package service

import (
	"errors"

	"evermos-backend/internal/dto"
	"evermos-backend/internal/model"
	"evermos-backend/internal/repository"
)

type AlamatService interface {
	GetMyAlamat(userID uint) ([]model.Alamat, error)
	GetAlamatByID(id, userID uint) (*model.Alamat, error)
	CreateAlamat(userID uint, req *dto.CreateAlamatRequest) (uint, error)
	UpdateAlamat(id, userID uint, req *dto.UpdateAlamatRequest) error
	DeleteAlamat(id, userID uint) error
}

type alamatService struct {
	alamatRepo repository.AlamatRepository
}

func NewAlamatService(alamatRepo repository.AlamatRepository) AlamatService {
	return &alamatService{alamatRepo: alamatRepo}
}

func (s *alamatService) GetMyAlamat(userID uint) ([]model.Alamat, error) {
	alamats, err := s.alamatRepo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}
	if alamats == nil {
		alamats = []model.Alamat{}
	}
	return alamats, nil
}

func (s *alamatService) GetAlamatByID(id, userID uint) (*model.Alamat, error) {
	alamat, err := s.alamatRepo.FindByIDAndUserID(id, userID)
	if err != nil {
		return nil, errors.New("record not found")
	}
	return alamat, nil
}

func (s *alamatService) CreateAlamat(userID uint, req *dto.CreateAlamatRequest) (uint, error) {
	alamat := model.Alamat{
		IdUser:       userID,
		JudulAlamat:  req.JudulAlamat,
		NamaPenerima: req.NamaPenerima,
		NoTelp:       req.NoTelp,
		DetailAlamat: req.DetailAlamat,
	}

	if err := s.alamatRepo.Create(&alamat); err != nil {
		return 0, err
	}
	return alamat.ID, nil
}

func (s *alamatService) UpdateAlamat(id, userID uint, req *dto.UpdateAlamatRequest) error {
	alamat, err := s.alamatRepo.FindByIDAndUserID(id, userID)
	if err != nil {
		return errors.New("record not found")
	}

	if req.JudulAlamat != "" {
		alamat.JudulAlamat = req.JudulAlamat
	}
	if req.NamaPenerima != "" {
		alamat.NamaPenerima = req.NamaPenerima
	}
	if req.NoTelp != "" {
		alamat.NoTelp = req.NoTelp
	}
	if req.DetailAlamat != "" {
		alamat.DetailAlamat = req.DetailAlamat
	}

	return s.alamatRepo.Update(alamat)
}

func (s *alamatService) DeleteAlamat(id, userID uint) error {
	return s.alamatRepo.Delete(id, userID)
}
