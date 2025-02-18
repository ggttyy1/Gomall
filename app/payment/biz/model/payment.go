package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type PaymentLog struct {
	gorm.Model
	UserId        uint32    `json:"user_id"`
	OrderId       string    `json:"order_id"`
	TransactionId string    `json:"transaction_id"`
	Amount        float32   `json:"amount"`
	PayTime       time.Time `json:"pay_time"`
}

func (PaymentLog) TableName() string {
	return "payment_log"
}

func CreatePaymentLog(ctx context.Context, db *gorm.DB, log *PaymentLog) error {
	return db.WithContext(ctx).Model(&PaymentLog{}).Create(log).Error
}
