// database connection
package repo

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

// TODO:環境変数化する
const (
	DATABASE_URL = "postgres://user:password@localhost:5432/rest-postgres"
	// DATABASE_URL = "host=localhost port=5432 user=user password=password dbname=rest-postgres sslmode=disable"
)

func ConnectDataBase(ctx context.Context) (*pgx.Conn, func(context.Context) error, error) {
	conn, err := pgx.Connect(ctx, DATABASE_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn, conn.Close, nil
}
