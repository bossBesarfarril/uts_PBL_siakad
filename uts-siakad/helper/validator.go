package helper

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func InitValidator() {
	Validate = validator.New()
}

// Fungsi bantu untuk memformat pesan error validasi
func FormatValidationError(err error) map[string][]string {
	errors := make(map[string][]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			field := strings.ToLower(e.Field())
			switch e.Tag() {
			case "required":
				errors[field] = []string{e.Field() + " wajib diisi"}
			case "email":
				errors[field] = []string{"Format email tidak valid"}
			case "min":
				errors[field] = []string{"Minimal " + e.Param() + " karakter"}
			case "max":
				errors[field] = []string{"Maksimal " + e.Param() + " karakter"}
			case "len":
				errors[field] = []string{"Harus persis " + e.Param() + " karakter"}
			default:
				errors[field] = []string{"Validasi " + e.Tag() + " gagal"}
			}
		}
	}
	return errors
}
