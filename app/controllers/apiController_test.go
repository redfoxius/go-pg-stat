package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/redfoxius/go-pg-stat/app/dto"
	"github.com/redfoxius/go-pg-stat/app/repositories/stat"
	"github.com/redfoxius/go-pg-stat/app/repositories/stat/models"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"gorm.io/gorm"
	"net/url"
	"testing"
)

type StatRepoMock struct {
	GetStatMock func(ctx context.Context, filter models.StatFilter) (*models.StatResult, error)
}

func (r StatRepoMock) GetStat(ctx context.Context, filter models.StatFilter) (*models.StatResult, error) {
	return r.GetStatMock(ctx, filter)
}

func TestGetRecords(t *testing.T) {
	tests := map[string]struct {
		repo  stat.StatRepository
		query map[string]string
		exp   dto.Response
		err   error
	}{
		"success": {
			repo: StatRepoMock{
				GetStatMock: func(ctx context.Context, filter models.StatFilter) (*models.StatResult, error) {
					return &models.StatResult{
						Items: []models.StatRow{
							{
								Query:    "SELECT 1",
								MaxTime:  20,
								MeanTime: 15,
							},
						},
					}, nil
				},
			},
			query: map[string]string{
				"type":  "",
				"sort":  "slow",
				"page":  "2",
				"limit": "5",
			},
			exp: dto.Response{
				Result: &models.StatResult{
					Items: []models.StatRow{
						{
							Query:    "SELECT 1",
							MaxTime:  20,
							MeanTime: 15,
						},
					},
				},
			},
		},
		"DB_error": {
			repo: StatRepoMock{
				GetStatMock: func(ctx context.Context, filter models.StatFilter) (*models.StatResult, error) {
					return nil, gorm.ErrRecordNotFound
				},
			},
			err: errors.New("record not found"),
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			c := InitController(test.repo)

			q := make(url.Values)
			for key, value := range test.query {
				q.Add(key, value)
			}

			var ctx fasthttp.RequestCtx
			var req fasthttp.Request
			req.SetRequestURI("http://127.0.0.1:13013/api/stat/get?" + q.Encode())
			ctx.Init(&req, nil, nil)
			err := c.GetRecords(fiber.New().AcquireCtx(&ctx))
			if err != nil {
				assert.IsType(t, test.err, err)
				assert.Equal(t, test.err.Error(), err.Error())

				return
			}

			var result dto.Response
			err = json.Unmarshal(ctx.Response.Body(), &result)
			if err != nil {
				t.Fatal(err)
			}

			assert.EqualValues(t, test.exp, result)
		})
	}
}
