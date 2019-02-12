package poker

import (
	"poker/card"
	"sort"
)

//Util 扑克牌组工具类
type Util struct {
}

//GetSameX 从cards牌组中获取牌组数目等于targetNum的牌组
func (u *Util) GetSameX(cards card.CSlice, targetNum int) (map[card.CDigit]card.CSlice, int) {
	tmpMap := make(map[card.CDigit]card.CSlice)
	for _, v := range cards {
		if _, ok := tmpMap[v.GetDigit()]; ok {
			tmpMap[v.GetDigit()] = append(tmpMap[v.GetDigit()], v)
			// fmt.Printf("add Card to exist group. key=0x%02x|val=0x%02x|len=%d\n", v.GetDigit(), v, len(tmpMap[v.GetDigit()]))
		} else {
			tmpSlice := make(card.CSlice, 0, 4)
			tmpSlice = append(tmpSlice, v)
			tmpMap[v.GetDigit()] = tmpSlice
			// fmt.Printf("add new group key=0x%02x\n", v.GetDigit())
		}
	}

	for _, val := range tmpMap {
		if count := len(val); count > 0 && count != targetNum {
			// fmt.Printf("delete num=%d digit=0x%02x\n", count, val[0].GetDigit())
			delete(tmpMap, val[0].GetDigit())
		}
	}
	return tmpMap, len(tmpMap)
}

//GetSameXEx 从cards牌组中获取牌组数目等于targetNum的牌组,结果放在同一个slice里
func (u *Util) GetSameXEx(cards card.CSlice, targetNum int) (result card.CSlice, count int) {
	tmpRst, _ := u.GetSameX(cards, targetNum)
	for _, group := range tmpRst {
		result = append(result, group...)
	}
	count = len(result)
	return
}

//GetSameEqX 从牌组中获取targetNum张相同的牌，尽管牌组中该牌的数量多于targetNum
func (u *Util) GetSameEqX(cards card.CSlice, targetNum int) (map[card.CDigit]card.CSlice, int) {
	tmpMap := make(map[card.CDigit]card.CSlice)
	for _, v := range cards {
		if _, ok := tmpMap[v.GetDigit()]; ok {
			tmpMap[v.GetDigit()] = append(tmpMap[v.GetDigit()], v)
		} else {
			tmpSlice := make(card.CSlice, 0, 4)
			tmpSlice = append(tmpSlice, v)
			tmpMap[v.GetDigit()] = tmpSlice
		}
	}

	for _, val := range tmpMap {
		if count := len(val); count >= targetNum {
			if count > targetNum {
				if tmp, ok := tmpMap[val[0].GetDigit()]; ok && targetNum > 0 {
					tmptmp := tmp[0:targetNum]
					tmpMap[val[0].GetDigit()] = tmptmp
				}
			}
		} else {
			if count > 0 {
				delete(tmpMap, val[0].GetDigit())
			}
		}
	}
	return tmpMap, len(tmpMap)
}

//GetSameEqXEx 从cards牌组中获取牌组数目等于targetNum的牌组,结果放在同一个slice里
func (u *Util) GetSameEqXEx(cards card.CSlice, targetNum int) (result card.CSlice, count int) {
	tmpRst, _ := u.GetSameEqX(cards, targetNum)
	for _, group := range tmpRst {
		result = append(result, group...)
	}
	count = len(result)
	return
}

// GetXCountVal 从牌组cards中获取count张val的牌
func (u *Util) GetXCountVal(cards card.CSlice, val card.CDigit, count int) (result card.CSlice) {
	for _, c := range cards {
		if len(result) < count && val == c.GetDigit() {
			result = append(result, c)
		}
		if len(result) == count {
			break
		}
	}
	return
}

//IsContainsX 牌组cards中是否包含有牌tc
func (u *Util) IsContainsX(cards card.CSlice, tc card.CType) bool {
	for _, c := range cards {
		if c == tc {
			return true
		}
	}
	return false
}

