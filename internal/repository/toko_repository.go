package repository

import (
	"evermos-backend/internal/model"

	"gorm.io/gorm"
)

type TokoRepository interface {
	Create(toko *model.Toko) error
	FindByUserID(userID uint) (*model.Toko, error)
	FindByID(id uint) (*model.Toko, error)
	Update(toko *model.Toko) error
	FindAll(page, limit int, nama string) ([]model.Toko, int64, error)
}

type tokoRepository struct {
	db *gorm.DB
}

func NewTokoRepository(db *gorm.DB) TokoRepository {
	return &tokoRepository{db: db}
}

func (r *tokoRepository) Create(toko *model.Toko) error {
	return r.db.Create(toko).Error
}

func (r *tokoRepository) FindByUserID(userID uint) (*model.Toko, error) {
	var toko model.Toko
	err := r.db.Where("id_user = ?", userID).First(&toko).Error
	if err != nil {
		return nil, err
	}
	return &toko, nil
}

func (r *tokoRepository) FindByID(id uint) (*model.Toko, error) {
	var toko model.Toko
	err := r.db.First(&toko, id).Error
	if err != nil {
		return nil, err
	}
	return &toko, nil
}

func (r *tokoRepository) Update(toko *model.Toko) error {
	return r.db.Save(toko).Error
}

func (r *tokoRepository) FindAll(page, limit int, nama string) ([]model.Toko, int64, error) {
	var tokos []model.Toko
	var total int64

	query := r.db.Model(&model.Toko{})
	if nama != "" {
		query = query.Where("nama_toko LIKE ?", "%"+nama+"%")
	}

	query.Count(&total)

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	err := query.Offset(offset).Limit(limit).Find(&tokos).Error
	if err != nil {
		return nil, 0, err
	}

	return tokos, total, nil
}
