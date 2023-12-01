package repository

import (
	"github.com/google/uuid"
	"github.com/uhkrowi/go-simple-crud/internal/app/model"
	"gorm.io/gorm"
)

type VariantRepository interface {
	GetList(filter model.VariantReqFilter) (res []model.Variant, err error)
	CreateTx(tx *gorm.DB, variant model.Variant) (err error)
}

type VariantRepositoryImpl struct {
	BaseRepository
}

func NewVariantRepository(DB *gorm.DB) VariantRepository {
	return &VariantRepositoryImpl{
		BaseRepository: BaseRepository{
			DB: DB,
		},
	}
}

func (r *VariantRepositoryImpl) GetList(filter model.VariantReqFilter) (res []model.Variant, err error) {
	stmt := `select
		v.id,
		v.product_id,
		v.name,
		p.name as product_name,
		v.price,
		v.stock 
	from
		variant v
	join product p on
		p.id = v.product_id
	where v.is_deleted is false`

	placeholders := []any{}

	if filter.ProductID != nil {
		stmt += ` and v.product_id = ?`
		placeholders = append(placeholders, filter.ProductID)
	}

	err = r.DB.Raw(stmt, placeholders...).Find(&res).Error

	return
}

func (r *VariantRepositoryImpl) CreateTx(tx *gorm.DB, variant model.Variant) (err error) {
	ID, err := uuid.NewRandom()
	if err != nil {
		return
	}

	err = tx.Exec(`INSERT INTO variant(id,product_id,name,price,stock) VALUES(?,?,?,?,?)`,
		ID,
		variant.ProductID,
		variant.Name,
		variant.Price,
		variant.Stock,
	).Error

	return
}
