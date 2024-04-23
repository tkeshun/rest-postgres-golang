// database connection
package repo

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var connPool *pgxpool.Pool

func ConnectDataBase(ctx context.Context) (func(), error) {
	dbpool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	connPool = dbpool
	return dbpool.Close, nil
}
