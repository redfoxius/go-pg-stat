package stat

import (
	"context"
	"github.com/redfoxius/go-pg-stat/app/repositories/stat/models"
	"gorm.io/gorm"
)

type StatRepository interface {
	GetStat(ctx context.Context, filter models.StatFilter) (*models.StatResult, error)
}

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) StatRepository {
	return &repository{
		db,
	}
}

func (r *repository) GetStat(ctx context.Context, filter models.StatFilter) (*models.StatResult, error) {
	result := models.StatResult{}
	query := r.db.WithContext(ctx)

	if filter.Type != "" {
		query = query.Where("lower(query) like ?", filter.Type+`%`)
	}

	if filter.Sort == `fast` {
		query = query.Order("max_exec_time ASC")
	} else {
		query = query.Order("max_exec_time DESC")
	}

	if err := query.Limit(filter.Limit).Offset((filter.Page - 1) * filter.Limit).Find(&result.Items).Error; err != nil {
		return nil, err
	}

	return &result, nil
}
