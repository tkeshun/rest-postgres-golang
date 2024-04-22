// database connection
package repo

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// TODO:環境変数化する
const (
	DATABASE_URL = "postgres://user:password@localhost:5432/rest-postgres"
	// DATABASE_URL = "host=localhost port=5432 user=user password=password dbname=rest-postgres sslmode=disable"
)

func ConnectDataBase(ctx context.Context) (*pgxpool.Pool, func(), error) {
	dbpool, err := pgxpool.New(ctx, DATABASE_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return dbpool, dbpool.Close, nil
}
