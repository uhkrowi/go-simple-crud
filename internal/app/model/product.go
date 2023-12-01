package model

import (
	"github.com/google/uuid"
)

type ProductReqFilter struct{}

type Product struct {
	ID       uuid.UUID `gorm:"primaryKey;column:id" json:"id" swaggerignore:"true"`
	Name     string    `gorm:"column:name" json:"name"`
	IsActive bool      `gorm:"column:is_active" json:"is_active"`
	Variants []Variant `gorm:"-" json:"variants"`
}
