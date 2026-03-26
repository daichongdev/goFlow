package repository

import (
	"goflow/internal/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	List(page, size int) ([]model.Product, int64, error)
	GetByID(id uint) (*model.Product, error)
	Create(product *model.Product) error
	Update(product *model.Product) error
	Delete(id uint) error
}

type ProductRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepository {
	return &ProductRepo{db: db}
}

// List 分页查询商品列表
func (r *ProductRepo) List(page, size int) ([]model.Product, int64, error) {
	var products []model.Product
	var total int64

	query := r.db.Model(&model.Product{})
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * size
	if err := query.Offset(offset).Limit(size).Order("id DESC").Find(&products).Error; err != nil {
		return nil, 0, err
	}
	return products, total, nil
}

// GetByID 根据 ID 查询商品
func (r *ProductRepo) GetByID(id uint) (*model.Product, error) {
	var product model.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

// Create 创建商品
func (r *ProductRepo) Create(product *model.Product) error {
	return r.db.Create(product).Error
}

// Update 更新商品
func (r *ProductRepo) Update(product *model.Product) error {
	return r.db.Save(product).Error
}

// Delete 删除商品（软删除）
func (r *ProductRepo) Delete(id uint) error {
	result := r.db.Delete(&model.Product{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
