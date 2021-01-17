package logic

import (
	"context"

	"payment/internal/svc"
	"payment/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type PaymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) PaymentLogic {
	return PaymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PaymentLogic) Payment(req types.Request) (types.Response, error) {
	// todo: add your logic here and delete this line

	return types.Response{}, nil
}
