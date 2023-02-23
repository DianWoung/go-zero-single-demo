# go-zero-gva

### 1、介绍
`go-zero`练习 demo；创建目录结构`goctl api new greet`

### 2、启动方式
cd greet && go run main.go -f etc/config.yaml

### 3、软件架构

```
.
├── api
│   └── upload.api
├── etc
│   └── config.yaml
├── main.go
├── go.mod
├── go.sum
└── internal
    ├── config
    │   └── config.go
    ├── handler
    │   ├── routes.go
    │   └── uploadhandler.go
    ├── logic
    │   └── uploadlogic.go
    ├── svc
    │   └── servicecontext.go
    │   └── model.go
    └── types
        └── types.go
```
1. api目录：我们前面定义的API接口描述文件，无需多言
2. etc目录：这个是用来放置yaml配置文件的
3. main.go：main函数所在文件
4. internal/config目录：服务的配置定义
5. internal/handler目录：API文件里定义的路由对应的handler实现
6. internal/logic目录：用来放每个路由对应的业务处理逻辑，之所以区分handler和logic是为了让业务处理部分尽可能减少依赖，把HTTP requests和逻辑处理代码隔离开，便于后续按需拆分成RPC service
7. internal/svc目录：用来定义业务逻辑处理的依赖，我们可以在main里面创建依赖的资源，然后通过ServiceContext传递给handler和logic
8. internal/types目录：定义了API的request、response的struct，根据api自动生成，不建议编辑

### 4、添加mysql/redis数据源
1. `/etc/config.yaml`写入数据库链接、账号、密钥等参数
2. `main.go`将config.yaml格式 转换成 struct格式（internal/config/config.go）
3. `svc/servicecontext.go`中上下文 加载 数据库配置

### 5、添加jwt配置
1. `/etc/config.yaml`写入数据库链接、账号、密钥等参数
2. `main.go`将config.yaml格式 转换成 struct格式（internal/config/config.go）
3. `svc/servicecontext.go`默认已经加载了`config.go`的数据，所以可以读取jwt的数据
```goLand
type CourseLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func (l *CourseLogic) GetCourseList(req *types.PageReq) (resp *types.Response, err error) {
 println(l.svcCtx.Config.Auth.AccessSecret)
}
```

### 6、goctl 生成model文件
```goctl
goctl model mysql datasource -url="account:password@tcp(rm-wz9445yi0pv0t68a2wo.mysql.rds.aliyuncs.com:3306)/evaluation" -table="course,course_snapshot,course_student" -dir="./"
```
1. 因model文件在 internal/svc 下，进入到该目录执行命令
2. `-table`参数表示生成指定的表
3. `-dir`参数表示在当前目录下 生成的model的目录结构

### 8、编写业务代码步骤
1. 编写流程 handler -> logic-> types -> routes
2. handler 写 控制器入口
3. logic 写 业务逻辑入口
4. types 写 业务层所需要的struct，比如 业务中间流程需要声明的结构
5. routes 写 路由

### 9、中间件搭建流程
1. middleware目录添加 中间件文件
2. 在路由文件新增middleware
```goLand
server.AddRoutes(
    rest.WithMiddlewares([]rest.Middleware{serverCtx.OperationRecord}, 
    []rest.Route{}...
)
```
3. `servicecontext.go` 文件添加中间件依赖

 

