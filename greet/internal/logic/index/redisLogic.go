package index

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-demo/greet/internal/svc"
	"go-zero-demo/greet/internal/types"
	"time"
)

type RedisLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRedisLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RedisLogic {
	return &RedisLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RedisLogic) TestRedis() (resp *types.Response, err error) {
	//官方文档： https://pkg.go.dev/github.com/go-redis/redis#section-readme
	l.Logger.Info("testtt")
	l.Logger.Info("testtt")

	//设置key
	err = l.svcCtx.Rdb.Set("liu", "huiling", time.Duration(600)*time.Second).Err()
	if err != nil {
		fmt.Println("设置key失败")
		return
	}

	//获取key
	value, err := l.svcCtx.Rdb.Get("liu").Result()
	if err != nil {
		fmt.Println("获取key失败")
		return
	}
	fmt.Println(value)

	return &types.Response{
		Code: 0,
		Msg:  "Hello go-zero",
		Data: map[string]string{
			"result": value,
		},
	}, nil
}
