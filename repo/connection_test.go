// database connection
package repo

import (
	"context"
	"os"
	"testing"

	"dagger.io/dagger"
	"github.com/stretchr/testify/assert"
)

func TestConnectDataBase(t *testing.T) {

	ctx := context.Background()

	// create a Dagger client
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// Database service used for application tests
	database, err := client.Container().From("postgres:16").
		WithEnvVariable("POSTGRES_PASSWORD", "password").
		WithEnvVariable("POSTGRES_DB", "rest-postgres").
		WithEnvVariable("POSTGRES_USER", "user").
		WithExec([]string{"postgres"}).
		WithExposedPort(5432).
		AsService().Start(ctx)
	if err != nil {
		panic(err)
	}

	defer database.Stop(ctx)
	os.Setenv("DATABASE_URL", "postgres://user:password@localhost:5432/rest-postgres")
	tests := []struct {
		name string
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			closer, err := ConnectDataBase(ctx)
			defer closer()
			assert.NoError(t, err)
		})
	}
}
