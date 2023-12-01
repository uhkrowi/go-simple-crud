package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/uhkrowi/go-simple-crud/internal/app/model"
	"github.com/uhkrowi/go-simple-crud/internal/app/usecase"
	"github.com/uhkrowi/go-simple-crud/pkg/helper"
)

type ProductHandler interface {
	GetList(ctx *fiber.Ctx) error
	GetSingle(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type ProductHandlerImpl struct {
	uc usecase.ProductUseCase
}

func NewProductHandler(uc usecase.ProductUseCase) ProductHandler {
	return &ProductHandlerImpl{
		uc: uc,
	}
}

func (h *ProductHandlerImpl) GetList(ctx *fiber.Ctx) (err error) {
	filter := model.ProductReqFilter{}
	err = ctx.QueryParser(&filter)
	if err != nil {
		return helper.GetWebResponse(ctx, nil, err)
	}

	result, err := h.uc.GetList(ctx, filter)
	if err != nil {
		return helper.GetWebResponse(ctx, nil, err)
	}

	return helper.GetWebResponse(ctx, result, err)
}

func (h *ProductHandlerImpl) GetSingle(ctx *fiber.Ctx) (err error) {
	productID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return
	}

	result, err := h.uc.GetSingle(ctx, productID)
	if err != nil {
		return helper.GetWebResponse(ctx, nil, err)
	}

	return helper.GetWebResponse(ctx, result, err)
}

func (h *ProductHandlerImpl) Create(ctx *fiber.Ctx) (err error) {
	var product model.Product

	err = ctx.BodyParser(&product)
	if err != nil {
		return helper.GetWebResponse(ctx, nil, err)
	}

	err = h.uc.Create(ctx, product)

	return helper.GetWebResponse(ctx, nil, err)
}

func (h *ProductHandlerImpl) Update(ctx *fiber.Ctx) (err error) {
	var product model.Product

	product.ID, err = uuid.Parse(ctx.Params("id"))
	if err != nil {
		return
	}

	err = ctx.BodyParser(&product)
	if err != nil {
		return helper.GetWebResponse(ctx, nil, err)
	}

	result, err := h.uc.Update(ctx, product)

	return helper.GetWebResponse(ctx, result, err)
}

func (h *ProductHandlerImpl) Delete(ctx *fiber.Ctx) (err error) {
	productID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return
	}

	err = h.uc.Delete(ctx, productID)

	return helper.GetWebResponse(ctx, nil, err)
}
