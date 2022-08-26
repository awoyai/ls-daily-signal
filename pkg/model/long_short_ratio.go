package model

import (
	"time"

	"gorm.io/gorm"
)

type LongShortFilter struct {
	TradingAt int64
	Varieties []string
}

type LongShortRatios []*LongShortRatio

type LongShortRatio struct {
	gorm.Model

	TradingDate time.Time
	TradingAt   int64 `gorm:"idx:idx_trading_at"`

	Variety string `gorm:"idx:idx_variety;type:varchar(10)"`
	Ratio   float64
}

func (s LongShortRatios) GetMap() map[string]*LongShortRatio {
	ratioMap := make(map[string]*LongShortRatio)
	for _, v := range s {
		ratioMap[v.Variety] = v
	}
	return ratioMap
}