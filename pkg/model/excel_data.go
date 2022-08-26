package model

type ExcelDatas []*ExcelData

type ExcelData struct {
	Date           string
	Plate          string
	Direction      string
	Varieties      string
	Symbol         string
	YesterdayClose float64
	TodayClose     float64
	Size           int
	YesterdayRatio float64
	TodayRatio     float64
	OpenRatio      string
	IsBetraied     string
}

func (e ExcelDatas) MergeRepeatSymbol() ExcelDatas {
	m := make(map[string]*ExcelData)
	for _, v := range e {
		if m[v.Symbol] == nil {
			m[v.Symbol].Size += v.Size
		}
		m[v.Symbol] = v
	}
	var newE ExcelDatas
	for _, v := range m {
		newE = append(newE, v)
	}
	return newE
}

func (e ExcelData) ToSlice() []any {
	// TODO
	return []any{}
}
