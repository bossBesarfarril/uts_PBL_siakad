package model

type Course struct {
	ID       int    `json:"id"`
	KodeMK   string `json:"kode_mk"`
	NamaMK   string `json:"nama_mk"`
	SKS      int    `json:"sks"`
	Semester int    `json:"semester"`
	Kuota    int    `json:"kuota"`
}

// Struct khusus untuk balasan Endpoint 8 (diminta menampilkan sisa kuota)
type CourseResponse struct {
	Course
	Terisi    int `json:"terisi"`
	SisaKuota int `json:"sisa_kuota"`
}
