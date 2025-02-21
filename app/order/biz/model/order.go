package model

import (
	"context"

	"gorm.io/gorm"
)

type Consignee struct {
	Email         string
	StreetAddress string
	City          string
	State         string
	Country       string
	ZipCode       string
}

type Order struct {
	gorm.Model
	UserId    uint32
	OrderId   string       `gorm:"type:varchar(100);uniqueIndex"`
	Consignee Consignee    `gorm:"embedded"`
	Items     []OrderItems `gorm:"foreignKey:OrderIdRefer;references:OrderId"`
}

type OrderItems struct {
	gorm.Model
	ProductId    uint32
	Qty          uint32
	Cost         float32
	OrderIdRefer string `gorm:"type:varchar(100);index"`
}

func ListOrder(ctx context.Context, db *gorm.DB, userId uint32) ([]*Order, error) {
	var orders []*Order
	err := db.WithContext(ctx).Unscoped().Where("user_id = ?", userId).Preload("Items").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func CreateOrder(ctx context.Context, db *gorm.DB, order *Order) error {
	return db.WithContext(ctx).Create(order).Error
}

func DeleteOrder(ctx context.Context, db *gorm.DB, orderId string) error {
	return db.WithContext(ctx).Where("order_id = ?", orderId).Delete(&Order{}).Error
}
