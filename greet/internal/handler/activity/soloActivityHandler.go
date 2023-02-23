package activity

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-demo/greet/internal/logic/activity"
	"go-zero-demo/greet/internal/svc"
	"go-zero-demo/greet/internal/types"
	"net/http"
)

/*
 *  @Description: 获取家庭活动列表（command+ctrl+/）
 *  @param svcCtx
 *  @return http.HandlerFunc
 */
func GetSoloActivityList(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//将请求参数 进行映射
		var req types.PageReq
		if err := httpx.Parse(r, &req); err != nil {
			print("huangg")
			httpx.Error(w, err)
			return
		}

		//实例化logic并调用其方法
		l := activity.NewSoloActivityLogic(r.Context(), svcCtx)
		resp, err := l.GetSoloActivityList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
