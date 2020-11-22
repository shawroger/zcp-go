package zcp

import "reflect"

// CardPoint 卡牌点数
type CardPoint int8

// CardList 卡牌点数组
type CardList []CardPoint

// CardListGroup 卡牌点数组组合
type CardListGroup []CardList

func (c CardListGroup) Print() {
	for _, list := range c {
		list.Print()
	}
}

// makeCardPoint 判断类型可否转为 CardPoint 型
func makeCardPoint(x interface{}) (CardPoint, bool) {
	kind := reflect.TypeOf(x).Kind()
	value := reflect.ValueOf(x)
	if kind == reflect.Ptr {
		kind = reflect.TypeOf(x).Elem().Kind()
		value = value.Elem()
	}

	if kind == reflect.Int ||
		kind == reflect.Int8 ||
		kind == reflect.Int16 ||
		kind == reflect.Int32 ||
		kind == reflect.Int64 ||
		kind == reflect.Uint ||
		kind == reflect.Uint8 ||
		kind == reflect.Uint16 ||
		kind == reflect.Uint32 ||
		kind == reflect.Uint64 {
		return CardPoint(value.Int()), true
	}

	return p0, false
}

// NewFromSlice 从切片新建卡组
func NewFromSlice(x []interface{}) CardList {
	r := CardList{}
	for _, i := range x {
		value, ok := makeCardPoint(i)
		if ok {
			r = append(r, value)
		}
	}
	return r.Sort()
}

// New 从空接口新建卡组
func New(x ...interface{}) CardList {
	return NewFromSlice(x)
}

// 枚举卡牌点数 A~K
// p0 点数为 0 不导出
const (
	p0 CardPoint = iota
	//P1  CardPoint = iota
	//P2  CardPoint = iota
	//P3  CardPoint = iota
	//P4  CardPoint = iota
	//P6  CardPoint = iota
	//P7  CardPoint = iota
	//P8  CardPoint = iota
	//P9  CardPoint = iota
	//P10 CardPoint = iota
	//P11 CardPoint = iota
	//P12 CardPoint = iota
	//P13 CardPoint = iota
)
