package index

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-demo/greet/internal/svc"
	"go-zero-demo/greet/internal/types"
	"go-zero-demo/greet/utils"
	"time"
)

type JwtLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJwtLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JwtLogic {
	return &JwtLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JwtLogic) TestJwtToken(req *types.PageReq) (resp *types.Response, err error) {
	err = l.svcCtx.Rdb.Set("liu", "huiling", time.Duration(600)*time.Second).Err()
	if err != nil {
		fmt.Println("设置key失败")
		return
	}

	jwtToken, err := utils.GetJwtToken(l.svcCtx.Config.Auth.AccessSecret, map[string]interface{}{
		"name": "huangbin123",
		"age":  "18",
	})
	if err != nil {
		return types.NewDefaultError(err.Error()), nil
	}
	println("token")
	println(jwtToken)

	mycaim, err := utils.ParseToken(l.svcCtx.Config.Auth.AccessSecret, jwtToken)

	if err != nil {
		return types.NewDefaultError(err.Error()), nil
	}

	fmt.Printf("%#v", mycaim)
	result := mycaim["data"].(map[string]interface{})

	return &types.Response{
		Code: 0,
		Msg:  "Hello go-zero",
		Data: map[string]string{
			"token": jwtToken,
			"age":   result["age"].(string),
			"name":  result["name"].(string),
		},
	}, nil
}
