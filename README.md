## 提醒小工具
现已完成 事项提醒、日期记录、签到功能

### 前端
微信小程序 taro3 + vue3 + nutui + typeScript + pinia

### 后端
基于大佬的 go-zero-looklook 项目快速开发

go-zero
nginx网关
prometheus
asynq
asynqmon
etcd
dtm
docker
docker-compose
mysql
redis
modd

## 功能介绍
### 事项提醒
使用 asynq 延迟队列，到点调用微信小程序通知接口

### 日期记录 

### 签到功能
基于 dtm 实现跨服务分布式事务控制，保障签到记录、任务完成、积分增减业务的数据一致性。

基于 asynq 的 shedule 实现用户订阅签到提醒