package repository

import (
	"errors"
	"fmt"

	"evermos-backend/internal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
	FindByNoTelp(noTelp string) (*model.User, error)
	FindByID(id uint) (*model.User, error)
	Update(user *model.User) error
	CheckEmailOrPhoneExists(email, noTelp string, excludeUserID uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByNoTelp(noTelp string) (*model.User, error) {
	var user model.User
	err := r.db.Where("notelp = ?", noTelp).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByID(id uint) (*model.User, error) {
	var user model.User
	err := r.db.Preload("Toko").First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) CheckEmailOrPhoneExists(email, noTelp string, excludeUserID uint) error {
	var count int64
	query := r.db.Model(&model.User{}).Where("email = ? AND id != ?", email, excludeUserID)
	query.Count(&count)
	if count > 0 {
		return fmt.Errorf("Error 1062: Duplicate entry '%s' for key 'users.email'", email)
	}

	query = r.db.Model(&model.User{}).Where("notelp = ? AND id != ?", noTelp, excludeUserID)
	query.Count(&count)
	if count > 0 {
		return errors.New("No telepon sudah terdaftar")
	}

	return nil
}
