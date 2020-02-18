# go-frame

- 本项目学习自[TomatoMr的awesomeframework项目](https://github.com/TomatoMr/awesomeframework)，目前代码基本一致，后期有待加入自己的特性。
- 附上TomatoMr的公众号(onepunchgo)以及原博客地址：[wx5e1abbbb0a5e5的博客](https://blog.51cto.com/14664952)

## 0. 项目目标

- 通用的开发框架，支持配置加载、日志记录、SQL数据库、Redis缓存、NSQ消息队列等功能，将应用接口(HTTP/WebSocket/RPC)与实际处理逻辑分离。

## 1. 项目结构

项目的主要结构如下图：

```go
├── cache   // 缓存模块
│   └── redis   // 具体的缓存模块，这里为redis
├── config  // 配置模块，基于yaml
│   ├── config.go
│   └── config.yaml
├── db      // 数据库模块，基于xorm和mysql
│   ├── db.go   
│   └── test.sql
├── logger  // 日志器模块，基于zap
│   └── logger.go
├── main.go // 入口文件
├── mq      // 消息队列模块
│   └── nsq     // nsq
├── process // 处理模块
│   ├── controller  // 真正的处理逻辑
│   ├── http        // http和websocket接口
│   └── rpc         // rpc接口，基于grpc和protobuf
└── utils   // 共用工具库

```

## 2. 项目启动流程

![workflow](/resources/images/workflow.png)

## 3. 配置

## 4. 日志

## 5. 存储

## 6. 缓存

## 7. 消息队列

## 8. HTTP接口

## 9. WebSocket

## 10. RPC

