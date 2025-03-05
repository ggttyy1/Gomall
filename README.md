# GoMall 项目介绍

## 项目概述

GoMall 是一个基于 Go 语言开发的电商平台，旨在为用户提供一个高效、稳定、可扩展的在线购物体验。项目采用微服务架构，将核心功能模块（如用户管理、商品管理、订单处理、支付系统等）进行解耦，每个服务独立部署和运行。服务之间通过 RPC（远程过程调用）进行通信，确保高效的数据交互和低延迟的性能表现。
GoMall 的设计目标是成为一个高性能、易于维护的电商解决方案，适用于中小型企业的快速部署和扩展。通过微服务架构和 RPC 通信机制，GoMall 能够灵活应对高并发场景，并支持服务的独立扩展和升级。  
请注意，GoMall 目前仍在积极开发和不断完善中，我们欢迎社区贡献和反馈，以帮助项目更快成熟。
---

## 主要功能

- **用户管理**：用户注册、登录、个人信息管理、地址管理等功能。
- **商品管理**：商品上传、商品分类、商品详情、商品搜索等功能。
- **购物车**：用户可以将商品加入购物车，进行批量结算。
- **订单管理**：订单创建、订单查询、订单状态更新等功能。
- **AI对话窗**：在AI对话框中实现页面跳转等功能。
---

## 技术栈

- **编程语言**：Go (Golang)
- **Web 框架**：hertz
- **RPC 框架**：Kitex
- **服务中心**：Consul
- **数据库**：MySQL
- **消息队列**：Nats
- **大模型**：llama3.2
- **缓存**：Redis
- **容器化**：Docker



---

## 快速开始

### 环境准备

1. 安装 Go 1.16 或更高版本。
2. 安装 Docker 和 Docker Compose。


### 克隆项目

```bash
git clone git@github.com:ggttyy1/Gomall.git
cd app
```
### 配置环境变量
将每个微服务中的 .env.example 文件并重命名为 .env，然后根据实际情况修改配置。
```bash
cp .env.example .env
```
### 启动中间件
使用 Docker Compose 启动所有中间件依赖：
```bash
docker_run.bat
```
### 运行项目前后端
```bash
go run cart/main.go
go run checkout/main.go
go run email/main.go
go run order/main.go
go run payment/main.go
go run person_infor/main.go
go run product/main.go
go run user/main.go
go run frontend/main.go
```
### 访问
```bash
http://localhost:8080
```

### 贡献指南
我们欢迎任何形式的贡献！如果你有任何建议或发现任何问题，请提交 Issue 或 Pull Request。

Fork 项目仓库。
1. 创建你的特性分支 (git checkout -b feature/AmazingFeature)。
2. 提交你的更改 (git commit -m 'Add some AmazingFeature')。
3. 推送到分支 (git push origin feature/AmazingFeature)。
4. 打开一个 Pull Request。

### 联系方式
如果你有任何问题或建议，请通过以下方式联系：

邮箱: jinshengli0412@163.com
感谢你对 GoMall 项目的关注和支持！希望这个项目能为你的电商业务带来便利和成功。