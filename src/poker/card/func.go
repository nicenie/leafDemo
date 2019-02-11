package card

import (
	"fmt"
)

//CGroupString 获取牌组对应的string
func CGroupString(cards []CType) string {
	var result string
	for _, c := range cards {
		result = fmt.Sprintf("%s0x%02x ", result, c)
	}
	return result
}

// CSort 排序升序
func CSort(a, b CType) bool {
	var is bool
	if a.GetDigit() == b.GetDigit() {
		is = a.GetColor() < b.GetColor()
	} else {
		is = a.GetDigit() < b.GetDigit()
	}
	return is
}
