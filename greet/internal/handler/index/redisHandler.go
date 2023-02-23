package index

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-demo/greet/internal/logic/index"
	"go-zero-demo/greet/internal/svc"
	"net/http"
)

/*
 *  @Description: jwt测试
 *  @param svcCtx
 *  @return http.HandlerFunc
 */
func TestRedis(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//实例化logic并调用其方法
		l := index.NewRedisLogic(r.Context(), svcCtx)
		resp, err := l.TestRedis()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
