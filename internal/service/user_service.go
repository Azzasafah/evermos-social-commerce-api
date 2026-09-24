package service

import (
	"evermos-backend/internal/dto"
	"evermos-backend/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetProfile(userID uint) (*dto.UserProfileResponse, error)
	UpdateProfile(userID uint, req *dto.UpdateProfileRequest) error
}

type userService struct {
	userRepo        repository.UserRepository
	provCityService ProvCityService
}

func NewUserService(userRepo repository.UserRepository, provCityService ProvCityService) UserService {
	return &userService{
		userRepo:        userRepo,
		provCityService: provCityService,
	}
}

func (s *userService) GetProfile(userID uint) (*dto.UserProfileResponse, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	prov, _ := s.provCityService.GetDetailProvince(user.IdProvinsi)
	city, _ := s.provCityService.GetDetailCity(user.IdKota)

	return &dto.UserProfileResponse{
		Nama:         user.Nama,
		NoTelp:       user.NoTelp,
		TanggalLahir: user.TanggalLahir,
		Tentang:      user.Tentang,
		Pekerjaan:    user.Pekerjaan,
		Email:        user.Email,
		IdProvinsi:   prov,
		IdKota:       city,
	}, nil
}

func (s *userService) UpdateProfile(userID uint, req *dto.UpdateProfileRequest) error {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return err
	}

	// Check duplicates if email or notelp is changed
	emailToCheck := user.Email
	if req.Email != "" {
		emailToCheck = req.Email
	}
	noTelpToCheck := user.NoTelp
	if req.NoTelp != "" {
		noTelpToCheck = req.NoTelp
	}

	if err := s.userRepo.CheckEmailOrPhoneExists(emailToCheck, noTelpToCheck, userID); err != nil {
		return err
	}

	if req.Nama != "" {
		user.Nama = req.Nama
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.NoTelp != "" {
		user.NoTelp = req.NoTelp
	}
	if req.TanggalLahir != "" {
		user.TanggalLahir = req.TanggalLahir
	}
	if req.Pekerjaan != "" {
		user.Pekerjaan = req.Pekerjaan
	}
	if req.IdProvinsi != "" {
		user.IdProvinsi = req.IdProvinsi
	}
	if req.IdKota != "" {
		user.IdKota = req.IdKota
	}
	if req.Tentang != "" {
		user.Tentang = req.Tentang
	}
	if req.KataSandi != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(req.KataSandi), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.KataSandi = string(hashed)
	}

	return s.userRepo.Update(user)
}
