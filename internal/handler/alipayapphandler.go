package handler

import (
	"net/http"

	"payment/internal/logic"
	"payment/internal/svc"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func AlipayAppHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewAlipayAppLogic(r.Context(), ctx)
		err := l.AlipayApp()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
