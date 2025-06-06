
将根目录下的 .env_example 替换为 .env 并将配置秘钥都填上

搭建请阅读项目根目录下的 /doc/chinese/ 下的完整教程，从 01-开发环境搭建环境。
## go-zero-looklook

### 创建 kafka topic
系统内使用了1个topic，默认是不允许程序自动创建topic的，进入kafka的容器创建1个topic

进入容器
```
$ docker exec -it kafka /bin/sh
$ cd /opt/kafka/bin/
```
创建topic
```
$ ./kafka-topics.sh --create --zookeeper 172.27.103.240:2181 --replication-factor 1 -partitions 1 --topic notice-send-topic
```

### 常见错误
#### 签到找不到服务
error: produced zero addresses

查看一下 dtm 服务是否注册到了 etcd 中
```
docker exec -it etcd /bin/bash
etcdctl get "dtmservice" --prefix
```
若没有 则重启一下 dtm

- k8s
- go-zero
- nginx网关
- filebeat
- kafka
- go-stash
- elasticsearch
- kibana
- prometheus
- grafana
- jaeger
- go-queue
- asynq
- asynqmon
- dtm
- docker
- docker-compose
- mysql
- redis
- modd
- jenkins
- gitlab
- harbor


## 项目简介

整个项目使用了go-zero开发的微服务，基本包含了go-zero以及相关go-zero作者开发的一些中间件，所用到的技术栈基本是go-zero项目组的自研组件，基本是go-zero全家桶了


项目目录结构如下：


- app：所有业务代码包含api、rpc以及mq（消息队列、延迟队列、定时任务）
- common：通用组件 error、middleware、interceptor、tool、ctxdata等
- data：该项目包含该目录依赖所有中间件(mysql、es、redis、grafana等)产生的数据，此目录下的所有内容应该在git忽略文件中，不需要提交。
- deploy：

    - filebeat: docker部署filebeat配置
    - go-stash：go-stash配置
    - nginx: nginx网关配置
    - prometheus ： prometheus配置
    - script：
        - gencode：生成api、rpc，以及创建kafka语句，复制粘贴使用
        - mysql：生成model的sh工具
    - goctl: 该项目goctl的template，goctl生成自定义代码模版，template用法可参考go-zero文档，复制到家目录下.goctl即可
- doc : 该项目系列文档
- modd.conf :  modd热加载配置文件，不要怕～它用起来很简单，关于modd更多用法可以去这里了解 ： https://github.com/cortesi/modd ， 本项目镜像只是将golang-1.24.1-alpine作为基础镜像安装了modd在内部，如果你想把goctl、protoc、golint等加进去，不用我的镜像直接制作一个镜像也一样的哈


## 网关

nginx做对外网关，网关前面是slb，另外，很多同学觉得nginx做网关不太好，这块原理基本一样，可以自行替换成apisix、kong等



## 开发模式

本项目使用的是微服务开发，api （http） + rpc（grpc） ， api充当聚合服务，复杂、涉及到其他业务调用的统一写在rpc中，如果一些不会被其他服务依赖使用的简单业务，可以直接写在api的logic中



## 日志

关于日志，统一使用filebeat收集，上报到kafka中，由于logstash懂得都懂，资源占用太夸张了，这里使用了go-stash替换了logstash

链接：https://github.com/kevwan/go-stash  ， go-stash是由go-zero开发团队开发的，性能很高不占资源，主要代码量没多少，只需要配置就可以使用，很简单。它是把kafka数据源同步到elasticsearch中，默认不支持elasticsearch账号密码，我fork了一份修改了一下，很简单支持了账号、密码



## 监控

监控采用prometheus，这个go-zero原生支持，只需要配置就可以了，这里可以看项目中的配置



## 链路追踪

go-zero默认jaeger、zipkin支持，只需要配置就可以了，可以看配置



## 发布订阅

kafka ， 发布订阅使用的是go-zero开发团队开发的go-queue，链接：https://github.com/zeromicro/go-queue

这里使用kq，kq是基于kafka做的高性能发布订阅



## 消息队列、延迟队列、定时任务

消息队列、延迟队列、定时任务本项目使用的是asynq ，基于redis开发的简单中间件，

当然，消息队列你也可以使用go-queue

链接：https://github.com/hibiken/asynq



## 分布式事务

分布式事务准备使用的是dtm， 嗯 ，很舒服，之前我写过一篇 "go-zero对接分布式事务dtm保姆式教程" 链接地址：https://github.com/Mikaelemmmm/gozerodtm ， 本项目目前还未使用到，后续准备直接集成就好了，如果读者使用直接去看那个源码就行了



## 部署

本项目开发环境推荐docker-compose，使用直链方式，放弃服务注册发现中间件（etcd、nacos、consul等）带来的麻烦

测试、线上部署使用k8s（也不需要etcd、nacos、consul等）有详细教程（搭建+部署），可以进go-zero社区群沟通，非常easy

文档地址：https://github.com/Mikaelemmmm/go-zero-looklook/tree/main/doc

gitlab + jenkins + harbor + k8s

在jenkins中点击部署对应的服务，会去gitlab拉取代码-->再去拉取线上配置（线上配置单独一个git库，为什么不用配置中心，部署文档中有介绍）---->自动构建镜像-->推送到harbor镜像仓库--->使用kubectl自动发布到k8s中---->前面要挂一个nginx做网关统一入口



## 感谢

go-zero-looklook: https://github.com/Mikaelemmmm/go-zero-looklook

go-zero 微服务: https://github.com/zeromicro/go-zero

dtm分布式事务：https://github.com/dtm-labs/dtm


