package model

import (
	"context"

	"gorm.io/gorm"
)

type Address struct {
	// gorm.Model
	StreetAddress string `json:"street_address"`
	City          string `json:"city"`
	State         string `json:"state"`
	Country       string `json:"country"`
	ZipCode       string `json:"zip_code"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
}

type CreditCardInfo struct {
	// gorm.Model
	CardNumber          string `json:"card_number"`
	CVV                 int32  `json:"cvv"`
	CardExpirationYear  int32  `json:"card_expiration_year"`
	CardExpirationMonth int32  `json:"card_expiration_month"`
}
type PersonInfor struct {
	gorm.Model
	Address        Address        `gorm:"embedded"`
	CreditCardInfo CreditCardInfo `gorm:"embedded"`
	UserId         uint32
	Email          string
}

func SearchPersonInforByPersonId(ctx context.Context, DB *gorm.DB, PersonInforId uint32) (*PersonInfor, error) {
	var personInfor PersonInfor                                                         // 定义一个 PersonInfor 类型的变量
	err := DB.WithContext(ctx).Where("id = ?", PersonInforId).First(&personInfor).Error // 使用 &personInfor 作为指针传入
	if err != nil {
		return nil, err
	}
	return &personInfor, nil // 返回指向 personInfor 的指针
}

func SearchPersonInforByUserId(ctx context.Context, DB *gorm.DB, UserId uint32) ([]*PersonInfor, error) {
	var PersonInfors []*PersonInfor
	err := DB.WithContext(ctx).Where("user_id = ?", UserId).Find(&PersonInfors).Error
	if err != nil {
		return nil, err
	}
	return PersonInfors, nil
}

func Create(ctx context.Context, DB *gorm.DB, p *PersonInfor) error {
	return DB.WithContext(ctx).Create(p).Error
}

func Delete(ctx context.Context, DB *gorm.DB, PersonInforId uint32) error {
	return DB.WithContext(ctx).Where("id = ?", PersonInforId).Delete(&PersonInfor{}).Error
}

func Update(ctx context.Context, DB *gorm.DB, PersonInforId uint32, p *PersonInfor) error {
	return DB.WithContext(ctx).Where("id = ?", PersonInforId).Updates(p).Error
}
