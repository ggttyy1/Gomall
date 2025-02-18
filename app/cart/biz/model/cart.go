package model

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID    uint32 `gorm:"type:int(11);not null;index:idx_user_id"`
	ProductID uint32 `gorm:"type:int(11);not null"`
	Quantity  uint32 `gorm:"type:int(11);not null"`
}

func AddItem(ctx context.Context, db *gorm.DB, item *Cart) error {
	var row Cart
	err := db.WithContext(ctx).Model(&Cart{}).Where(&Cart{UserID: item.UserID, ProductID: item.ProductID}).First(&row).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if row.ID != 0 {
		// update
		row.Quantity += item.Quantity
		return db.WithContext(ctx).Model(&Cart{}).Where(&Cart{UserID: item.UserID, ProductID: item.ProductID}).Updates(&row).Error
	}
	return db.WithContext(ctx).Model(&Cart{}).Create(item).Error
}

func DeleteItem(ctx context.Context, db *gorm.DB, userid uint32) error {
	if userid == 0 {
		return errors.New("userid is empty")
	}
	return db.WithContext(ctx).Delete(&Cart{}, "user_id = ?", userid).Error

}

func SearchUserItem(ctx context.Context, db *gorm.DB, userid uint32) ([]*Cart, error) {
	var cart []*Cart
	err := db.WithContext(ctx).Model(&Cart{}).Where(&Cart{UserID: userid}).Find(&cart).Error
	return cart, err
}
