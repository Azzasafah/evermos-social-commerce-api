package repository

import (
	"evermos-backend/internal/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *model.Produk) error
	FindByID(id uint) (*model.Produk, error)
	FindByIDAndTokoID(id, tokoID uint) (*model.Produk, error)
	FindAll(page, limit int, nama, categoryID string, minHarga, maxHarga int) ([]model.Produk, int64, error)
	Update(product *model.Produk) error
	Delete(id, tokoID uint) error
	CreatePhotos(photos []model.FotoProduk) error
	DeletePhotosByProductID(productID uint) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(product *model.Produk) error {
	return r.db.Create(product).Error
}

func (r *productRepository) FindByID(id uint) (*model.Produk, error) {
	var product model.Produk
	err := r.db.Preload("Toko").Preload("Category").Preload("Photos").First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) FindByIDAndTokoID(id, tokoID uint) (*model.Produk, error) {
	var product model.Produk
	err := r.db.Where("id = ? AND id_toko = ?", id, tokoID).Preload("Toko").Preload("Category").Preload("Photos").First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) FindAll(page, limit int, nama, categoryID string, minHarga, maxHarga int) ([]model.Produk, int64, error) {
	var products []model.Produk
	var total int64

	query := r.db.Model(&model.Produk{})

	if nama != "" {
		query = query.Where("nama_produk LIKE ?", "%"+nama+"%")
	}
	if categoryID != "" {
		query = query.Where("id_category = ?", categoryID)
	}
	if minHarga > 0 {
		query = query.Where("harga_konsumen >= ?", minHarga)
	}
	if maxHarga > 0 {
		query = query.Where("harga_konsumen <= ?", maxHarga)
	}

	query.Count(&total)

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	err := query.Preload("Toko").Preload("Category").Preload("Photos").Offset(offset).Limit(limit).Find(&products).Error
	if err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

func (r *productRepository) Update(product *model.Produk) error {
	return r.db.Save(product).Error
}

func (r *productRepository) Delete(id, tokoID uint) error {
	result := r.db.Where("id = ? AND id_toko = ?", id, tokoID).Delete(&model.Produk{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	// Also delete photos
	r.db.Where("id_produk = ?", id).Delete(&model.FotoProduk{})
	return nil
}

func (r *productRepository) CreatePhotos(photos []model.FotoProduk) error {
	if len(photos) == 0 {
		return nil
	}
	return r.db.Create(&photos).Error
}

func (r *productRepository) DeletePhotosByProductID(productID uint) error {
	return r.db.Where("id_produk = ?", productID).Delete(&model.FotoProduk{}).Error
}
