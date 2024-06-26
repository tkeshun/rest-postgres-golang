package repo

import "context"

type repository interface {
	ExecQuery(ctx context.Context, sqlQuery string) (map[string]any, error)
	Get(ctx context.Context, sqlQuery string) any
	Create(ctx context.Context, sqlQuery string) any
	Update(ctx context.Context, sqlQuery string) any
	Delete(ctx context.Context, sqlQuery string) any
}
