package main

import (
	"flag"
	"fmt"

	"go-zero-demo/greet/internal/config"
	"go-zero-demo/greet/internal/handler"
	"go-zero-demo/greet/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

// 1、flag用于获取 命令行参数 -f
var configFile = flag.String("f", "etc/config.yaml", "the config file")

func main() {
	flag.Parse()

	//2、加载配置文件信息 到 config.go
	var c config.Config
	conf.MustLoad(*configFile, &c)

	//3、实例化服务器
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	//4、将配置信息 赋值给 ServiceContext
	ctx := svc.NewServiceContext(c)
	//5、注册路由到 server实例
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)

	//6、启动服务
	server.Start()
}
