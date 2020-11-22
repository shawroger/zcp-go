package zcp

import (
	"fmt"
	"testing"
)

func TestErgodicBase(t *testing.T) {
	r := ErgodicBase(6, 3)
	for index, item := range r {
		fmt.Printf("%v -> %#v\n", index+1, item)
	}
}

func TestErgodicBase2(t *testing.T) {
	r := ErgodicBase2(5, 3)
	for index, item := range r {
		fmt.Printf("%v -> %#v\n", index+1, item)
	}
}

func TestErgodicCard(t *testing.T) {
	temp := CardList{
		10, 11, 12, 13, 14,
	}
	r := ErgodicCardList(temp, 3)
	for index, item := range r {
		fmt.Printf("%v -> %#v\n", index+1, item)
	}

}

func TestErgodicCard2(t *testing.T) {
	temp := CardList{
		10, 11, 12, 13, 14,
	}
	r := ErgodicCardList2(temp, 3)
	for index, item := range r {
		fmt.Printf("%v -> %#v\n", index+1, item)
	}

}

func TestCardList_Ergodic(t *testing.T) {
	list := New(1, 2, 4, 5, 6, 14)
	list.Run().Print()
}
