package repository

import (
	"context"
	"fmt"
	"uts-siakad/app/model"
	"uts-siakad/database"
)

func FindAllCourses(semester string, search string, available bool) ([]model.CourseResponse, error) {
	courses := []model.CourseResponse{}

	// Menghitung jumlah yang terisi secara langsung dari tabel enrollments
	query := `
		SELECT c.id, c.kode_mk, c.nama_mk, c.sks, c.semester, c.kuota, 
		       (SELECT COUNT(*) FROM enrollments e WHERE e.course_id = c.id) as terisi
		FROM courses c
		WHERE 1=1
	`

	if semester != "" {
		query += fmt.Sprintf(" AND c.semester = %s", semester)
	}
	if search != "" {
		query += fmt.Sprintf(" AND (c.kode_mk ILIKE '%%%s%%' OR c.nama_mk ILIKE '%%%s%%')", search, search)
	}

	rows, err := database.DB.Query(context.Background(), query)
	if err != nil {
		return courses, err
	}
	defer rows.Close()

	for rows.Next() {
		var c model.CourseResponse
		err := rows.Scan(&c.ID, &c.KodeMK, &c.NamaMK, &c.SKS, &c.Semester, &c.Kuota, &c.Terisi)
		if err != nil {
			return courses, err
		}

		c.SisaKuota = c.Kuota - c.Terisi

		// Filter available (hanya tampilkan jika sisa kuota > 0)
		if available && c.SisaKuota <= 0 {
			continue
		}
		courses = append(courses, c)
	}
	return courses, nil
}
