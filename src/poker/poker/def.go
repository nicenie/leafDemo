package poker

import (
	"poker/card"
)

// ShunziT 顺子在single-cards中的起始位置和长度
type ShunziT struct {
	Len   int
	Cards card.CSlice
}

// NewShunziT 构造函数
func NewShunziT() ShunziT {
	return ShunziT{
		Len:   0,
		Cards: make(card.CSlice, 0),
	}
}
