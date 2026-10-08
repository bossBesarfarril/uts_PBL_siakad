package model

type Student struct {
	ID          int     `json:"id"`
	UserID      int     `json:"user_id"`
	NIM         string  `json:"nim"`
	Nama        string  `json:"nama"`
	Prodi       string  `json:"prodi"`
	Angkatan    int     `json:"angkatan"`
	IPKTerakhir float64 `json:"ipk_terakhir"`
}

// Untuk saat admin membuat mahasiswa baru (Endpoint 4)
type CreateStudentRequest struct {
	NIM         string   `json:"nim" validate:"required,len=12"`
	Nama        string   `json:"nama" validate:"required"`
	Email       string   `json:"email" validate:"required,email"`
	Prodi       string   `json:"prodi" validate:"required"`
	Angkatan    int      `json:"angkatan" validate:"required"`
	IPKTerakhir *float64 `json:"ipk_terakhir" validate:"omitempty,min=0,max=4"`
}

// Untuk saat admin update mahasiswa (Endpoint 6)
type UpdateStudentRequest struct {
	Nama        *string  `json:"nama" validate:"omitnil,min=3"`
	Prodi       *string  `json:"prodi" validate:"omitnil"`
	Angkatan    *int     `json:"angkatan" validate:"omitnil"`
	IPKTerakhir *float64 `json:"ipk_terakhir" validate:"omitnil,min=0,max=4"`
}
