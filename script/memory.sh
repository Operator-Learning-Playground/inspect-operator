# 检查内存的空闲率是否大于20%
a=`cat /proc/meminfo | sed -n '/MemAvailable/p'| awk '{print $2}'`;
t=`cat /proc/meminfo | sed -n '/MemTotal/p'| awk '{print $2}'`;
r=`echo "scale=4; ( $a / $t) * 100" | bc`;
echo -e "FreeMem Ratio (%): $r"

if [ `echo ${r} | awk -v tem=20 '{print($1>tem)? "1":"0"}'` -eq "0" ]
then
    echo caseName:内存的使用率小于80%, caseDesc:, result:fail, resultDesc:内存的使用率大于80%
else
    echo caseName:内存的使用率小于80%, caseDesc:, result:success, resultDesc:内存的使用率小于80%
fi