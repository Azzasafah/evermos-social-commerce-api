package service

import (
	"errors"

	"evermos-backend/internal/dto"
	"evermos-backend/internal/helper"
	"evermos-backend/internal/model"
	"evermos-backend/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(req *dto.RegisterRequest) error
	Login(req *dto.LoginRequest) (*dto.UserProfileResponse, error)
}

type authService struct {
	userRepo        repository.UserRepository
	tokoRepo        repository.TokoRepository
	provCityService ProvCityService
}

func NewAuthService(
	userRepo repository.UserRepository,
	tokoRepo repository.TokoRepository,
	provCityService ProvCityService,
) AuthService {
	return &authService{
		userRepo:        userRepo,
		tokoRepo:        tokoRepo,
		provCityService: provCityService,
	}
}

func (s *authService) Register(req *dto.RegisterRequest) error {
	// Check duplicate email or phone
	if err := s.userRepo.CheckEmailOrPhoneExists(req.Email, req.NoTelp, 0); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.KataSandi), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := model.User{
		Nama:         req.Nama,
		KataSandi:    string(hashedPassword),
		NoTelp:       req.NoTelp,
		TanggalLahir: req.TanggalLahir,
		Pekerjaan:    req.Pekerjaan,
		Email:        req.Email,
		IdProvinsi:   req.IdProvinsi,
		IdKota:       req.IdKota,
		IsAdmin:      false,
	}

	if err := s.userRepo.Create(&user); err != nil {
		return err
	}

	// Requirement 2 & 6: Toko otomatis terbuat ketika user mendaftar
	toko := model.Toko{
		IdUser:   user.ID,
		NamaToko: user.Nama,
		UrlFoto:  "",
	}
	if err := s.tokoRepo.Create(&toko); err != nil {
		return err
	}

	return nil
}

func (s *authService) Login(req *dto.LoginRequest) (*dto.UserProfileResponse, error) {
	user, err := s.userRepo.FindByNoTelp(req.NoTelp)
	if err != nil {
		return nil, errors.New("No Telp atau kata sandi salah")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.KataSandi), []byte(req.KataSandi)); err != nil {
		return nil, errors.New("No Telp atau kata sandi salah")
	}

	token, err := helper.GenerateToken(user.ID, user.Email, user.IsAdmin)
	if err != nil {
		return nil, err
	}

	prov, _ := s.provCityService.GetDetailProvince(user.IdProvinsi)
	city, _ := s.provCityService.GetDetailCity(user.IdKota)

	res := &dto.UserProfileResponse{
		Nama:         user.Nama,
		NoTelp:       user.NoTelp,
		TanggalLahir: user.TanggalLahir,
		Tentang:      user.Tentang,
		Pekerjaan:    user.Pekerjaan,
		Email:        user.Email,
		IdProvinsi:   prov,
		IdKota:       city,
		Token:        token,
	}

	return res, nil
}
