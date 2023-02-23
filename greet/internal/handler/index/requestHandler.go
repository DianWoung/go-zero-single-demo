package index

import (
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-demo/greet/internal/logic/index"
	"go-zero-demo/greet/internal/svc"
	"go-zero-demo/greet/internal/types"
	"net/http"
	"strconv"
)

/*
 *  @Description: Get请求参数获取
 *  @param svcCtx
 *  @return http.HandlerFunc
 */
func TestGetRequest(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//正确请求 http://localhost:8888/get-request/test?page=21&pageSize=10
		//错误请求 http://localhost:8888/get-request/test?page000=21

		var req types.PageReq
		var err error

		query := r.URL.Query()
		fmt.Println(query)             //map[page:[21] pageSize:[10]]
		fmt.Println(query.Get("page")) //21

		if req.Page, err = strconv.Atoi(query.Get("page")); err != nil {
			//如果参数不存在就会报错
			httpx.OkJson(w, types.NewDefaultError(err.Error()))
			return
		}

		if req.PageSize, err = strconv.Atoi(query.Get("pageSize")); err != nil {
			httpx.OkJson(w, types.NewDefaultError(err.Error()))
			return
		}

		fmt.Println(req)

		//实例化logic并调用其方法
		l := index.NewJwtLogic(r.Context(), svcCtx)
		resp, err := l.TestJwtToken(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

/*
 *  @Description: Body请求参数获取
 *  @param svcCtx
 *  @return http.HandlerFunc
 */
func TestBodyRequest(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//http://localhost:8888/body-request/test json格式 {"page":1,"pageSize":11}

		//将请求参数 进行映射
		var req types.PageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		//实例化logic并调用其方法
		l := index.NewJwtLogic(r.Context(), svcCtx)
		resp, err := l.TestJwtToken(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
