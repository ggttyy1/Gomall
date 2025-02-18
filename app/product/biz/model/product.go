package model

import (
	"context"

	"gorm.io/gorm"
)

type Product struct {
	Base
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Picture     string     `json:"picture"`
	Price       float32    `json:"price"`
	Categories  []Category `json:"categories" gorm:"many2many:product_category"`
}

// 定义Product结构体的TableName方法，返回字符串"product"
func (p Product) TableName() string {
	return "product"
}

type ProductQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func (p ProductQuery) GetById(productID int) (product Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).First(&product, productID).Error
	return
}

func (p ProductQuery) SearchProduct(q string) (products []Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).Where("NAME LIKE ? or DESCRIPTION LIKE ?", "%"+q+"%", "%"+q+"%").Find(&products).Error

	return
}

func NewProductQuery(ctx context.Context, db *gorm.DB) *ProductQuery {
	return &ProductQuery{ctx: ctx, db: db}
}

func SearchCategory(ctx context.Context, db *gorm.DB, categoryType string) (category *Category, err error) {
	category = &Category{} // 创建一个 Category 指针
	err = db.WithContext(ctx).Model(&Category{}).Where("name = ?", categoryType).First(category).Error
	return
}

func CreateProduct(ctx context.Context, db *gorm.DB, product Product) error {
	return db.WithContext(ctx).Create(&product).Error
}
