# 检查cpu的空闲率是否小于20%
cpuCount=$[$(vmstat -SM | awk '{if ($15 < 20) print $0}' | wc -l)-1]
if [ $cpuCount -gt 0 ]
then
    echo caseName:cpu的使用率小于80%, caseDesc:, result:fail, resultDesc:有${cpuCount}个cpu的使用率大于80%
else
    echo caseName:cpu的使用率小于80%, caseDesc:, result:success, resultDesc:cpu的使用率都小于80%
fi