package service

import (
	"errors"
	"mime/multipart"

	"evermos-backend/internal/dto"
	"evermos-backend/internal/helper"
	"evermos-backend/internal/model"
	"evermos-backend/internal/repository"
)

type TokoService interface {
	GetMyToko(userID uint) (*model.Toko, error)
	UpdateToko(userID, tokoID uint, namaToko string, photoFile *multipart.FileHeader) error
	GetTokoByID(id uint) (*model.Toko, error)
	GetAllToko(page, limit int, nama string) (*dto.TokoPaginationResponse, error)
}

type tokoService struct {
	tokoRepo repository.TokoRepository
}

func NewTokoService(tokoRepo repository.TokoRepository) TokoService {
	return &tokoService{tokoRepo: tokoRepo}
}

func (s *tokoService) GetMyToko(userID uint) (*model.Toko, error) {
	return s.tokoRepo.FindByUserID(userID)
}

func (s *tokoService) UpdateToko(userID, tokoID uint, namaToko string, photoFile *multipart.FileHeader) error {
	toko, err := s.tokoRepo.FindByID(tokoID)
	if err != nil {
		return errors.New("Toko tidak ditemukan")
	}

	// Requirement 13: User tidak dapat mengelola data toko dari data user lain
	if toko.IdUser != userID {
		return errors.New("Akses ditolak: Anda bukan pemilik toko ini")
	}

	if namaToko != "" {
		toko.NamaToko = namaToko
	}

	if photoFile != nil {
		filename, err := helper.SaveUploadedFile(photoFile, "./uploads")
		if err != nil {
			return err
		}
		toko.UrlFoto = filename
	}

	return s.tokoRepo.Update(toko)
}

func (s *tokoService) GetTokoByID(id uint) (*model.Toko, error) {
	toko, err := s.tokoRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("Toko tidak ditemukan")
	}
	return toko, nil
}

func (s *tokoService) GetAllToko(page, limit int, nama string) (*dto.TokoPaginationResponse, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	tokos, _, err := s.tokoRepo.FindAll(page, limit, nama)
	if err != nil {
		return nil, err
	}

	if tokos == nil {
		tokos = []model.Toko{}
	}

	return &dto.TokoPaginationResponse{
		Page:  page,
		Limit: limit,
		Data:  tokos,
	}, nil
}
