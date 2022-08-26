package model

import (
	"time"

	"gorm.io/gorm"
)

type SymbolPriceFilter struct {
	Symbols   []string
	TradingAt int64
}

type SymbolPrices []*SymbolPrice

type SymbolPrice struct {
	gorm.Model
	TradingDate time.Time
	TradingAt   int64  `gorm:"index:idx_trading_at"`
	Variety     string `gorm:"type:varchar(10);index:idx_variety;comment:期货品种代码"`
	Symbol      string `gorm:"type:varchar(10);comment:期货合约代码"`
	OpenPrice   float64
	ClosePrice  float64
	HighPrice   float64
	LowPrice    float64
}

func (s SymbolPrices) GetMap() map[string]*SymbolPrice {
	priceMap := make(map[string]*SymbolPrice)
	for _, v := range s {
		priceMap[v.Symbol] = v
	}
	return priceMap
}
