package common

import (
	"github.com/gorhill/cronexpr"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

// IsCronExpressionValid
// Same as org.quartz.CronExpression.IsExpressValid(string) boolean
func IsCronExpressionValid(spec string) bool {
	_, err := cronexpr.Parse(spec)
	return err == nil
}

func CronNextExecTime(spec string) time.Time {
	expression, _ := cronexpr.Parse(spec)
	return expression.Next(time.Now())
}

var cronNormalizer = strings.NewReplacer(
	"@yearly", "0 0 0 1 1 * *",
	"@annually", "0 0 0 1 1 * *",
	"@monthly", "0 0 0 1 * * *",
	"@weekly", "0 0 0 * * 0 *",
	"@daily", "0 0 0 * * * *",
	"@hourly", "0 0 * * * * *")
var fieldFinder = regexp.MustCompile(`\S+`)

var (
	layoutValue               = regexp.MustCompile(`^(0?[0-9]|1[0-9]|2[0-3])$`)
	layoutMultiValue          = regexp.MustCompile(`^((0?[0-9]|1[0-9]|2[0-3]),)+(0?[0-9]|1[0-9]|2[0-3])$`)
	layoutRange               = regexp.MustCompile(`^(0?[0-9]|1[0-9]|2[0-3])-(0?[0-9]|1[0-9]|2[0-3])$`)
	layoutWildcardAndInterval = regexp.MustCompile(`^\*/(\d+)$`)
	layoutValueAndInterval    = regexp.MustCompile(`^(0?[0-9]|1[0-9]|2[0-3])/(\d+)$`)
	layoutRangeAndInterval    = regexp.MustCompile(`^(0?[0-9]|1[0-9]|2[0-3])-(0?[0-9]|1[0-9]|2[0-3])/(\d+)$`)
)

func ChangeCronExpressionTimeZone(spec string) string {
	if k8sTimeZone == CronTimeZone {
		return spec
	}
	// Maybe one of the built-in aliases is being used
	cron := cronNormalizer.Replace(spec)
	indices := fieldFinder.FindAllStringIndex(cron, -1)

	hourSetting := spec[indices[1][0]:indices[1][1]]
	b := []byte(hourSetting)
	after := hourSetting
	if layoutValue.Match(b) {
		// 5
		after = strconv.Itoa(changeHourTimeZone(atoi(hourSetting)))
	} else if layoutMultiValue.Match(b) {
		// 5,6,7
		hours := strings.Split(hourSetting, ",")
		buf := strings.Builder{}
		for i, hour := range hours {
			buf.WriteString(strconv.Itoa(changeHourTimeZone(atoi(hour))))
			if i != len(hours)-1 {
				buf.WriteByte(',')
			}
		}
		after = buf.String()
	} else if layoutRange.Match(b) {
		// 5-10
		rangeHour := strings.Split(hourSetting, "-")
		after = changeRangeHourTimeZone(atoi(rangeHour[0]), atoi(rangeHour[1]), 1)
	} else if layoutWildcardAndInterval.Match(b) {
		// */2
		cycleHour := strings.Split(hourSetting, "/")
		after = changeRangeHourTimeZone(0, 23, atoi(cycleHour[1]))
	} else if layoutValueAndInterval.Match(b) {
		// 5/2
		cycleHour := strings.Split(hourSetting, "/")
		after = changeRangeHourTimeZone(atoi(cycleHour[0]), 23, atoi(cycleHour[1]))
	} else if layoutRangeAndInterval.Match(b) {
		// 5-10/2
		cycleHour := strings.Split(hourSetting, "/")
		rangeHour := strings.Split(cycleHour[0], "-")
		after = changeRangeHourTimeZone(atoi(rangeHour[0]), atoi(rangeHour[1]), atoi(cycleHour[1]))
	}
	return cron[:indices[1][0]] + after + cron[indices[1][1]:]
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

var (
	k8sTimeZone  = 8
	CronTimeZone = 8
)

func changeHourTimeZone(hour int) int {
	offset := k8sTimeZone - CronTimeZone
	return (hour + offset + 24) % 24
}

func changeRangeHourTimeZone(start, end, period int) string {
	result := make([]int, 0, 24)
	for i := start; i <= end; i += period {
		result = append(result, changeHourTimeZone(i))
	}
	sort.Ints(result)
	buf := strings.Builder{}
	for i, hour := range result {
		buf.WriteString(strconv.Itoa(hour))
		if i != len(result)-1 {
			buf.WriteByte(',')
		}
	}
	return buf.String()
}
