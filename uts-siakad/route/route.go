package route

import (
	"strconv"
	"time"
	"uts-siakad/app/model"
	"uts-siakad/app/repository"
	"uts-siakad/app/service"
	"uts-siakad/helper"
	"uts-siakad/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func SetupRoutes(app *fiber.App) {
	// Inisialisasi custom validator
	helper.InitValidator()

	api := app.Group("/api/v1")

	// --- ENDPOINT 1: LOGIN (Publik + Rate Limiter) ---
	// Rate Limiter: Maksimal 5x request per menit (Syarat UTS: 429 Too Many Requests)
	loginLimiter := limiter.New(limiter.Config{
		Max:        5,
		Expiration: 1 * time.Minute,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(429).JSON(model.WebResponse{
				Success: false,
				Message: "Terlalu banyak percobaan login. Coba lagi dalam 1 menit.",
			})
		},
	})

	api.Post("/auth/login", loginLimiter, func(c *fiber.Ctx) error {
		var req model.LoginRequest
		if err := c.BodyParser(&req); err != nil {
			return middleware.ErrorHandler(c, helper.BadRequest("Format JSON tidak valid"))
		}

		res, err := service.Login(req)
		if err != nil {
			return middleware.ErrorHandler(c, err)
		}
		return helper.Success(c, 200, "Login berhasil", res)
	})

	// ==========================================
	// GROUP AUTHENTICATION (Wajib Pakai Token)
	// ==========================================
	auth := api.Group("", middleware.AuthMiddleware)

	// --- ENDPOINT 2: PROFIL (Semua Role) ---
	auth.Get("/auth/me", func(c *fiber.Ctx) error {
		userID := c.Locals("user_id").(int)
		role := c.Locals("role").(string)

		profile, err := service.GetProfile(userID, role)
		if err != nil {
			return middleware.ErrorHandler(c, err)
		}
		return helper.Success(c, 200, "Profil berhasil diambil", profile)
	})

	// --- ENDPOINT 8: DAFTAR MATA KULIAH (Semua Role) ---
	auth.Get("/courses", func(c *fiber.Ctx) error {
		semester := c.Query("semester")
		search := c.Query("search")
		available := c.QueryBool("available", false)

		courses, err := service.ListCourses(semester, search, available)
		if err != nil {
			return middleware.ErrorHandler(c, err)
		}
		return helper.Success(c, 200, "Daftar mata kuliah berhasil diambil", courses)
	})

	// --- ENDPOINT 5: DETAIL MAHASISWA & SKS ---
	// Bisa diakses Admin, ATAU Mahasiswa (tapi hanya untuk datanya sendiri)
	auth.Get("/students/:id", func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))
		role := c.Locals("role").(string)
		userID := c.Locals("user_id").(int)

		// Syarat UTS: Mahasiswa tidak boleh melihat data mahasiswa lain (403 Forbidden)
		if role == "mahasiswa" {
			s, err := repository.FindStudentByUserID(userID)
			if err != nil || s.ID != id {
				return middleware.ErrorHandler(c, helper.Forbidden("Anda tidak berhak mengakses data KRS mahasiswa lain"))
			}
		}

		detail, err := service.GetStudentDetail(id)
		if err != nil {
			return middleware.ErrorHandler(c, err)
		}
		return helper.Success(c, 200, "Detail mahasiswa dan KRS berhasil diambil", detail)
	})

	// ==========================================
	// GROUP ADMIN (Khusus Role Admin)
	// ==========================================
	admin := auth.Group("/students", middleware.RoleGuard("admin"))

	// --- ENDPOINT 3: DAFTAR MAHASISWA ---
	admin.Get("/", func(c *fiber.Ctx) error {
		page, _ := strconv.Atoi(c.Query("page", "1"))
		perPage, _ := strconv.Atoi(c.Query("per_page", "10"))
		prodi := c.Query("prodi")
		angkatan := c.Query("angkatan")
		search := c.Query("search")
		sort := c.Query("sort")

		students, meta, err := service.ListStudents(page, perPage, prodi, angkatan, search, sort)
		if err != nil {
			return middleware.ErrorHandler(c, err)
		}
		return helper.SuccessList(c, "Daftar mahasiswa berhasil diambil", students, meta)
	})

	// --- ENDPOINT 4: TAMBAH MAHASISWA BARU ---
	admin.Post("/", func(c *fiber.Ctx) error {
		var req model.CreateStudentRequest
		if err := c.BodyParser(&req); err != nil {
			return middleware.ErrorHandler(c, helper.BadRequest("Format JSON tidak valid"))
		}

		err := service.CreateStudent(req)
		if err != nil {
			return middleware.ErrorHandler(c, err)
		}
		return helper.Success(c, 201, "Mahasiswa dan User berhasil ditambahkan", nil) // 201 Created
	})

	// --- ENDPOINT 6: UPDATE DATA MAHASISWA ---
	admin.Put("/:id", func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))
		var req model.UpdateStudentRequest
		if err := c.BodyParser(&req); err != nil {
			return middleware.ErrorHandler(c, helper.BadRequest("Format JSON tidak valid"))
		}

		err := service.UpdateStudent(id, req)
		if err != nil {
			return middleware.ErrorHandler(c, err)
		}
		return helper.Success(c, 200, "Data mahasiswa berhasil diperbarui", nil)
	})

	// --- ENDPOINT 7: SOFT DELETE MAHASISWA ---
	admin.Delete("/:id", func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))
		err := service.DeleteStudent(id)
		if err != nil {
			return middleware.ErrorHandler(c, err)
		}
		return helper.NoContent(c) // 204 No Content
	})

	// ==========================================
	// GROUP MAHASISWA (Khusus Role Mahasiswa)
	// ==========================================
	mhs := auth.Group("/enrollments", middleware.RoleGuard("mahasiswa"))

	// --- ENDPOINT 9: AMBIL MATA KULIAH (KRS) ---
	mhs.Post("/", func(c *fiber.Ctx) error {
		var req model.CreateEnrollmentRequest
		if err := c.BodyParser(&req); err != nil {
			return middleware.ErrorHandler(c, helper.BadRequest("Format JSON tidak valid"))
		}
		userID := c.Locals("user_id").(int)

		err := service.CreateEnrollment(req, userID)
		if err != nil {
			return middleware.ErrorHandler(c, err)
		}
		return helper.Success(c, 201, "Mata kuliah berhasil ditambahkan ke KRS", nil) // 201 Created
	})

	// --- ENDPOINT 10: BATALKAN MATA KULIAH DARI KRS ---
	mhs.Delete("/:id", func(c *fiber.Ctx) error {
		enrollmentID, _ := strconv.Atoi(c.Params("id"))
		userID := c.Locals("user_id").(int)

		err := service.DeleteEnrollment(enrollmentID, userID)
		if err != nil {
			return middleware.ErrorHandler(c, err)
		}
		return helper.NoContent(c) // 204 No Content
	})
}
