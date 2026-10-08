package main

import (
	"log"
	"os"
	"uts-siakad/config"
	"uts-siakad/database"
	"uts-siakad/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// 1. Load config & database
	config.LoadEnv()
	database.ConnectDB()

	// 2. Inisialisasi Fiber
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("SIAKAD Mini API is Running! 🚀")
	})

	// 3. Pasang semua rute (10 Endpoint)
	route.SetupRoutes(app)

	// 4. Jalankan Server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}
	
	log.Printf("Menyalakan server di port %s...", port)
	log.Fatal(app.Listen(":" + port))
}