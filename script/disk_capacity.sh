# 检查是否有磁盘使用率大于80%
count=0
for diskUsage in $(df -h | grep -v "overlay\|tmpfs\|Filesystem" | awk '{print $5}')
do
    diskUsage=${diskUsage::-1}
    if [ `echo "$diskUsage < 80" |bc` -eq 1 ]
    then
            count=`expr $count + 1`
    fi
done

if [ $count -gt 0 ]
then
    echo caseName:磁盘使用率小于80%, caseDesc:, result:fail, resultDesc:有${count}个磁盘使用率大于80%
else
    echo caseName:磁盘使用率小于80%, caseDesc:, result:success, resultDesc:磁盘使用率都小于80%
fi