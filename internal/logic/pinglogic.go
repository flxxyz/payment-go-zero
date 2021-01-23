package logic

import (
	"context"
	"net/http"

	"payment/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type PingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) PingLogic {
	return PingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingLogic) Ping(w http.ResponseWriter) (err error) {
	_, err = w.Write([]byte("pong"))
	return
}
