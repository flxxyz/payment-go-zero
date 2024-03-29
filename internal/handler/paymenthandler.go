package handler

import (
	"net/http"

	"payment/internal/logic"
	"payment/internal/svc"
	"payment/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func PaymentHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewPaymentLogic(r.Context(), ctx)
		resp, err := l.Payment(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
