package course

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-demo/greet/internal/logic/course"
	"go-zero-demo/greet/internal/svc"
	"go-zero-demo/greet/internal/types"
	"net/http"
)

/*
 *  @Description: 获取课程信息（command+ctrl+/）
 *  @param svcCtx
 *  @return http.HandlerFunc
 */
func GetCourseList(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//将请求参数 进行映射
		var req types.PageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		//实例化logic并调用其方法
		l := course.NewCourseLogic(r.Context(), svcCtx)
		resp, err := l.GetCourseList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
