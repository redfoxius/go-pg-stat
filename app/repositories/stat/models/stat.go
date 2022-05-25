package models

type StatFilter struct {
	Type  string
	Sort  string
	Page  int
	Limit int
}

type StatResult struct {
	Items []StatRow `json:"items"`
}

type StatRow struct {
	Query    string  `gorm:"column:query" json:"query"`
	MaxTime  float64 `gorm:"column:max_exec_time" json:"max_time"`
	MeanTime float64 `gorm:"column:mean_exec_time" json:"mean_time"`
}

func (g StatRow) TableName() string {
	return "pg_stat_statements"
}
