package service

import (
	"uts-siakad/app/model"
	"uts-siakad/app/repository"
	"uts-siakad/helper"

	"golang.org/x/crypto/bcrypt"
)

func Login(req model.LoginRequest) (model.LoginResponse, error) {
	// 1. Cari user berdasarkan email
	user, err := repository.FindUserByEmail(req.Email)
	if err != nil {
		return model.LoginResponse{}, helper.Unauthorized("Email atau password salah")
	}

	// 2. Cek apakah password cocok
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return model.LoginResponse{}, helper.Unauthorized("Email atau password salah")
	}

	// 3. Khusus mahasiswa: cek apakah akun sudah di-soft delete
	if user.Role == "mahasiswa" {
		_, err := repository.FindStudentByUserID(user.ID)
		if err != nil {
			return model.LoginResponse{}, helper.Unauthorized("Akun ini sudah dinonaktifkan")
		}
	}

	// 4. Generate JWT Token
	token, expiresIn, err := helper.GenerateToken(user.ID, user.Role)
	if err != nil {
		return model.LoginResponse{}, helper.BadRequest("Gagal membuat token")
	}

	// 5. Susun Response
	res := model.LoginResponse{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   expiresIn,
		User:        user,
	}

	return res, nil
}

func GetProfile(userID int, role string) (interface{}, error) {
	// Jika admin, kembalikan data user saja
	if role == "admin" {
		user, err := repository.FindUserByID(userID)
		if err != nil {
			return nil, helper.NotFound("User tidak ditemukan")
		}
		return user, nil
	}

	// Jika mahasiswa, kembalikan data user + data profil student
	student, err := repository.FindStudentByUserID(userID)
	if err != nil {
		return nil, helper.NotFound("Profil mahasiswa tidak ditemukan")
	}

	// Gabungkan data user dan student menggunakan struct sementara (anonymous struct)
	profile := struct {
		model.Student
		Email string `json:"email"`
		Role  string `json:"role"`
	}{
		Student: student,
		Email:   "",
		Role:    "mahasiswa",
	}

	// Cari email aslinya
	user, _ := repository.FindUserByID(userID)
	profile.Email = user.Email

	return profile, nil
}
