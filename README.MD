

# Chess

[![Build Status](https://img.shields.io/badge/build-1.01-brightgreen)](https://travis-ci.org/pibigstar/go-todo)

> 双人象棋

## 1.项目结构

<details>
<summary>展开查看</summary>
<pre>
    ├── app                        (项目文件)
        ├── api                    (前端api层)
        ├── cache                  (redis调用)
        ├── global                 (全局对象)
        ├── internal               (内部逻辑)
            ├── consts             (常量定义)
            ├── middleware         (中间件层)
            ├── model              (模型层)
            ├── service            (服务层)
        ├── resource               (资源文件)
            ├── public             (静态资源)
            ├── template           (模板文件)
        ├── router                 (路由层)
        ├── service                (微服务)
        ├── ui                     (象棋ui)
        ├── websocket              (websocket服务)
    ├── manifest                   (交付清单)
        ├── config                 (配置管理)
        ├── docs                   (swag接口文档)
        ├── docker                 (镜像文件)
        ├── sql                    (mysql初始化配置文件)
    ├── boot                       (项目启动文件)
    ├── log                        (日志目录)
    ├── utils                      (工具包)          
    ├── drone.yml                  (drone配置文件)
    ├── docker-compose.yaml        (多GRPC容器编排)
    ├── modd.conf                  (热配置更新)
</pre>
</details>

## 2. 使用技术

- [x] 基于`GoFrame v2.0`基础开发框架+`go-zero`微服务框架
- [x] Consul服务注册发现中心
- [x] Jaeger链路追踪
- [x] 基于gogs+drone的自动化集成部署
- [x] 
- [ ] 

### 2.1 后端框架

使用`GoFrame`框架, golang 版本: `1.8.2`

##### 特性

`GoFrame`不是一款`WEB/RPC`框架，而是一款通用性的基础开发框架，是`Golang`标准库的一个增强扩展级，包含通用核心的基础开发组件，优点是实战化、模块化、文档全面、模块丰富、易用性高、通用性强、面向团队。

使用``go-zero`框架，

##### 特性

- 高可用、高并发、易扩展

###### 

## 3. 部署

将代码push至gogs仓库，由web钩子自动请求drone进行集成部署



运行`go run main.go`开始游戏

### 

