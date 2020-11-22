package calc_test

import (
	"fmt"
	"github.com/shawroger/zcp-go/calc"
	"testing"
)

type testContentEqualGroup struct {
	p1     calc.CardList
	p2     calc.CardList
	should bool
}

func TestContentEqual(t *testing.T) {
	testGroup := []testContentEqualGroup{
		{
			p1:     calc.CardList{1, 2, 3, 4, 5},
			p2:     calc.CardList{1, 4, 2, 3, 5},
			should: true,
		},
		{
			p1:     calc.CardList{1, 2, 3, 4, 5},
			p2:     calc.CardList{1, 4, 3, 3, 5},
			should: false,
		},
		{
			p1:     calc.CardList{1, 2, 3, 4, 5},
			p2:     calc.CardList{3, 4, 1, 2, 5},
			should: true,
		},
		{
			p1:     calc.CardList{1, 2, 3, 4, 5},
			p2:     calc.CardList{2, 1, 2, 5},
			should: false,
		},
	}

	for index, group := range testGroup {
		testResult := calc.ContentEqual(group.p1, group.p2)
		if testResult != group.should {
			t.Errorf("Test failed at group %d", index)
		}
	}
}

func TestCardList_Copy(t *testing.T) {
	p := calc.CardList{1, 5, 2}
	copy := p.Sort()
	copy[1] = 3

	fmt.Printf("%#v\n%#v\n", p, copy)
}
