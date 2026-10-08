package service

import (
	"context"
	"fmt"
	"uts-siakad/app/model"
	"uts-siakad/app/repository"
	"uts-siakad/database"
	"uts-siakad/helper"
)

// Endpoint 9: POST /enrollments
func CreateEnrollment(req model.CreateEnrollmentRequest, userID int) error {
	// 1. Validasi Input
	err := helper.Validate.Struct(req)
	if err != nil {
		return err
	}

	// 2. Ambil data mahasiswa pemilik token
	student, err := repository.FindStudentByUserID(userID)
	if err != nil {
		return helper.Forbidden("Hanya mahasiswa yang bisa mengambil KRS")
	}

	// 3. Cek Duplikasi (Business Rule)
	if repository.CheckDuplicateEnrollment(student.ID, req.CourseID, req.TahunAkademik) {
		return helper.Conflict("Mata kuliah ini sudah diambil pada tahun akademik yang sama")
	}

	// --- TRANSACTION + ROW LOCKING DIMULAI ---
	ctx := context.Background()
	tx, err := database.DB.Begin(ctx)
	if err != nil {
		return helper.BadRequest("Gagal memulai transaksi")
	}
	defer tx.Rollback(ctx)

	// 4. Cek Kuota dengan SELECT ... FOR UPDATE (Row Locking)
	var kuota, terisi, sksMatkul int
	queryLock := `
		SELECT c.kuota, c.sks, (SELECT COUNT(*) FROM enrollments e WHERE e.course_id = c.id) 
		FROM courses c WHERE c.id = $1 FOR UPDATE
	`
	err = tx.QueryRow(ctx, queryLock, req.CourseID).Scan(&kuota, &sksMatkul, &terisi)
	if err != nil {
		return helper.NotFound("Mata kuliah tidak ditemukan")
	}

	if terisi >= kuota {
		return helper.ValidationError("Kuota mata kuliah ini sudah penuh")
	}

	// 5. Cek Batas SKS (Business Rule)
	batasSKS := 18
	if student.IPKTerakhir >= 3.00 {
		batasSKS = 24
	} else if student.IPKTerakhir >= 2.50 {
		batasSKS = 21
	}

	var totalSksSaatIni int
	querySks := `SELECT COALESCE(SUM(c.sks), 0) FROM enrollments e JOIN courses c ON e.course_id = c.id WHERE e.student_id = $1 AND e.tahun_akademik = $2`
	tx.QueryRow(ctx, querySks, student.ID, req.TahunAkademik).Scan(&totalSksSaatIni)
	
	if totalSksSaatIni+sksMatkul > batasSKS {
		sisaSks := batasSKS - totalSksSaatIni
		msg := fmt.Sprintf("Gagal mengambil KRS. Sisa SKS anda hanya %d SKS, sedangkan matkul ini %d SKS.", sisaSks, sksMatkul)
		return helper.ValidationError(msg)
	}

	// 6. Insert KRS
	_, err = tx.Exec(ctx, "INSERT INTO enrollments (student_id, course_id, tahun_akademik) VALUES ($1, $2, $3)", student.ID, req.CourseID, req.TahunAkademik)
	if err != nil {
		return helper.BadRequest("Gagal menyimpan KRS")
	}

	return tx.Commit(ctx)
}

// Endpoint 10: DELETE /enrollments/{id}
func DeleteEnrollment(enrollmentID int, userID int) error {
	// Ambil data mahasiswa pemilik token (memastikan Kepemilikan)
	student, err := repository.FindStudentByUserID(userID)
	if err != nil {
		return helper.Forbidden("Hanya mahasiswa yang bisa membatalkan KRS")
	}

	sukses, err := repository.DeleteEnrollment(enrollmentID, student.ID)
	if err != nil {
		return helper.BadRequest("Gagal membatalkan KRS")
	}
	
	if !sukses {
		return helper.NotFound("KRS tidak ditemukan atau bukan milik Anda")
	}

	return nil
}
