package logic

import (
	"context"

	"github.com/tal-tech/go-zero/core/logx"
	"payment/internal/svc"
)

type AlipayWapLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAlipayWapLogic(ctx context.Context, svcCtx *svc.ServiceContext) AlipayWapLogic {
	return AlipayWapLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AlipayWapLogic) AlipayWap() error {
	// todo: add your logic here and delete this line

	return nil
}
