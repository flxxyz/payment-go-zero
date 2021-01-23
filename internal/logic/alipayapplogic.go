package logic

import (
	"context"

	"github.com/tal-tech/go-zero/core/logx"
	"payment/internal/svc"
)

type AlipayAppLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAlipayAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) AlipayAppLogic {
	return AlipayAppLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AlipayAppLogic) AlipayApp() error {
	// todo: add your logic here and delete this line

	return nil
}
