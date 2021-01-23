package handler

import (
	"net/http"

	"payment/internal/logic"
	"payment/internal/svc"

	"github.com/tal-tech/go-zero/rest/httpx"
)

// PingHandler ping响应
func PingHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewPingLogic(r.Context(), ctx)
		err := l.Ping(w)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
