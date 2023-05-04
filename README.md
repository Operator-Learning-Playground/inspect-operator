## inspect-operator 简易型集群内巡检中心
![](https://github.com/googs1025/message-operator/blob/main/image/%E6%B5%81%E7%A8%8B%E5%9B%BE%20(1).jpg?raw=true)
### 项目思路与设计
设计背景：本项目基于k8s的扩展功能，实现Inspect的自定义资源，实现一个集群内的执行bash脚本或是自定义镜像的controller应用。调用方可在cluster中部署与启动相关配置即可使用。
思路：当应用启动后，会启动一个controller，controller会监听所需的资源，并执行相应的业务逻辑(如：执行巡检脚本或镜像，再使用集群内的消息中心进行通知)。

### 项目功能
1. 支持对集群内使用job执行用户自定义镜像内容功能(用户必须完成image开发部分，可参考test/try目录)。
2. 支持对本地节点执行bash脚本功能，其中提供内置巡检bash脚本或用户可自定义bash脚本内容。
3. 提供发送结果通知功能(使用集群内消息中心operator实现)。


### RoadMap
1. 实现远端局点执行脚本的能力
2. 优化发送结果回调通知的能力
3. 实现下发cronjob定时巡检能力
