// database access
package repo

import (
	"context"
	"fmt"
	"os"
)

var _ repository = (*Repository)(nil)

type Repository struct {
}

func (r *Repository) ExecQuery(ctx context.Context, sqlQuery string) (map[string]any, error) {
	// Queryの実行
	dbResp, err := connPool.Query(context.Background(), "select 'Hello, world!'")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return nil, fmt.Errorf("Failed to execute query. error is %s.", err)
	}

	// カラムの取得
	columns := dbResp.FieldDescriptions()
	// クエリから返される各行のデータを一時的に格納する
	values := make([]interface{}, len(columns))
	// values スライスの各要素へのポインタを格納する
	valuePtrs := make([]interface{}, len(columns))
	// カラム名をキーにして値を保管する
	rowMap := make(map[string]interface{})
	for i := range values {
		valuePtrs[i] = &values[i]
	}
	for dbResp.Next() {
		if err := dbResp.Scan(valuePtrs...); err != nil {
			break
		}

	}
	return dbResp, nil
}

func (r *Repository) Get(ctx context.Context, sqlQuery string) any {
	return nil
}

func (r *Repository) Create(ctx context.Context, sqlQuery string) any {
	return nil
}

func (r *Repository) Update(ctx context.Context, sqlQuery string) any {
	return nil
}

func (r *Repository) Delete(ctx context.Context, sqlQuery string) any {
	return nil
}
