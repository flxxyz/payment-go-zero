package logic

import (
	"context"

	"github.com/tal-tech/go-zero/core/logx"
	"payment/internal/svc"
)

type WechatAppLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWechatAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) WechatAppLogic {
	return WechatAppLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WechatAppLogic) WechatApp() error {
	// todo: add your logic here and delete this line

	return nil
}
