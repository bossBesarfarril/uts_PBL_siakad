package repository

import (
	"context"
	"uts-siakad/database"
)

// Cek apakah mahasiswa ini sudah ambil matkul ini di tahun yang sama
func CheckDuplicateEnrollment(studentID int, courseID int, tahun string) bool {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM enrollments WHERE student_id = $1 AND course_id = $2 AND tahun_akademik = $3)"
	database.DB.QueryRow(context.Background(), query, studentID, courseID, tahun).Scan(&exists)
	return exists
}

// Menghitung SKS yang sudah diambil di suatu tahun akademik
func CountTotalSKS(studentID int, tahun string) int {
	var total int
	query := `
		SELECT COALESCE(SUM(c.sks), 0) FROM enrollments e
		JOIN courses c ON e.course_id = c.id
		WHERE e.student_id = $1 AND e.tahun_akademik = $2
	`
	database.DB.QueryRow(context.Background(), query, studentID, tahun).Scan(&total)
	return total
}

func DeleteEnrollment(id int, studentID int) (bool, error) {
	tag, err := database.DB.Exec(context.Background(), "DELETE FROM enrollments WHERE id = $1 AND student_id = $2", id, studentID)
	return tag.RowsAffected() > 0, err // Mengembalikan true jika beneran ada yang terhapus
}