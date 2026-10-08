package service

import (
	"context"
	"fmt"
	"math"
	"strings"
	"uts-siakad/app/model"
	"uts-siakad/app/repository"
	"uts-siakad/database"
	"uts-siakad/helper"

	"golang.org/x/crypto/bcrypt"
)

// Endpoint 3: GET /students
func ListStudents(page int, perPage int, prodi string, angkatan string, search string, sort string) ([]model.Student, *model.Meta, error) {
	// Menghindari paginasi ngawur
	if page < 1 {
		page = 1
	}
	if perPage < 1 {
		perPage = 10
	}
	if perPage > 50 {
		perPage = 50
	}

	// Dasar Query
	query := "SELECT id, user_id, nim, nama, prodi, angkatan, ipk_terakhir FROM students WHERE deleted_at IS NULL"
	countQuery := "SELECT COUNT(*) FROM students WHERE deleted_at IS NULL"
	
	args := []interface{}{}
	argId := 1

	// Filter Prodi
	if prodi != "" {
		query += fmt.Sprintf(" AND prodi = $%d", argId)
		countQuery += fmt.Sprintf(" AND prodi = $%d", argId)
		args = append(args, prodi)
		argId++
	}

	// Filter Angkatan
	if angkatan != "" {
		query += fmt.Sprintf(" AND angkatan = $%d", argId)
		countQuery += fmt.Sprintf(" AND angkatan = $%d", argId)
		args = append(args, angkatan)
		argId++
	}

	// Filter Search (NIM / Nama)
	if search != "" {
		query += fmt.Sprintf(" AND (nim ILIKE '%%' || $%d || '%%' OR nama ILIKE '%%' || $%d || '%%')", argId, argId)
		countQuery += fmt.Sprintf(" AND (nim ILIKE '%%' || $%d || '%%' OR nama ILIKE '%%' || $%d || '%%')", argId, argId)
		args = append(args, search)
		argId++
	}

	// Sorting
	if sort != "" {
		if strings.HasPrefix(sort, "-") {
			column := strings.TrimPrefix(sort, "-")
			query += " ORDER BY " + column + " DESC"
		} else {
			query += " ORDER BY " + sort + " ASC"
		}
	} else {
		query += " ORDER BY id DESC" // Default sorting
	}

	// Hitung Total Data untuk Meta
	var totalData int
	database.DB.QueryRow(context.Background(), countQuery, args...).Scan(&totalData)

	// Pagination OFFSET
	offset := (page - 1) * perPage
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argId, argId+1)
	args = append(args, perPage, offset)

	// Eksekusi Ambil Data
	rows, err := database.DB.Query(context.Background(), query, args...)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var students []model.Student
	for rows.Next() {
		var s model.Student
		if err := rows.Scan(&s.ID, &s.UserID, &s.NIM, &s.Nama, &s.Prodi, &s.Angkatan, &s.IPKTerakhir); err == nil {
			students = append(students, s)
		}
	}

	// Susun Meta
	meta := &model.Meta{
		CurrentPage: page,
		PerPage:     perPage,
		Total:       totalData,
		LastPage:    int(math.Ceil(float64(totalData) / float64(perPage))),
	}

	return students, meta, nil
}

