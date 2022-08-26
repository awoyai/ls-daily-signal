package model

import (
	"time"
)

// offset
const (
	Open  = "open"
	Close = "close"
)

type PreOrderFilter struct {
	CreateDate *time.Time
	PlateName  string
}

// PreOrder 预订单
type PreOrder struct {
	ID         uint       `gorm:"primaryKey"`
	PlateName  string     `gorm:"type:varchar(50);index:idx_long_short_date,unique;comment:板块名"`
	CreateDate *time.Time `gorm:"type:date;index:idx_long_short_date,unique;comment:产生日期"`
	Long       string     `gorm:"type:varchar(25);index:idx_long_short_date,unique;comment:做多品种"`
	LongSize   float64    `gorm:"type:float(25);comment:做多数量"`
	Short      string     `gorm:"type:varchar(25);index:idx_long_short_date,unique;comment:做空品种"`
	ShortSize  float64    `gorm:"type:float(25);comment:做空数量"`
	Ratio      float64    `gorm:"type:float(25);comment:做多品种合约价格/做空品种合约价格"`
	State      string     `gorm:"-:all"`
}
