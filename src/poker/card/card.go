package card

//CColor 扑克牌花色类型
type CColor uint8

//CDigit 扑克牌值类型
type CDigit uint8

//CType 扑克牌类型
type CType uint8

//GetColor 获取牌的花色
func (c CType) GetColor() (col CColor) {
	col = CColor(c & 0xf0)
	return
}

//GetDigit 获取牌的数值
func (c CType) GetDigit() (val CDigit) {
	val = CDigit(c & 0x0f)
	return
}

/* ********************************************************************** */

//CSlice 牌组容器
type CSlice []CType

// Len 长度
func (cs CSlice) Len() int {
	return len(cs)
}

// Swap 交换i,j位置数据
func (cs CSlice) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}

// Less 排序（升序）
func (cs CSlice) Less(i, j int) bool {
	var is bool
	if cs[i].GetDigit() == cs[j].GetDigit() {
		is = cs[i].GetColor() < cs[j].GetColor()
	} else {
		is = cs[i].GetDigit() < cs[j].GetDigit()
	}
	return is
}
