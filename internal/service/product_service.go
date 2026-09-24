package service

import (
	"errors"
	"mime/multipart"

	"evermos-backend/internal/dto"
	"evermos-backend/internal/helper"
	"evermos-backend/internal/model"
	"evermos-backend/internal/repository"
)

type ProductService interface {
	GetAllProducts(page, limit int, nama, categoryID string, minHarga, maxHarga int) (*dto.ProductPaginationResponse, error)
	GetProductByID(id uint) (*model.Produk, error)
	CreateProduct(userID uint, req *dto.CreateProductRequest, photos []*multipart.FileHeader) (uint, error)
	UpdateProduct(id, userID uint, req *dto.UpdateProductRequest, photos []*multipart.FileHeader) error
	DeleteProduct(id, userID uint) error
}

type productService struct {
	productRepo repository.ProductRepository
	tokoRepo    repository.TokoRepository
}

func NewProductService(productRepo repository.ProductRepository, tokoRepo repository.TokoRepository) ProductService {
	return &productService{
		productRepo: productRepo,
		tokoRepo:    tokoRepo,
	}
}

func (s *productService) GetAllProducts(page, limit int, nama, categoryID string, minHarga, maxHarga int) (*dto.ProductPaginationResponse, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	products, _, err := s.productRepo.FindAll(page, limit, nama, categoryID, minHarga, maxHarga)
	if err != nil {
		return nil, err
	}

	if products == nil {
		products = []model.Produk{}
	}

	return &dto.ProductPaginationResponse{
		Data:  products,
		Page:  page,
		Limit: limit,
	}, nil
}

func (s *productService) GetProductByID(id uint) (*model.Produk, error) {
	product, err := s.productRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("No Data Product")
	}
	return product, nil
}

func (s *productService) CreateProduct(userID uint, req *dto.CreateProductRequest, photos []*multipart.FileHeader) (uint, error) {
	toko, err := s.tokoRepo.FindByUserID(userID)
	if err != nil {
		return 0, errors.New("Toko tidak ditemukan untuk user ini")
	}

	slug := helper.GenerateSlug(req.NamaProduk)

	product := model.Produk{
		NamaProduk:    req.NamaProduk,
		Slug:          slug,
		HargaReseller: req.HargaReseller,
		HargaKonsumen: req.HargaKonsumen,
		Stok:          req.Stok,
		Deskripsi:     req.Deskripsi,
		IdToko:        toko.ID,
		IdCategory:    req.CategoryID,
	}

	if err := s.productRepo.Create(&product); err != nil {
		return 0, err
	}

	// Save photos
	var fotoList []model.FotoProduk
	for _, p := range photos {
		filename, err := helper.SaveUploadedFile(p, "./uploads")
		if err == nil {
			fotoList = append(fotoList, model.FotoProduk{
				IdProduk: product.ID,
				Url:      filename,
			})
		}
	}

	if len(fotoList) > 0 {
		_ = s.productRepo.CreatePhotos(fotoList)
	}

	return product.ID, nil
}

func (s *productService) UpdateProduct(id, userID uint, req *dto.UpdateProductRequest, photos []*multipart.FileHeader) error {
	toko, err := s.tokoRepo.FindByUserID(userID)
	if err != nil {
		return errors.New("Toko tidak ditemukan")
	}

	// Requirement 14: User tidak dapat mengelola data product dari data user lain
	product, err := s.productRepo.FindByIDAndTokoID(id, toko.ID)
	if err != nil {
		return errors.New("record not found")
	}

	if req.NamaProduk != "" {
		product.NamaProduk = req.NamaProduk
		product.Slug = helper.GenerateSlug(req.NamaProduk)
	}
	if req.CategoryID > 0 {
		product.IdCategory = req.CategoryID
	}
	if req.HargaReseller > 0 {
		product.HargaReseller = req.HargaReseller
	}
	if req.HargaKonsumen > 0 {
		product.HargaKonsumen = req.HargaKonsumen
	}
	if req.Stok >= 0 {
		product.Stok = req.Stok
	}
	if req.Deskripsi != "" {
		product.Deskripsi = req.Deskripsi
	}

	if err := s.productRepo.Update(product); err != nil {
		return err
	}

	// If new photos provided, add them
	var fotoList []model.FotoProduk
	for _, p := range photos {
		filename, err := helper.SaveUploadedFile(p, "./uploads")
		if err == nil {
			fotoList = append(fotoList, model.FotoProduk{
				IdProduk: product.ID,
				Url:      filename,
			})
		}
	}
	if len(fotoList) > 0 {
		_ = s.productRepo.CreatePhotos(fotoList)
	}

	return nil
}

func (s *productService) DeleteProduct(id, userID uint) error {
	toko, err := s.tokoRepo.FindByUserID(userID)
	if err != nil {
		return errors.New("record not found")
	}

	// Requirement 14: User tidak dapat mengelola data product dari data user lain
	return s.productRepo.Delete(id, toko.ID)
}
