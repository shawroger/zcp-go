package zcp

import (
	"fmt"
	"testing"
)

type testContentEqualGroup struct {
	p1     CardList
	p2     CardList
	should bool
}

func TestContentEqual(t *testing.T) {
	testGroup := []testContentEqualGroup{
		{
			p1:     CardList{1, 2, 3, 4, 5},
			p2:     CardList{1, 4, 2, 3, 5},
			should: true,
		},
		{
			p1:     CardList{1, 2, 3, 4, 5},
			p2:     CardList{1, 4, 3, 3, 5},
			should: false,
		},
		{
			p1:     CardList{1, 2, 3, 4, 5},
			p2:     CardList{3, 4, 1, 2, 5},
			should: true,
		},
		{
			p1:     CardList{1, 2, 3, 4, 5},
			p2:     CardList{2, 1, 2, 5},
			should: false,
		},
	}

	for index, group := range testGroup {
		testResult := ContentEqual(group.p1, group.p2)
		if testResult != group.should {
			t.Errorf("Test failed at group %d", index)
		}
	}
}

func TestCardList_Copy(t *testing.T) {
	p := CardList{1, 5, 2}
	c := p.Sort()
	c[1] = 3

	fmt.Printf("%#v\n%#v\n", p, c)
}
