package handler

import (
	"net/http"

	"payment/internal/logic"
	"payment/internal/svc"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func WechatWapHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewWechatWapLogic(r.Context(), ctx)
		err := l.WechatWap()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
