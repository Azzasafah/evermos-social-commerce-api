package dto

type CategoryRequest struct {
	NamaCategory string `json:"nama_category" binding:"required"`
}
