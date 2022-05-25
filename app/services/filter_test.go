package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/redfoxius/go-pg-stat/app/repositories/stat/models"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"net/url"
	"testing"
)

func TestBuildFilter(t *testing.T) {
	types := [5]string{
		"", "select", "insert", "update", "delete",
	}
	sorts := [2]string{
		"fast", "slow",
	}

	for _, r := range types {
		for _, s := range sorts {
			queryParams := map[string]string{
				"type":  r,
				"sort":  s,
				"page":  "2",
				"limit": "5",
			}

			q := make(url.Values)
			for key, value := range queryParams {
				q.Add(key, value)
			}

			var ctx fasthttp.RequestCtx
			var req fasthttp.Request
			req.SetRequestURI("http://unit.test/queries?" + q.Encode())
			ctx.Init(&req, nil, nil)
			testFilter := BuildFilter(fiber.New().AcquireCtx(&ctx))

			sampleFilter := models.StatFilter{
				Type:  r,
				Sort:  s,
				Page:  2,
				Limit: 5,
			}

			assert.Equal(t, testFilter, sampleFilter)
		}
	}
}

func TestBuildFilterFixed(t *testing.T) {
	sampleFilter := models.StatFilter{
		Type:  "",
		Sort:  "slow",
		Page:  1,
		Limit: 10,
	}

	queryParams := map[string]string{
		"type": "qqq",
		"sort": "www",
	}

	q := make(url.Values)
	for key, value := range queryParams {
		q.Add(key, value)
	}

	var ctx fasthttp.RequestCtx
	var req fasthttp.Request
	req.SetRequestURI("http://unit.test/queries?" + q.Encode())
	ctx.Init(&req, nil, nil)
	testFilter := BuildFilter(fiber.New().AcquireCtx(&ctx))

	assert.Equal(t, testFilter, sampleFilter)
}
