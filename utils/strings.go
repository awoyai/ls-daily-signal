package utils

import (
	"sort"
	"strings"
	"unicode"
)

func Contain[T string | int | float64](arr []T, x T) bool {
	switch any(x).(type) {
	case string:
		sort.Strings(any(arr).([]string))
	case int:
		sort.Ints(any(arr).([]int))
	case float64:
		sort.Float64s(any(arr).([]float64))
	}

	index := sort.Search(len(arr), func(i int) bool { return arr[i] >= x })

	return index < len(arr) && arr[index] == x
}

// ParseVarietyFromSymbol 对symbol不校验，symbol必须类似于RU2205
func ParseVarietyFromSymbol(symbol string) string {
	if symbol == "" {
		return symbol
	}

	var end int
	for _, c := range symbol {
		if unicode.IsLetter(c) {
			end++
			continue
		}
		if unicode.IsNumber(c) { // 遇到数字就break
			break
		}
	}
	
	return strings.ToLower(symbol[:end])
}


