package calc_test

import (
	"fmt"
	"github.com/shawroger/zcp-go/calc"
)

func ExampleGenerateCardList() {
	// 生成卡组，长度为 5
	r := calc.CardListGenerate(5)

	fmt.Printf("%#v", r)

	// Output:
	// CardList{1,2,3,4,5}
}
