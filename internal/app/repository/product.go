package repository

import (
	"github.com/google/uuid"
	"github.com/uhkrowi/go-simple-crud/internal/app/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetList() (products []model.Product, err error)
	GetSingle(productID uuid.UUID) (res model.Product, err error)
	Create(product model.Product) (res model.Product, err error)
	Update(product model.Product) (res model.Product, err error)
	Delete(productID uuid.UUID) (err error)
}

type ProductRepositoryImpl struct {
	BaseRepository
}

func NewProductRepository(DB *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{
		BaseRepository: BaseRepository{
			DB:        DB,
			TableName: "product",
		},
	}
}

func (r *ProductRepositoryImpl) GetList() (res []model.Product, err error) {
	err = r.DB.Table(r.TableName).
		Where("is_deleted = ?", false).
		Find(&res).Error

	return
}

func (r *ProductRepositoryImpl) GetSingle(productID uuid.UUID) (res model.Product, err error) {
	err = r.DB.Table(r.TableName).
		Where("id = ?", productID).
		Where("is_deleted = ?", false).
		First(&res).Error

	return
}

func (r *ProductRepositoryImpl) Create(product model.Product) (res model.Product, err error) {
	tx := r.DB.Begin()

	productID, err := uuid.NewRandom()
	if err != nil {
		return
	}

	err = tx.Exec(`INSERT INTO product(id,name) VALUES(?,?)`,
		productID,
		product.Name,
	).Error
	if err != nil {
		return
	}

	variantRepo := NewVariantRepository(r.DB)

	for i := range product.Variants {
		product.Variants[i].ProductID = productID

		err = variantRepo.CreateTx(tx, product.Variants[i])
		if err != nil {
			return
		}
	}

	tx.Commit()

	return
}

func (r *ProductRepositoryImpl) Update(product model.Product) (res model.Product, err error) {
	err = r.DB.Table(r.TableName).Model(&model.Product{}).
		Where("id = ?", product.ID).
		Updates(product).Error

	return
}

func (r *ProductRepositoryImpl) Delete(productID uuid.UUID) (err error) {
	_, err = r.GetSingle(productID)
	if err != nil {
		return
	}

	stmt := `UPDATE product SET is_deleted = ? WHERE id = ?`

	err = r.DB.Exec(stmt, true, productID).Error

	return
}
