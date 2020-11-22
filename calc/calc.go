package calc

import (
	"math"
)

type ValidGroup [2]CardList

type Result []ValidGroup

// addUtilHalf 判断目标在卡组长度的一半以上
func addUtilHalf(x int, y CardList) bool {
	return float64(x) >= math.Ceil(0.5*float64(len(y)))
}

// RunWithAll 获取由全部卡点组合的结果
func RunWithAll(c CardList) Result {
	var r Result
	for width := len(c); addUtilHalf(width, c); width-- {
		baseList := ErgodicCard2(c, width)

		// 直接获取二分组的结果
		for _, list := range baseList {
			if SumEqual(list[0], list[1]) {
				r = append(r, ValidGroup{list[0], list[1]})
			}
		}
	}

	return r
}

// Run 获取结果
func Run(c CardList) Result {
	var r Result
	for width := len(c); addUtilHalf(width, c); width-- {
		baseList := ErgodicCard2(c, width)

		for _, list := range baseList {

			checker1 := make(map[string]string)
			// 闭包函数
			runAndFilter := func(list CardList) {
				withAll := RunWithAll(list)
				// 消去顺序相反的情况
				for _, item := range withAll {
					str1, exist1 := checker1[item[0].Str()]
					if exist1 && str1 == item[1].Str() {
						continue
					}

					str2, exist2 := checker1[item[1].Str()]
					if exist2 && str2 == item[0].Str() {
						continue
					}

					if len(item[0]) > 0 && len(item[1]) > 0 {
						checker1[item[0].Str()] = item[1].Str()
						r = append(r, item)
					}
				}
			}
			runAndFilter(list[0])
			runAndFilter(list[1])
		}
	}

	return r
}
