package stat

import (
	"context"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/redfoxius/go-pg-stat/app/repositories/stat/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"regexp"
	"testing"
)

type params struct {
	ctx    context.Context
	filter models.StatFilter
}

func NewMock(t *testing.T) (*sql.DB, *gorm.DB, sqlmock.Sqlmock) {
	conn, mock, err := sqlmock.New()

	if err != nil {
		log.Fatal("sqlmock.New error: ", err)
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "sqlmock",
		DriverName:           "postgres",
		Conn:                 conn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("gorm.Open error: ", err)
	}

	return conn, db, mock
}

func TestRepository_GetStat(t *testing.T) {
	ctx := context.Background()

	rows := []string{"query", "max_exec_time", "mean_exec_time"}

	tests := map[string]struct {
		callParams params
		mock       func(mock sqlmock.Sqlmock)
		result     []models.StatRow
		err        error
	}{
		"default_params": {
			callParams: params{
				ctx: ctx,
				filter: models.StatFilter{
					Type:  "",
					Sort:  "slow",
					Page:  1,
					Limit: 10,
				},
			},
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "pg_stat_statements" ORDER BY max_exec_time DESC LIMIT 10`)).
					WillReturnRows(sqlmock.NewRows(rows).AddRow("SELECT 2", 1.5, 0.01).AddRow("SELECT 1", 0.5, 0.05))
			},
			result: []models.StatRow{
				{
					Query:    "SELECT 2",
					MaxTime:  1.5,
					MeanTime: 0.01,
				},
				{
					Query:    "SELECT 1",
					MaxTime:  0.5,
					MeanTime: 0.05,
				},
			},
		},
		"with_type": {
			callParams: params{
				ctx: ctx,
				filter: models.StatFilter{
					Type:  "select",
					Sort:  "slow",
					Page:  1,
					Limit: 10,
				},
			},
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "pg_stat_statements" WHERE lower(query) like $1 ORDER BY max_exec_time DESC LIMIT 10`)).
					WithArgs("select%").
					WillReturnRows(sqlmock.NewRows(rows).AddRow("SELECT 2", 1.5, 0.01).AddRow("SELECT 1", 0.5, 0.05))
			},
			result: []models.StatRow{
				{
					Query:    "SELECT 2",
					MaxTime:  1.5,
					MeanTime: 0.01,
				},
				{
					Query:    "SELECT 1",
					MaxTime:  0.5,
					MeanTime: 0.05,
				},
			},
		},
		"db_error": {
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "pg_stat_statements" ORDER BY max_exec_time DESC`)).
					WillReturnError(errors.New("db_error"))
			},
			result: []models.StatRow{},
			err:    errors.New("db_error"),
		},
	}

	for test, item := range tests {
		t.Run(test, func(t *testing.T) {

			conn, db, mock := NewMock(t)
			defer conn.Close()

			item.mock(mock)

			result, err := New(db).GetStat(
				item.callParams.ctx,
				item.callParams.filter,
			)

			if err != nil {
				assert.Equal(t, item.err, err)
			} else {
				res := *result

				assert.Equal(t, item.result, res.Items)
			}
		})
	}
}
