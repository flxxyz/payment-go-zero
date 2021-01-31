package payment

import (
	"payment/internal/model"
	"payment/internal/pay"

	"gorm.io/gorm"
)

type OrderPay struct {
	ID      uint        `gorm:"primaryKey"`
	OrderId string      `gorm:"column:orderId;type:varchar(30);unique"`
	TradeNo string      `gorm:"column:tradeNo;type:varchar(64);index"`
	Account string      `gorm:"column:account"`
	Amount  int64       `gorm:"column:amount"`
	IP      string      `gorm:"column:ip"`
	PayType pay.PayType `gorm:"column:payType;index"`
}

func (OrderPay) TableName() string {
	return model.Prefix("order_pay")
}

func (OrderPay) CreateOrderPay(op *OrderPay) *gorm.DB {
	return model.Orm().Create(op)
}

func (op *OrderPay) Update(updated map[string]interface{}) *gorm.DB {
	return model.Orm().Model(op).Updates(updated)
}
