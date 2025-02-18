package model

import (
	"context"

	"gorm.io/gorm"
)

type Category struct {
	Base
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Products    []Product `json:"product" gorm:"many2many:product_category"`
}

func (c Category) TableName() string {
	return "category"
}

type CatagoryQuery struct {
	ctx context.Context
	db  *gorm.DB
}

// Find函数用于查找Category类型的切片
func (c CatagoryQuery) GetProductsByCategoryName(name string) (categories []Category, err error) {
	err = c.db.WithContext(c.ctx).Model(&Category{}).Where(&Category{Name: name}).Preload("Products").Find(&categories).Error
	return
}

func NewCateGoryQuery(ctx context.Context, db *gorm.DB) *CatagoryQuery {
	return &CatagoryQuery{ctx: ctx, db: db}
}
