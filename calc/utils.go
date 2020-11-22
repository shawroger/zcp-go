package calc

// Sum 一系列点数求和
func Sum(p ...CardPoint) CardPoint {
	r := p0
	for _, value := range p {
		r += value
	}
	return r
}

// SumList 卡牌组求和
func SumList(p CardList) CardPoint {
	r := p0
	for _, value := range p {
		r += value
	}
	return r
}

// SumEqual 判断卡牌和相等
func SumEqual(p1 CardList, p2 CardList) bool {
	return SumList(p1) == SumList(p2)
}

// ContentEqual 判断卡牌内容相等
func ContentEqual(p1 CardList, p2 CardList) bool {

	if len(p1) != len(p2) {
		return false
	}

	for _, value := range p1 {
		for index, same := range p2 {
			if value == same {
				break
			}

			if index == len(p2)-1 {
				return false
			}
		}
	}

	return true
}

// CardListGenerate 生成卡组序列
func CardListGenerate(length int) CardList {
	r := CardList{}
	for i := 1; i <= length; i++ {
		r = append(r, CardPoint(i))
	}

	return r
}

// CardListCopy 复制卡组
func CardListCopy(p CardList) CardList {
	r := CardList{}
	for _, value := range p {
		r = append(r, value)
	}
	return r
}

// CardListContain 检查是否有重复卡组
func CardListContain(list []CardList, p CardList) bool {
	set := make(map[string]string)
	for _, c := range list {
		set[c.Str()] = c.Str()
	}
	if _, exist := set[p.Str()]; exist {
		return true
	}
	return false
}

// CardPointContain 检查是否有重复卡牌点数
func CardPointContain(list CardList, p CardPoint) bool {
	set := make(map[CardPoint]CardPoint)
	for _, c := range list {
		set[c] = c
	}
	if _, exist := set[p]; exist {
		return true
	}
	return false
}