//IsContainsXVal 牌组cards中是否包含有牌值td
func (u *Util) IsContainsXVal(cards card.CSlice, td card.CDigit) bool {
	for _, c := range cards {
		if c.GetDigit() == td {
			return true
		}
	}
	return false
}

//GetSingles 获取牌组的单牌牌组
func (u *Util) GetSingles(cards card.CSlice) (result card.CSlice, count int) {
	tmp, tc := u.GetSameEqXEx(cards, 1)
	count = tc
	result = append(result, tmp...)
	return
}

//GetMoreBigThanXVal 获取牌组中比td大的牌组
func (u *Util) GetMoreBigThanXVal(cards card.CSlice, td card.CDigit) (result card.CSlice) {
	for _, c := range cards {
		if c.GetDigit() > td {
			result = append(result, c)
		}
	}
	return
}

//GetExcepts 从牌组cards获取的牌组不包含tcards的牌组,按照real-value比较
func (u *Util) GetExcepts(cards card.CSlice, tcards card.CSlice) (result card.CSlice) {
	for _, c := range cards {
		if !u.IsContainsX(tcards, c) {
			result = append(result, c)
		}
	}
	return
}

//GetExceptsXVals 从牌组cards获取的牌组不包含tcards的牌组,按照Digit比较
func (u *Util) GetExceptsXVals(cards card.CSlice, tcards card.CSlice) (result card.CSlice) {
	for _, c := range cards {
		if !u.IsContainsXVal(tcards, c.GetDigit()) {
			result = append(result, c)
		}
	}
	return
}

//GetMaxXShun 获取cards牌组中最长的X顺子
func (u *Util) GetMaxXShun(cards card.CSlice, x, limit int) (result card.CSlice) {
	ers := []card.CType{card.CType(card.CardErVal)}
	exceptEr := u.GetExceptsXVals(cards, ers)
	thanLimit, _ := u.GetSameEqXEx(exceptEr, x)
	// fmt.Print("thanLimit:")
	// fmt.Println(card.CGroupString(thanLimit))
	singles, _ := u.GetSameEqXEx(thanLimit, 1)
	sort.Sort(singles)
	// fmt.Print("singles:")
	// fmt.Println(card.CGroupString(singles))
	sLen := len(singles)
	var index, startIndex int = 0, -1
	tmpShunzis := make([]ShunziT, 0)
	tmpSz := make(card.CSlice, 0, len(singles))
	for index < sLen-1 {
		if startIndex < 0 {
			startIndex = index
		}
		var ok, ok1 bool
		if singles[index].GetDigit()+1 == singles[index+1].GetDigit() {
			tmpSz = append(tmpSz, singles[index])
			if len(tmpSz) == sLen-1 {
				tmpSz = append(tmpSz, singles[index+1])
				ok = true
			}
		} else { //broken
			ok1 = true
		}
		if ok || ok1 {
			if !ok {
				tmpSz = append(tmpSz, singles[index])
			}
			szT := NewShunziT()
			szT.Len = index - startIndex + 1
			if szT.Len >= limit {
				szT.Cards = append(szT.Cards, tmpSz...)
				tmpShunzis = append(tmpShunzis, szT)
				// fmt.Printf("tmpShunzi:%s\n", card.CGroupString(szT.Cards))
			}
			startIndex = -1
			tmpSz = make(card.CSlice, 0, len(singles))
		}
		index++
	}
	index = 0
	var maxLen int
	for idx, sz := range tmpShunzis {
		if maxLen <= 0 {
			maxLen = sz.Len
			index = idx
		}
		if sz.Len > maxLen {
			maxLen = sz.Len
			index = idx
		}
	}
	if index >= 0 && index < len(tmpShunzis) && maxLen >= limit {
		for _, c := range tmpShunzis[index].Cards {
			tmp := u.GetXCountVal(cards, c.GetDigit(), x)
			result = append(result, tmp...)
		}
	}
	return
}
