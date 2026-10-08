package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	godotenv.Load("../.env") // Load env from parent directory

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		fmt.Println("POOL ERROR:", err)
		return
	}

	email := "admin@siakad.ac.id"
	password := "admin123"

	var hash string
	err = pool.QueryRow(context.Background(), "SELECT password FROM users WHERE email = $1", email).Scan(&hash)
	if err != nil {
		fmt.Println("DB QUERY ERROR:", err)
		return
	}

	fmt.Println("HASH KETEMU:", hash)

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		fmt.Println("BCRYPT ERROR:", err)
	} else {
		fmt.Println("LOGIN SUKSES!")
	}
}
