package repository

import (
	"evermos-backend/internal/model"

	"gorm.io/gorm"
)

type AlamatRepository interface {
	Create(alamat *model.Alamat) error
	FindByUserID(userID uint) ([]model.Alamat, error)
	FindByIDAndUserID(id, userID uint) (*model.Alamat, error)
	FindByID(id uint) (*model.Alamat, error)
	Update(alamat *model.Alamat) error
	Delete(id, userID uint) error
}

type alamatRepository struct {
	db *gorm.DB
}

func NewAlamatRepository(db *gorm.DB) AlamatRepository {
	return &alamatRepository{db: db}
}

func (r *alamatRepository) Create(alamat *model.Alamat) error {
	return r.db.Create(alamat).Error
}

func (r *alamatRepository) FindByUserID(userID uint) ([]model.Alamat, error) {
	var alamats []model.Alamat
	err := r.db.Where("id_user = ?", userID).Find(&alamats).Error
	return alamats, err
}

func (r *alamatRepository) FindByIDAndUserID(id, userID uint) (*model.Alamat, error) {
	var alamat model.Alamat
	err := r.db.Where("id = ? AND id_user = ?", id, userID).First(&alamat).Error
	if err != nil {
		return nil, err
	}
	return &alamat, nil
}

func (r *alamatRepository) FindByID(id uint) (*model.Alamat, error) {
	var alamat model.Alamat
	err := r.db.First(&alamat, id).Error
	if err != nil {
		return nil, err
	}
	return &alamat, nil
}

func (r *alamatRepository) Update(alamat *model.Alamat) error {
	return r.db.Save(alamat).Error
}

func (r *alamatRepository) Delete(id, userID uint) error {
	result := r.db.Where("id = ? AND id_user = ?", id, userID).Delete(&model.Alamat{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
