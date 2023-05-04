# 检查是否有CPU降频
count=0
for cpuHz in $(cat /proc/cpuinfo | grep MHz | awk '{print $4}')
do
    if [ `echo "$cpuHz < 2000.0" |bc` -eq 1 ]
    then
        count=`expr $count + 1`
    fi
done
if [ $count -gt 0 ]
then
    echo caseName:无CPU降频, caseDesc:, result:fail, resultDesc:有${count}个CPU的频率低于2000MHz, 可能发生降频
else
    echo caseName:无CPU降频, caseDesc:, result:success, resultDesc:CPU频率都大于2000MHz, 无降频
fi