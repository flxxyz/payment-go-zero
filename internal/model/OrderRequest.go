package payment

import (
	"payment/internal/model"

	"gorm.io/gorm"
)

type OrderRequest struct {
	ID      uint   `gorm:"primaryKey"`
	OrderId string `gorm:"column:orderId;type:varchar(30);unique"`
	TradeNo string `gorm:"column:tradeNo;type:varchar(64);index"`
}

func (OrderRequest) TableName() string {
	return model.Prefix("order_request")
}

func (OrderRequest) Create(orderId string) (or OrderRequest, db *gorm.DB) {
	or.OrderId = orderId
	db = model.Orm().Select("orderId").Create(&or)
	return
}

func (OrderRequest) FindByOrderId(orderId string) (or OrderRequest, db *gorm.DB) {
	db = model.Orm().Where("orderId = ?", orderId).First(&or)
	return
}

func (or *OrderRequest) UpdateOrderRequestTradeNo(tradeNo string) *gorm.DB {
	return model.Orm().Model(or).Update("tradeNo", tradeNo)
}