// Endpoint 4: POST /students
func CreateStudent(req model.CreateStudentRequest) error {
	// Cek Validasi Struct
	err := helper.Validate.Struct(req)
	if err != nil {
		return err // Biar ditangani ErrorHandler
	}

	// Cek Duplikat NIM / Email
	exists, err := repository.CheckNimOrEmailExists(req.NIM, req.Email)
	if err != nil {
		return helper.BadRequest("Gagal mengecek duplikasi data")
	}
	if exists {
		return helper.ValidationError("NIM atau Email sudah terdaftar")
	}

	// --- TRANSACTION DIMULAI ---
	ctx := context.Background()
	tx, err := database.DB.Begin(ctx)
	if err != nil {
		return helper.BadRequest("Gagal memulai transaksi database")
	}
	defer tx.Rollback(ctx) // Pastikan rollback jika error di tengah jalan

	// 1. Buat User (Password Hash = NIM)
	hashedNIM, _ := bcrypt.GenerateFromPassword([]byte(req.NIM), bcrypt.DefaultCost)
	var userID int
	err = tx.QueryRow(ctx, "INSERT INTO users (email, password, role) VALUES ($1, $2, 'mahasiswa') RETURNING id", req.Email, string(hashedNIM)).Scan(&userID)
	if err != nil {
		return helper.BadRequest("Gagal membuat user mahasiswa")
	}

	// 2. Buat Student
	ipk := 0.00
	if req.IPKTerakhir != nil {
		ipk = *req.IPKTerakhir
	}

	_, err = tx.Exec(ctx, "INSERT INTO students (user_id, nim, nama, prodi, angkatan, ipk_terakhir) VALUES ($1, $2, $3, $4, $5, $6)",
		userID, req.NIM, req.Nama, req.Prodi, req.Angkatan, ipk)
	if err != nil {
		return helper.BadRequest("Gagal membuat profil mahasiswa")
	}

	// --- TRANSACTION SELESAI ---
	return tx.Commit(ctx)
}

// Endpoint 5: GET /students/{id}
func GetStudentDetail(id int) (interface{}, error) {
	// Ambil profil dasar
	student, err := repository.FindStudentByID(id)
	if err != nil {
		return nil, helper.NotFound("Mahasiswa tidak ditemukan atau sudah dihapus")
	}

	// Tentukan Batas SKS berdasarkan Business Rule UTS
	batasSKS := 18
	if student.IPKTerakhir >= 3.00 {
		batasSKS = 24
	} else if student.IPKTerakhir >= 2.50 {
		batasSKS = 21
	}

	// Ambil total SKS yang sudah diambil di semua tahun akademik
	var totalSKS int
	database.DB.QueryRow(context.Background(), "SELECT COALESCE(SUM(c.sks), 0) FROM enrollments e JOIN courses c ON e.course_id = c.id WHERE e.student_id = $1", student.ID).Scan(&totalSKS)

	// Ambil daftar matkul (Enrollments)
	type MatkulKRS struct {
		IDEnrollment  int    `json:"id_enrollment"`
		KodeMK        string `json:"kode_mk"`
		NamaMK        string `json:"nama_mk"`
		SKS           int    `json:"sks"`
		TahunAkademik string `json:"tahun_akademik"`
	}
	var krsList []MatkulKRS

	rows, _ := database.DB.Query(context.Background(), "SELECT e.id, c.kode_mk, c.nama_mk, c.sks, e.tahun_akademik FROM enrollments e JOIN courses c ON e.course_id = c.id WHERE e.student_id = $1 ORDER BY e.tahun_akademik DESC", student.ID)
	defer rows.Close()
	for rows.Next() {
		var mk MatkulKRS
		rows.Scan(&mk.IDEnrollment, &mk.KodeMK, &mk.NamaMK, &mk.SKS, &mk.TahunAkademik)
		krsList = append(krsList, mk)
	}

	// Gabungkan Response
	detail := struct {
		model.Student
		TotalSKS  int         `json:"total_sks"`
		BatasSKS  int         `json:"batas_sks"`
		DaftarKRS []MatkulKRS `json:"daftar_krs"`
	}{
		Student:   student,
		TotalSKS:  totalSKS,
		BatasSKS:  batasSKS,
		DaftarKRS: krsList,
	}

	return detail, nil
}

// Endpoint 6: PUT /students/{id}
func UpdateStudent(id int, req model.UpdateStudentRequest) error {
	err := helper.Validate.Struct(req)
	if err != nil {
		return err
	}

	_, err = repository.FindStudentByID(id)
	if err != nil {
		return helper.NotFound("Mahasiswa tidak ditemukan")
	}

	return repository.UpdateStudent(id, req)
}

// Endpoint 7: DELETE /students/{id}
func DeleteStudent(id int) error {
	_, err := repository.FindStudentByID(id)
	if err != nil {
		return helper.NotFound("Mahasiswa tidak ditemukan")
	}
	return repository.SoftDeleteStudent(id)
}
