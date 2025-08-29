# Kratos Project Template

## Install Kratos
```
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```
## Create a service
```
# Create a template project
kratos new server

cd server
# Add a proto template
kratos proto add api/server/server.proto
# Generate the proto code
kratos proto client api/server/server.proto
# Generate the source code of service by proto file
kratos proto server api/server/server.proto -t internal/service

go generate ./...
go build -o ./bin/ ./...
./bin/server -conf ./configs
```
## Generate other auxiliary files by Makefile
```
# Download and update dependencies
make init
# Generate API files (include: pb.go, http, grpc, validate, swagger) by proto file
make api
# Generate all files
make all
```
## Automated Initialization (wire)
```
# install wire
go get github.com/google/wire/cmd/wire

# generate wire
cd cmd/server
wire
```

## Docker
```bash
# build
docker build -t <your-docker-image-name> .

# run
docker run --rm -p 8000:8000 -p 9000:9000 -v </path/to/your/configs>:/data/conf <your-docker-image-name>
```
# Review Job Service

## 项目简介

评价系统数据同步服务，负责将 MySQL 中的评价数据实时同步到 Elasticsearch。

## 核心功能

- **Kafka 消息消费**：监听评价数据变更消息
- **数据同步**：将评价数据写入 Elasticsearch
- **搜索支持**：为评价系统提供全文搜索能力

## 技术栈

- **框架**：Kratos v2
- **消息队列**：Kafka
- **搜索引擎**：Elasticsearch
- **语言**：Go

## 配置说明

```yaml
kafka:
  brokers: ["localhost:9092"]
  group_id: "review-job"
  topic: "topic3"

elasticsearch:
  addresses: ["http://localhost:9200"]
  index: "review"
```

## 快速启动

```bash
# 生成配置文件
make config
# 或者：
protoc --proto_path=./internal \
       --proto_path=./third_party \
       --go_out=paths=source_relative:./internal \
       internal/conf/conf.proto

# 编译运行
go build -o ./bin/ ./...
./bin/review-job -conf ./configs
```

## 架构角色
MySQL → Kafka → review-job → Elasticsearch → 搜索功能

本服务作为数据管道，确保评价数据的实时同步和搜索能力。