package usecase

import (
	"github.com/go-playground/validator"

	"gorm.io/gorm"
)

type BaseUseCase struct {
	DB        *gorm.DB
	Validate  *validator.Validate
	TableName string
}
