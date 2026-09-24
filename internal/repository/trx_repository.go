package repository

import (
	"evermos-backend/internal/model"

	"gorm.io/gorm"
)

type TrxRepository interface {
	CreateTrxWithDetails(trx *model.Trx, details []model.DetailTrx, logs []model.LogProduk, stockUpdates map[uint]int) error
	FindByUserID(userID uint, page, limit int) ([]model.Trx, int64, error)
	FindByIDAndUserID(id, userID uint) (*model.Trx, error)
}

type trxRepository struct {
	db *gorm.DB
}

func NewTrxRepository(db *gorm.DB) TrxRepository {
	return &trxRepository{db: db}
}

func (r *trxRepository) CreateTrxWithDetails(
	trx *model.Trx,
	details []model.DetailTrx,
	logs []model.LogProduk,
	stockUpdates map[uint]int,
) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// 1. Create Trx header
		if err := tx.Create(trx).Error; err != nil {
			return err
		}

		// 2. Create LogProduk entries & DetailTrx entries
		for i := range logs {
			if err := tx.Create(&logs[i]).Error; err != nil {
				return err
			}
			details[i].IdTrx = trx.ID
			details[i].IdLogProduk = logs[i].ID

			if err := tx.Create(&details[i]).Error; err != nil {
				return err
			}
		}

		// 3. Deduct stock for each product
		for prodID, newStock := range stockUpdates {
			if err := tx.Model(&model.Produk{}).Where("id = ?", prodID).Update("stok", newStock).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *trxRepository) FindByUserID(userID uint, page, limit int) ([]model.Trx, int64, error) {
	var trxes []model.Trx
	var total int64

	query := r.db.Model(&model.Trx{}).Where("id_user = ?", userID)
	query.Count(&total)

	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		query = query.Offset(offset).Limit(limit)
	}

	err := query.Preload("AlamatKirim").
		Preload("DetailTrx").
		Preload("DetailTrx.Toko").
		Preload("DetailTrx.Product").
		Preload("DetailTrx.Product.Category").
		Preload("DetailTrx.Product.Toko").
		Order("created_at desc").
		Find(&trxes).Error

	if err != nil {
		return nil, 0, err
	}

	// Populate photos for each LogProduk
	r.populateLogPhotos(trxes)

	return trxes, total, nil
}

func (r *trxRepository) FindByIDAndUserID(id, userID uint) (*model.Trx, error) {
	var trx model.Trx
	err := r.db.Where("id = ? AND id_user = ?", id, userID).
		Preload("AlamatKirim").
		Preload("DetailTrx").
		Preload("DetailTrx.Toko").
		Preload("DetailTrx.Product").
		Preload("DetailTrx.Product.Category").
		Preload("DetailTrx.Product.Toko").
		First(&trx).Error

	if err != nil {
		return nil, err
	}

	// Populate photos for each LogProduk
	r.populateLogPhotos([]model.Trx{trx})

	return &trx, nil
}

func (r *trxRepository) populateLogPhotos(trxes []model.Trx) {
	for i := range trxes {
		for j := range trxes[i].DetailTrx {
			if trxes[i].DetailTrx[j].Product != nil {
				var photos []model.FotoProduk
				r.db.Where("id_produk = ?", trxes[i].DetailTrx[j].Product.IdProduk).Find(&photos)
				trxes[i].DetailTrx[j].Product.Photos = photos
			}
		}
	}
}
