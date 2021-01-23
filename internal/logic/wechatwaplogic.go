package logic

import (
	"context"

	"github.com/tal-tech/go-zero/core/logx"
	"payment/internal/svc"
)

type WechatWapLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWechatWapLogic(ctx context.Context, svcCtx *svc.ServiceContext) WechatWapLogic {
	return WechatWapLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WechatWapLogic) WechatWap() error {
	// todo: add your logic here and delete this line

	return nil
}
