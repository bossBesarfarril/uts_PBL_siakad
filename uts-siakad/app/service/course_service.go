package service

import (
	"uts-siakad/app/model"
	"uts-siakad/app/repository"
)

// Endpoint 8: GET /courses
func ListCourses(semester string, search string, available bool) ([]model.CourseResponse, error) {
	return repository.FindAllCourses(semester, search, available)
}
