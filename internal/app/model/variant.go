package model

import (
	"github.com/google/uuid"
)

type VariantReqFilter struct {
	ProductID *uuid.UUID `query:"product_id"`
}

type Variant struct {
	ID          uuid.UUID `gorm:"primaryKey;column:id" json:"id" search:"id" swaggerignore:"true"`
	Name        string    `gorm:"column:name" json:"name"`
	ProductID   uuid.UUID `gorm:"column:product_id" json:"product_id"`
	ProductName string    `gorm:"column:product_name" json:"product_name" `
	Price       float64   `gorm:"column:price" json:"price" `
	Stock       int       `gorm:"column:stock" json:"stock"`
}
