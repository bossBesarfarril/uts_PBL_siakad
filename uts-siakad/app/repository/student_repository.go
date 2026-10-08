package repository

import (
	"context"
	"uts-siakad/app/model"
	"uts-siakad/database"
)

func FindStudentByID(id int) (model.Student, error) {
	var s model.Student
	query := "SELECT id, user_id, nim, nama, prodi, angkatan, ipk_terakhir FROM students WHERE id = $1 AND deleted_at IS NULL"
	err := database.DB.QueryRow(context.Background(), query, id).Scan(&s.ID, &s.UserID, &s.NIM, &s.Nama, &s.Prodi, &s.Angkatan, &s.IPKTerakhir)
	return s, err
}

func FindStudentByUserID(userID int) (model.Student, error) {
	var s model.Student
	query := "SELECT id, user_id, nim, nama, prodi, angkatan, ipk_terakhir FROM students WHERE user_id = $1 AND deleted_at IS NULL"
	err := database.DB.QueryRow(context.Background(), query, userID).Scan(&s.ID, &s.UserID, &s.NIM, &s.Nama, &s.Prodi, &s.Angkatan, &s.IPKTerakhir)
	return s, err
}

func CheckNimOrEmailExists(nim string, email string) (bool, error) {
	var exists bool
	query := `
		SELECT EXISTS (
			SELECT 1 FROM students s JOIN users u ON s.user_id = u.id 
			WHERE s.nim = $1 OR u.email = $2
		)
	`
	err := database.DB.QueryRow(context.Background(), query, nim, email).Scan(&exists)
	return exists, err
}

// Untuk update mahasiswa
func UpdateStudent(id int, req model.UpdateStudentRequest) error {
	query := "UPDATE students SET updated_at = NOW()"

	// Teknik modifikasi query hanya untuk field yang dikirim (omitnil)
	if req.Nama != nil {
		query += ", nama = '" + *req.Nama + "'"
	}
	if req.Prodi != nil {
		query += ", prodi = '" + *req.Prodi + "'"
	}
	if req.Angkatan != nil {
		// Konversi int ke query aman dilakukan manual atau lewat argumen (disingkat untuk modul)
		database.DB.Exec(context.Background(), "UPDATE students SET angkatan = $1 WHERE id = $2", *req.Angkatan, id)
	}
	if req.IPKTerakhir != nil {
		database.DB.Exec(context.Background(), "UPDATE students SET ipk_terakhir = $1 WHERE id = $2", *req.IPKTerakhir, id)
	}

	// Eksekusi final jika ada update nama/prodi string
	_, err := database.DB.Exec(context.Background(), query+" WHERE id = $1", id)
	return err
}

func SoftDeleteStudent(id int) error {
	_, err := database.DB.Exec(context.Background(), "UPDATE students SET deleted_at = NOW() WHERE id = $1", id)
	return err
}
