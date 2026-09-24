package dto

import "evermos-backend/internal/model"

type CreateTrxItemRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	Kuantitas int  `json:"kuantitas" binding:"required,min=1"`
}

type CreateTrxRequest struct {
	MethodBayar string                 `json:"method_bayar" binding:"required"`
	AlamatKirim uint                   `json:"alamat_kirim" binding:"required"`
	DetailTrx   []CreateTrxItemRequest `json:"detail_trx" binding:"required,min=1"`
}

type TrxPaginationResponse struct {
	Data  []model.Trx `json:"data"`
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
}
