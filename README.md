## inspect-operator 简易型集群内巡检中心
![](https://github.com/Operator-Learning-Playground/inspect-operator/blob/main/image/%E6%B5%81%E7%A8%8B%E5%9B%BE%20(1).jpg?raw=true)
### 项目思路与设计
设计背景：本项目基于k8s的扩展功能，实现Inspect的自定义资源，实现一个集群内的执行bash脚本或是自定义镜像的controller应用。调用方可在cluster中部署与启动相关配置即可使用。
思路：当应用启动后，会启动一个controller，controller会监听所需的资源，并执行相应的业务逻辑(如：执行巡检脚本或镜像，再使用集群内的消息中心进行通知)。

### 项目功能
1. 支持对集群内使用job执行用户自定义镜像内容功能(用户必须完成image开发部分，可参考test/try目录)。
2. 支持对本地节点执行bash脚本功能，其中提供内置巡检bash脚本或用户可自定义bash脚本内容。
3. 提供发送结果通知功能(使用集群内消息中心operator实现)。

- 自定义资源如下所示
```yaml
apiVersion: api.practice.com/v1alpha1
kind: Inspect
metadata:
  name: myinspect
spec:
  tasks:
    - task:
        task_name: task1
        type: script             # type字段：可填写脚本或镜像 image script 两种
        source: test.sh          # source字段：可执行 是bash 脚本 py 脚本或是镜像。需要把东西放入./script中
        script_location: local   # script_location字段(目前未支持此功能) 可选填 local remote all 三种，分别对应 本地节点 远端节点 全部节点
        # 选取local本地节点 就不需要再填写远端ip地址
        # 远端要执行的目标node
        # script字段：可填写bash脚本内容，controller默认如果有script字段，优先执行自定义脚本内容，"不执行"source字段脚本内容
        script: |
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
    - task:
        task_name: task2
        type: image
        source: try:latest  # source字段：镜像名称
        restart: true       # 用于标示是否重新执行。 如
    - task:
        task_name: task3
        type: script             # type字段：可填写脚本或镜像 image script 两种
        script_location: remote   # script_location字段(目前未支持此功能) 可选填 local remote all 三种，分别对应 本地节点 远端节点 全部节点
        # 远端要执行的目标node的信息：user password ip等
        remote_ips:
          - user: "root"
            password: "xxxxxx"
            ip: "xxxxxx"
        script: |
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


```

### RoadMap
1. 实现远端局点执行脚本的能力
2. 优化发送结果回调通知的能力
3. 实现下发cronjob定时巡检能力
