package dto

type RegisterRequest struct {
	Nama         string `json:"nama" binding:"required"`
	KataSandi    string `json:"kata_sandi" binding:"required"`
	NoTelp       string `json:"no_telp" binding:"required"`
	TanggalLahir string `json:"tanggal_Lahir"`
	Pekerjaan    string `json:"pekerjaan"`
	Email        string `json:"email" binding:"required,email"`
	IdProvinsi   string `json:"id_provinsi" binding:"required"`
	IdKota       string `json:"id_kota" binding:"required"`
}

type LoginRequest struct {
	NoTelp    string `json:"no_telp" binding:"required"`
	KataSandi string `json:"kata_sandi" binding:"required"`
}

type UserProfileResponse struct {
	Nama         string      `json:"nama"`
	NoTelp       string      `json:"no_telp"`
	TanggalLahir string      `json:"tanggal_Lahir"`
	Tentang      string      `json:"tentang"`
	Pekerjaan    string      `json:"pekerjaan"`
	Email        string      `json:"email"`
	IdProvinsi   interface{} `json:"id_provinsi"`
	IdKota       interface{} `json:"id_kota"`
	Token        string      `json:"token,omitempty"`
}
