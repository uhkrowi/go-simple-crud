package usecase

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/uhkrowi/go-simple-crud/internal/app/model"
	"github.com/uhkrowi/go-simple-crud/internal/app/repository"
	"gorm.io/gorm"
)

type ProductUseCase interface {
	GetList(ctx *fiber.Ctx, filter model.ProductReqFilter) (res []model.Product, err error)
	GetSingle(ctx *fiber.Ctx, productID uuid.UUID) (res model.Product, err error)
	Create(ctx *fiber.Ctx, product model.Product) (err error)
	Update(ctx *fiber.Ctx, product model.Product) (res model.Product, err error)
	Delete(ctx *fiber.Ctx, productID uuid.UUID) (err error)
}

type ProductUseCaseImpl struct {
	BaseUseCase
}

func NewProductUseCase(DB *gorm.DB, validate *validator.Validate) ProductUseCase {
	return &ProductUseCaseImpl{
		BaseUseCase: BaseUseCase{
			DB:       DB,
			Validate: validate,
		},
	}
}

func (uc *ProductUseCaseImpl) GetList(ctx *fiber.Ctx, filter model.ProductReqFilter) (res []model.Product, err error) {
	productRepo := repository.NewProductRepository(uc.DB)
	variantRepo := repository.NewVariantRepository(uc.DB)

	res, err = productRepo.GetList()
	if err != nil {
		return
	}

	for i := range res {
		res[i].Variants, err = variantRepo.GetList(model.VariantReqFilter{ProductID: &res[i].ID})
		if err != nil {
			return
		}
	}

	return
}

func (uc *ProductUseCaseImpl) GetSingle(ctx *fiber.Ctx, productID uuid.UUID) (res model.Product, err error) {
	productRepo := repository.NewProductRepository(uc.DB)

	res, err = productRepo.GetSingle(productID)
	if err != nil {
		return
	}

	variantRepo := repository.NewVariantRepository(uc.DB)

	res.Variants, err = variantRepo.GetList(model.VariantReqFilter{ProductID: &productID})
	if err != nil {
		return
	}

	return
}

func (uc *ProductUseCaseImpl) Create(ctx *fiber.Ctx, product model.Product) (err error) {
	productRepo := repository.NewProductRepository(uc.DB)

	_, err = productRepo.Create(product)

	return
}

func (uc *ProductUseCaseImpl) Update(ctx *fiber.Ctx, product model.Product) (res model.Product, err error) {
	productRepo := repository.NewProductRepository(uc.DB)

	_, err = productRepo.Update(product)
	if err != nil {
		return
	}

	res, err = uc.GetSingle(ctx, product.ID)

	return
}

func (uc *ProductUseCaseImpl) Delete(ctx *fiber.Ctx, productID uuid.UUID) (err error) {
	productRepo := repository.NewProductRepository(uc.DB)

	err = productRepo.Delete(productID)

	return
}
