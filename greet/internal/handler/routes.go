// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"
	"go-zero-demo/greet/internal/svc"
	"github.com/zeromicro/go-zero/rest"
	activity "go-zero-demo/greet/internal/handler/activity"
	course "go-zero-demo/greet/internal/handler/course"
	index "go-zero-demo/greet/internal/handler/index"
)

//注册路由到 server实例
func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.OperationRecord},
			[]rest.Route{
				{
					//请求方式： http://localhost:8888/soloActivity/list json格式 {"page":1,"pageSize":11}
					Method:  http.MethodGet,
					Path:    "/soloActivity/list",
					Handler: activity.GetSoloActivityList(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.OperationRecord},
			[]rest.Route{
				{
					//请求方式： http://localhost:8888/course/list json格式 {"page":1,"pageSize":11}
					Method:  http.MethodGet,
					Path:    "/course/list",
					Handler: course.GetCourseList(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.OperationRecord},
			[]rest.Route{
				{
					//请求方式： http://localhost:8888/jwt/test json格式 {"page":1,"pageSize":11}
					Method:  http.MethodGet,
					Path:    "/jwt/test",
					Handler: index.TestJwtToken(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/redis/test",
					Handler: index.TestRedis(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/get-request/test",
					Handler: index.TestGetRequest(serverCtx),
				},
				{
					//推荐所有请求都是用raw body；
					Method:  http.MethodGet,
					Path:    "/body-request/test",
					Handler: index.TestBodyRequest(serverCtx),
				},
			}...,
		),
	)
}
