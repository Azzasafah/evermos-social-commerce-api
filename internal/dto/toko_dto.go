package dto

import "evermos-backend/internal/model"

type UpdateTokoRequest struct {
	NamaToko string `form:"nama_toko"`
}

type TokoPaginationResponse struct {
	Page  int          `json:"page"`
	Limit int          `json:"limit"`
	Data  []model.Toko `json:"data"`
}
