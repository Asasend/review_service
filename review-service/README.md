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

# 开发接口流程

## 1. 定义API文件

按照需要编写proto文件：
```proto
// 创建服务
rpc CreateReview (CreateReviewRequest) returns (CreateReviewReply){
    option (google.api.http) = {
        post: "/v1/review",
        body: "*"
    };
};
```

```proto
// 创建评价的参数
message CreateReviewRequest {
	int64 userID = 1;
	int64 orderID = 2;
	int32 score = 3;
	int32 serviceScore = 4;
	int32 expressScore = 5;
	string content = 6;
	string picInfo = 7;
	string videoInfo = 8;
	bool annoymous = 9;
}
// 评价的回复
message CreateReviewReply {
	int64 reviewID = 1;
}

```

 
## 2. 生成客户端和服务端代码
```bash
kratos proto server api/review/v1/review.proto -t internal/service

kratos proto client api/review/v1/review.proto 
```

## 3. 填充业务逻辑
internal目录下

`server -> service -> biz -> data`

## 4. 更新ProvideSet执行Wire实现依赖注入
默认的还是GreetService，需要换成我们的服务。



# validate参数校验和错误处理
## 当前版本代码的参数校验和错误处理实现详细总结
基于对代码的深入分析，当前版本已经完整实现了Kratos框架的参数校验（validate）和错误处理机制。以下是详细的实现流程和机制：

### 1. 参数校验（Validate）实现 1.1 Proto文件中的校验规则定义
在 `review.proto` 中定义了详细的校验规则：

```
// 创建评价的参数
message CreateReviewRequest {
    int64 userID = 1 [(validate.
    rules).int64 = {gt: 0}];
    int64 orderID = 2 [(validate.
    rules).int64 = {gt: 0}];
    int32 score = 3 [(validate.
    rules).int32 = {gt: 0, lte: 5}];
    int32 serviceScore = 4 
    [(validate.rules).int32 = {in: 
    [1,2,3,4,5]}];
    int32 expressScore = 5 
    [(validate.rules).int32 = {in: 
    [1,2,3,4,5]}];
    string content = 6 [(validate.
    rules).string = {min_len: 8, 
    max_len: 255}];
    string picInfo = 7;
    string videoInfo = 8;
    bool annoymous = 9;
}
```
校验规则说明：

- userID 和 orderID ：必须大于0
- score ：必须大于0且小于等于5
- serviceScore 和 expressScore ：只能是1-5中的值
- content ：字符串长度必须在8-255之间 1.2 自动生成校验代码
通过 protoc 命令生成 `review.pb.validate.go` 文件，包含：

- Validate() 方法：校验单个字段，遇到第一个错误即返回
- ValidateAll() 方法：校验所有字段，收集所有错误后返回
- 自定义错误类型： CreateReviewRequestValidationError 1.3 框架层自动校验
在 `http.go` 和 `grpc.go` 中配置了校验中间件：

```go
func NewHTTPServer(c *conf.Server, 
reviewer *service.ReviewService, 
logger log.Logger) *http.Server {
    var opts = []http.ServerOption{
        http.Middleware(
            recovery.Recovery(),
            validate.Validator
            (), // 参数校验中间件
        ),
    }
    // ... existing code ...
}
```
### 2. 错误处理（Errors）实现
#### 2.1 错误码定义
在 `review_error.proto` 中定义业务错误码：

```go
enum ErrorReason {
  // 设置缺省错误码
  option (errors.default_code) = 
  500;

  // 为某个枚举单独设置错误码
   NEED_LOGIN = 0 [(errors.code) = 
   401];
   DB_FAILED = 1 [(errors.code) = 
   500];
   ORDER_REVIEWED = 2 [(errors.
   code) = 400];
}
``` 
#### 2.2 自动生成错误处理函数
通过 protoc 命令生成 `review_error_errors.pb.go` 文件，包含：

```go
// 错误创建函数
func ErrorNeedLogin(format string, 
args ...interface{}) *errors.Error {
    return errors.New(401, 
    ErrorReason_NEED_LOGIN.String
    (), fmt.Sprintf(format, 
    args...))
}

func ErrorDbFailed(format string, 
args ...interface{}) *errors.Error {
    return errors.New(500, 
    ErrorReason_DB_FAILED.String(), 
    fmt.Sprintf(format, args...))
}

func ErrorOrderReviewed(format 
string, args ...interface{}) 
*errors.Error {
    return errors.New(400, 
    ErrorReason_ORDER_REVIEWED.
    String(), fmt.Sprintf(format, 
    args...))
}

// 错误判断函数
func IsNeedLogin(err error) bool { /
* ... */ }
func IsDbFailed(err error) bool { /
* ... */ }
func IsOrderReviewed(err error) 
bool { /* ... */ }
``` 
#### 2.3 业务层错误处理
在 `review.go` 中使用标准错误处理：

```go
func (uc *ReviewUsecase) 
CreateReview(ctx context.Context, 
review *model.ReviewInfo) (*model.
ReviewInfo, error) {
    // ... existing code ...
    
    // 数据库查询错误处理
    reviews, err := uc.repo.
    GetReviewByOrderID(ctx, review.
    OrderID)
    if err != nil {
        return nil, v1.ErrorDbFailed
        ("查询数据库失败")
    }
    
    // 业务逻辑错误处理
    if len(reviews) > 0 {
        return nil, v1.
        ErrorOrderReviewed("订单%d已
        评价", review.OrderID)
    }
    
    // ... existing code ...
}
```
### 3. 完整的实现流程 
#### 3.1 请求处理流程
1. HTTP/gRPC请求接收 ：框架接收请求并解析参数
2. 参数绑定 ：将请求参数绑定到proto消息结构体
3. 自动校验 ： validate.Validator() 中间件自动调用 req.Validate() 方法
4. 校验失败处理 ：如果校验失败，直接返回400错误，不进入业务逻辑
5. 业务逻辑处理 ：校验通过后，进入service层和biz层处理
6. 业务错误处理 ：在biz层使用标准错误码返回业务错误
7. 统一错误响应 ：框架自动将错误转换为标准HTTP/gRPC响应 
#### 3.2 代码生成流程
8. 定义proto文件 ：在 .proto 文件中定义校验规则和错误码
9. 生成校验代码 ：使用 protoc --go-validate_out 生成校验代码
10. 生成错误码 ：使用 protoc --go-errors_out 生成错误处理代码
11. 集成中间件 ：在server配置中添加校验和错误处理中间件
### 4. 关键特性
- 自动化 ：参数校验完全自动化，无需手动编写校验代码
- 类型安全 ：基于proto定义，编译时类型检查
- 标准化 ：错误码和错误消息格式统一
- 可扩展 ：支持自定义校验规则和错误类型
- 多协议支持 ：同时支持HTTP和gRPC协议
- 中间件集成 ：通过中间件实现横切关注点的统一处理