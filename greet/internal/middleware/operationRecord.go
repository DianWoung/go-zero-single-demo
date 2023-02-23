package middleware

import (
	"net/http"
)

type OperationRecord struct {
}

func NewOperationRecord() *OperationRecord {
	return &OperationRecord{}
}

func (m *OperationRecord) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		println("operationRecord middleware...")

		////将请求参数 进行映射
		//var req types.PageReq
		//fmt.Println(httpx.Parse(r, &req))

		//if err := httpx.Parse(r, &req); err != nil {
		//	println("huangbbbbbbb.")
		//
		//	//httpx.Error(w, err)
		//	//return
		//}
		////fmt.Println(req)

		next(w, r)
	}
}
