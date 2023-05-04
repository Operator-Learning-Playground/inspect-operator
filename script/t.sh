# 检查是否有僵尸进程
export TERM=xterm
zombieCount=$(top -bn 1 |grep 'Tasks' | awk '{print $10}')
if [ $zombieCount -gt 0 ]
then
    echo caseName:无僵尸进程, caseDesc:, result:fail, resultDesc:有${zombieCount}个僵尸进程
else
    echo caseName:无僵尸进程, caseDesc:, result:success, resultDesc:无僵尸进程
fi