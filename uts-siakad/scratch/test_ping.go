package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		fmt.Println("POOL ERR:", err)
		return
	}
	err = pool.Ping(context.Background())
	if err != nil {
		fmt.Println("PING ERR:", err)
	} else {
		fmt.Println("PING SUCCESS!")
	}
}
