package zcp

// ErgodicBase 遍历基函数
func ErgodicBase(size, width int) CardListGroup {
	var r CardListGroup

	if size == width {
		base := CardListGenerate(width)
		r = append(r, base.Sort())
	} else if size > width {
		base := ErgodicBase(size-1, width)
		for _, value := range base {
			for index := range value {
				var temp CardList
				temp = append(temp, value[:index]...)
				temp = append(temp, value[index+1:]...)
				temp = append(temp, CardPoint(size))
				temp = temp.Sort()
				if !CardListContain(r, temp) {
					r = append(r, temp)
				}
			}
			r = append(r, value)
		}
	}
	return r
}

// ErgodicBase2 遍历基函数，且包含对应剩余分组
func ErgodicBase2(size, width int) Result {
	var r Result
	base := CardListGenerate(size)
	list := ErgodicBase(size, width)
	for _, item := range list {
		var rest CardList

		for _, value := range base {
			if !CardPointContain(item, value) {
				rest = append(rest, value)
			}
		}
		temp := ValidGroup{item.Copy(), rest}
		r = append(r, temp)
	}
	return r
}

// ErgodicCardList 列出所有卡组组合
func ErgodicCardList(p CardList, width int) CardListGroup {
	var r CardListGroup
	base := ErgodicBase(len(p), width)
	for _, list := range base {
		var temp CardList
		for _, val := range list {
			// -1 很重要，因为 ErgodicBase 数组是从 1 开始的
			temp = append(temp, p[val-1])
		}
		if !CardListContain(r, temp) {
			r = append(r, temp)
		}
	}
	return r
}

// ErgodicCardList2
//
// 列出所有卡组组合，且包含对应剩余分组
func ErgodicCardList2(p CardList, width int) Result {
	var r Result
	base := ErgodicBase2(len(p), width)
	for _, list := range base {
		var temp1, temp2 CardList
		for _, val := range list[0] {
			// -1 很重要，因为 ErgodicBase 数组是从 1 开始的
			temp1 = append(temp1, p[val-1])
		}

		for _, val := range list[1] {
			temp2 = append(temp2, p[val-1])
		}
		r = append(r, ValidGroup{temp1, temp2})
	}
	return r
}

// Ergodic 直接遍历组合
func (p CardList) Ergodic(width int) Result {
	return ErgodicCardList2(p, width)
}
