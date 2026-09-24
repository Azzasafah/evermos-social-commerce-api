package service

import (
	"errors"
	"fmt"
	"time"

	"evermos-backend/internal/dto"
	"evermos-backend/internal/model"
	"evermos-backend/internal/repository"
)

type TrxService interface {
	CreateTrx(userID uint, req *dto.CreateTrxRequest) (uint, error)
	GetAllTrx(userID uint, page, limit int) (*dto.TrxPaginationResponse, error)
	GetTrxByID(id, userID uint) (*model.Trx, error)
}

type trxService struct {
	trxRepo     repository.TrxRepository
	alamatRepo  repository.AlamatRepository
	productRepo repository.ProductRepository
}

func NewTrxService(
	trxRepo repository.TrxRepository,
	alamatRepo repository.AlamatRepository,
	productRepo repository.ProductRepository,
) TrxService {
	return &trxService{
		trxRepo:     trxRepo,
		alamatRepo:  alamatRepo,
		productRepo: productRepo,
	}
}

func (s *trxService) CreateTrx(userID uint, req *dto.CreateTrxRequest) (uint, error) {
	// 1. Verify shipping address belongs to user
	alamat, err := s.alamatRepo.FindByIDAndUserID(req.AlamatKirim, userID)
	if err != nil || alamat == nil {
		return 0, errors.New("Alamat pengiriman tidak ditemukan atau bukan milik Anda")
	}

	var logs []model.LogProduk
	var details []model.DetailTrx
	stockUpdates := make(map[uint]int)
	totalTrx := 0

	// 2. Validate products and calculate totals
	for _, item := range req.DetailTrx {
		product, err := s.productRepo.FindByID(item.ProductID)
		if err != nil || product == nil {
			return 0, fmt.Errorf("Produk dengan ID %d tidak ditemukan", item.ProductID)
		}

		currentStock := product.Stok
		if updated, exists := stockUpdates[product.ID]; exists {
			currentStock = updated
		}

		if currentStock < item.Kuantitas {
			return 0, fmt.Errorf("Stok tidak mencukupi untuk produk: %s", product.NamaProduk)
		}

		stockUpdates[product.ID] = currentStock - item.Kuantitas

		itemSubtotal := product.HargaKonsumen * item.Kuantitas
		totalTrx += itemSubtotal

		// Requirement 16 & 17: Tabel log product diisi ketika melakukan transaksi
		logEntry := model.LogProduk{
			IdProduk:      product.ID,
			NamaProduk:    product.NamaProduk,
			Slug:          product.Slug,
			HargaReseller: product.HargaReseller,
			HargaKonsumen: product.HargaKonsumen,
			Deskripsi:     product.Deskripsi,
			IdToko:        product.IdToko,
			IdCategory:    product.IdCategory,
		}
		logs = append(logs, logEntry)

		detailEntry := model.DetailTrx{
			IdToko:     product.IdToko,
			Kuantitas:  item.Kuantitas,
			HargaTotal: itemSubtotal,
		}
		details = append(details, detailEntry)
	}

	// 3. Create Trx header
	invoiceCode := fmt.Sprintf("INV-%d", time.Now().Unix())
	trx := model.Trx{
		IdUser:           userID,
		AlamatPengiriman: alamat.ID,
		HargaTotal:       totalTrx,
		KodeInvoice:      invoiceCode,
		MethodBayar:      req.MethodBayar,
	}

	// 4. Save via atomic repository transaction
	if err := s.trxRepo.CreateTrxWithDetails(&trx, details, logs, stockUpdates); err != nil {
		return 0, err
	}

	return trx.ID, nil
}

func (s *trxService) GetAllTrx(userID uint, page, limit int) (*dto.TrxPaginationResponse, error) {
	trxes, _, err := s.trxRepo.FindByUserID(userID, page, limit)
	if err != nil {
		return nil, err
	}

	if trxes == nil {
		trxes = []model.Trx{}
	}

	return &dto.TrxPaginationResponse{
		Data:  trxes,
		Page:  page,
		Limit: limit,
	}, nil
}

func (s *trxService) GetTrxByID(id, userID uint) (*model.Trx, error) {
	trx, err := s.trxRepo.FindByIDAndUserID(id, userID)
	if err != nil {
		return nil, errors.New("No Data Trx")
	}
	return trx, nil
}
