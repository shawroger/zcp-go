package zcp

import (
	"fmt"
)

func ExampleGenerateCardList() {
	// 生成卡组，长度为 5
	r := CardListGenerate(5)

	fmt.Printf("%#v", r)

	// Output:
	// CardList{1,2,3,4,5}
}
