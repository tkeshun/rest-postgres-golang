// database connection
package repo

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func setEnv(key, value string) func() {
// 	originalValue, originalValuePresent := os.LookupEnv(key)
// 	os.Setenv(key, value)

// 	return func() {
// 		if originalValuePresent {
// 			os.Setenv(key, originalValue) // Restore the original value
// 		} else {
// 			os.Unsetenv(key) // Remove the variable if it wasn't set originally
// 		}
// 	}
// }

func TestConnectDataBase(t *testing.T) {
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
