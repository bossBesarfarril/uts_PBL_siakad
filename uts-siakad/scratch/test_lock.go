package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}

	rows, err := pool.Query(context.Background(), "SELECT pid, wait_event_type, wait_event, query FROM pg_stat_activity WHERE state = 'active'")
	if err != nil {
		fmt.Println("Error querying:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var pid int
		var wet, we, q *string
		rows.Scan(&pid, &wet, &we, &q)
		if we != nil && q != nil {
			fmt.Printf("PID: %d, Wait: %v, Query: %s\n", pid, *we, *q)
		} else if q != nil {
			fmt.Printf("PID: %d, Wait: NONE, Query: %s\n", pid, *q)
		}
	}
}
