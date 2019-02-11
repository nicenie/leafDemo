package poker

import (
	"fmt"
	"os"
	"poker/card"
	"sort"
	"testing"
)

var util Util

func TestGetSameX(t *testing.T) {
	cards := card.CSlice{0x01, 0x11, 0x21, 0x31, 0x03}
	result, count := util.GetSameX(cards, 4)
	fmt.Printf("count=%d\n", count)
	for _, group := range result {
		fmt.Println(card.CGroupString(group))
	}

	if count != 1 {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

func TestGetSameEqX(t *testing.T) {
	cards := card.CSlice{0x01, 0x11, 0x21, 0x31, 0x03}
	result, count := util.GetSameEqX(cards, 3)
	fmt.Printf("count=%d\n", count)
	for _, group := range result {
		fmt.Println(card.CGroupString(group))
	}

	if count != 1 {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

func TestGetSameEqXEx(t *testing.T) {
	cards := card.CSlice{0x01, 0x11, 0x21, 0x31, 0x03}
	result, count := util.GetSameEqXEx(cards, 3)
	fmt.Printf("count=%d\n", count)
	fmt.Println(card.CGroupString(result))

	if count <= 0 {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

func TestGetExcepts(t *testing.T) {
	cards := card.CSlice{0x05, 0x15, 0x04, 0x08, 0x01, 0x11, 0x21, 0x31, 0x03}
	sort.Sort(cards)
	fmt.Println(card.CGroupString(cards))
	ecards := card.CSlice{0x01, 0x03}
	result := util.GetExcepts(cards, ecards)
	fmt.Println(card.CGroupString(result))
	if len(result) == len(cards) {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
