package storage

import (
	"context"
	"fmt"
	"kanlam/lib/config"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectMainDb() {
	db := config.GetDbConfig()
	pool, err := pgxpool.New(context.Background(), db.Url)
	MainPool = pool

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(db)
	fmt.Println("Connected to main DB")
}
