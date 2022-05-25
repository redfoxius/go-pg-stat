package database

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetConnectionString(t *testing.T) {
	tests := map[string]struct {
		setEnv  func()
		dropEnv func()
		result  string
	}{
		"connect": {
			setEnv: func() {
				os.Setenv("PG_HOST", "127.0.0.1")
				os.Setenv("PG_PORT", "5432")
				os.Setenv("PG_USER", "postgres")
				os.Setenv("PG_PASSW", "postgres")
				os.Setenv("PG_DBNAME", "test")
			},
			dropEnv: func() {
				os.Unsetenv("PG_HOST")
				os.Unsetenv("PG_PORT")
				os.Unsetenv("PG_USER")
				os.Unsetenv("PG_PASSW")
				os.Unsetenv("PG_DBNAME")
			},
			result: "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=test sslmode=disable",
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.setEnv()

			res := GetConnectionString()
			assert.Equal(t, test.result, res)

			test.dropEnv()
		})
	}
}

func TestGetDB(t *testing.T) {
	tests := map[string]struct {
		setEnv   func()
		dropEnv  func()
		resEmpty bool
		errEmpty bool
	}{
		"fail": {
			setEnv: func() {
				os.Setenv("PG_HOST", "")
				os.Setenv("PG_PORT", "")
				os.Setenv("PG_USER", "")
				os.Setenv("PG_PASSW", "")
				os.Setenv("PG_DBNAME", "")
			},
			dropEnv: func() {
				os.Unsetenv("PG_HOST")
				os.Unsetenv("PG_PORT")
				os.Unsetenv("PG_USER")
				os.Unsetenv("PG_PASSW")
				os.Unsetenv("PG_DBNAME")
			},
			resEmpty: true,
			errEmpty: false,
		},
		"connect": {
			setEnv: func() {
				os.Setenv("PG_HOST", "127.0.0.1")
				os.Setenv("PG_PORT", "5432")
				os.Setenv("PG_USER", "postgres")
				os.Setenv("PG_PASSW", "postgres")
				os.Setenv("PG_DBNAME", "dvdrental")
			},
			dropEnv: func() {
				os.Unsetenv("PG_HOST")
				os.Unsetenv("PG_PORT")
				os.Unsetenv("PG_USER")
				os.Unsetenv("PG_PASSW")
				os.Unsetenv("PG_DBNAME")
			},
			resEmpty: false,
			errEmpty: true,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.setEnv()

			res, err := GetDB()

			assert.Equalf(t, res == nil, test.resEmpty, `check empty result`)
			assert.Equalf(t, err == nil, test.errEmpty, `check empty error`)

			test.dropEnv()
		})
	}
}
