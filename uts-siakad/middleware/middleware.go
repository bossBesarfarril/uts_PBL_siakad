package middleware

import (
	"os"
	"strings"
	"uts-siakad/app/model"
	"uts-siakad/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// 1. Error Handler Terpusat (Dari Modul 7)
func ErrorHandler(c *fiber.Ctx, err error) error {
	// Kalau errornya dari helper.AppError kita sendiri
	if e, ok := err.(*helper.AppError); ok {
		return c.Status(e.Code).JSON(model.WebResponse{
			Success: false,
			Message: e.Message,
		})
	}

	// Kalau errornya karena validasi gagal (Status 422)
	errorsMap := helper.FormatValidationError(err)
	if len(errorsMap) > 0 {
		return c.Status(422).JSON(model.WebResponse{
			Success: false,
			Message: "Validasi gagal",
			Errors:  errorsMap,
		})
	}

	// Kalau error yang nggak terduga (500)
	return c.Status(500).JSON(model.WebResponse{
		Success: false,
		Message: "Terjadi kesalahan pada server",
	})
}

// 2. Middleware Cek Token (Autentikasi)
func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return helper.Unauthorized("Token tidak ada atau salah format")
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return helper.Unauthorized("Token kedaluwarsa atau tidak valid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return helper.Unauthorized("Gagal membaca token")
	}

	// Simpan data user_id dan role ke dalam loker Fiber biar bisa dibaca di Endpoint
	c.Locals("user_id", int(claims["user_id"].(float64)))
	c.Locals("role", claims["role"].(string))

	return c.Next()
}

// 3. Middleware Cek Role (Admin / Mahasiswa)
func RoleGuard(allowedRole string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := c.Locals("role").(string)
		if role != allowedRole {
			return helper.Forbidden("Akses ditolak: khusus " + allowedRole)
		}
		return c.Next()
	}
}
