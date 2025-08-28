# Review Service 评价服务系统
一个基于 Go + Kratos 框架构建的微服务评价系统，支持用户评价、商家回复、申诉审核等完整的评价业务流程。

## 系统架构
本项目采用微服务架构，包含两个主要服务：

- review-service : 核心评价服务（端口：HTTP 8082, gRPC 9092）
- review-b : 商家端服务（端口：HTTP 8083, gRPC 9093）
## 功能特性
### 核心功能
-  用户评价管理 : 创建评价、查看评价详情、用户评价列表
-  商家回复系统 : 商家可对用户评价进行回复
-  申诉审核流程 : 商家申诉不当评价，运营审核处理
-  评价审核机制 : 运营人员审核用户评价内容
-  多媒体支持 : 支持图片和视频的评价内容
-  匿名评价 : 支持用户匿名发表评价
### 技术特性
-  基于 Kratos v2 微服务框架
-  使用 gRPC 进行服务间通信
-  MySQL 数据持久化 + Redis 缓存
-  GORM ORM 框架，支持数据库迁移
-  Snowflake 分布式ID生成
-  Protocol Buffers API 定义和验证
-  Docker 容器化部署
##  项目结构
```
review_service/
├── review-service/          # 核心评价服务
│   ├── api/                # API 定义 
(protobuf)
│   ├── cmd/                # 应用入口
│   ├── configs/            # 配置文件
│   ├── internal/           # 内部业务逻辑
│   │   ├── biz/           # 业务逻辑层
│   │   ├── data/          # 数据访问层
│   │   ├── service/       # 服务层
│   │   └── server/        # 服务器配置
│   ├── pkg/               # 公共包
│   └── review.sql         # 数据库表结构
├── review-b/               # 商家端服务
│   ├── api/               # API 定义
│   ├── internal/          # 业务逻辑
│   └── configs/           # 配置文件
└── README.md
```
## 快速开始
### 环境要求
- Go 1.19+
- MySQL 5.7+
- Redis 6.0+
- Docker (可选)
### 1. 克隆项目
```
git clone https://github.com/Asasend/
review_service.git
cd review_service
```
### 2. 数据库初始化
```
# 创建数据库
mysql -u root -p -e "CREATE DATABASE 
review_service;"

# 导入表结构
mysql -u root -p review_service < 
review-service/review.sql
```
### 3. 配置文件
修改配置文件中的数据库连接信息：

- review-service/configs/config.yaml
- review-b/configs/config.yaml
### 4. 启动服务
启动核心评价服务:

```
cd review-service
go mod tidy
go run cmd/review-service/main.go -conf 
configs/
```
启动商家端服务:

```
cd review-b
go mod tidy
go run cmd/review-b/main.go -conf configs/
```
### 5. Docker 部署 (可选)
```
# 构建镜像
docker build -t review-service ./
review-service
docker build -t review-b ./review-b

# 运行容器
docker run -p 8082:8082 -p 9092:9092 
review-service
docker run -p 8083:8083 -p 9093:9093 review-b
```
## API 接口
### 核心评价服务 (review-service)
- POST /v1/review - 创建评价
- GET /v1/review/{reviewID} - 获取评价详情
- POST /v1/review/audit - 审核评价
- POST /v1/review/reply - 回复评价
- POST /v1/review/appeal - 申诉评价
- POST /v1/appeal/audit - 审核申诉
- GET /v1/{userID}/reviews - 用户评价列表
### 商家端服务 (review-b)
- POST /business/v1/review/reply - 商家回复评价
## 数据库设计
### 主要数据表
- review_info : 评价信息表
- review_reply_info : 评价回复表
- review_appeal_info : 评价申诉表
## 开发指南
### 生成 API 代码
```
# 安装 kratos 工具
go install github.com/go-kratos/kratos/cmd/
kratos/v2@latest

# 生成 proto 代码
make api

# 生成服务代码
kratos proto server api/review/v1/review.
proto -t internal/service
```
### 依赖注入
```
# 安装 wire
go install github.com/google/wire/cmd/
wire@latest

# 生成依赖注入代码
cd cmd/review-service && wire
```
