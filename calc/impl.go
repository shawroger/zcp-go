package calc

import (
	"fmt"
	"sort"
)

type CardListFeature interface {
	Sum() CardPoint
	SumEqual(t CardList) bool
	ContentEqual(t CardList) bool
	Copy() CardList
	Sort() CardList
	Str() string
	sort.Interface
}

// Sum 卡牌组求和
func (p CardList) Sum() CardPoint {
	return SumList(p)
}

// SumEqual 判断卡牌和相等
func (p CardList) SumEqual(t CardList) bool {
	return SumList(p) == SumList(t)
}

// ContentEqual 判断卡牌内容相等
func (p CardList) ContentEqual(t CardList) bool {
	return ContentEqual(p, t)
}

// Copy 复制卡组
func (p CardList) Copy() CardList {
	return CardListCopy(p)
}

func (p CardList) Len() int {
	return len(p)
}

func (p CardList) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p CardList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Sort 对自己排序
//
// 在 Copy 基础上，不修改原数据
func (p CardList) Sort() CardList {
	copy := p.Copy()
	sort.Sort(copy)

	return copy
}

// Str 转字符串
func (p CardList) Str() string {
	return fmt.Sprintf("%v", p)
}
