package dto

import "evermos-backend/internal/model"

type CreateProductRequest struct {
	NamaProduk    string `form:"nama_produk" binding:"required"`
	CategoryID    uint   `form:"category_id" binding:"required"`
	HargaReseller int    `form:"harga_reseller" binding:"required"`
	HargaKonsumen int    `form:"harga_konsumen" binding:"required"`
	Stok          int    `form:"stok" binding:"required"`
	Deskripsi     string `form:"deskripsi"`
}

type UpdateProductRequest struct {
	NamaProduk    string `form:"nama_produk"`
	CategoryID    uint   `form:"category_id"`
	HargaReseller int    `form:"harga_reseller"`
	HargaKonsumen int    `form:"harga_konsumen"`
	Stok          int    `form:"stok"`
	Deskripsi     string `form:"deskripsi"`
}

type ProductPaginationResponse struct {
	Data  []model.Produk `json:"data"`
	Page  int            `json:"page"`
	Limit int            `json:"limit"`
}
