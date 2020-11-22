package calc_test

import (
	"fmt"
	"github.com/shawroger/zcp-go/calc"
	"testing"
)

func TestErgodicBase(t *testing.T) {
	r := calc.ErgodicBase(6, 3)
	for index, item := range r {
		fmt.Printf("%v -> %#v\n", index+1, item)
	}
}

func TestErgodicBase2(t *testing.T) {
	r := calc.ErgodicBase2(5, 3)
	for index, item := range r {
		fmt.Printf("%v -> %#v\n", index+1, item)
	}
}

func TestErgodicCard(t *testing.T) {
	temp := calc.CardList{
		10, 11, 12, 13, 14,
	}
	r := calc.ErgodicCard(temp, 3)
	for index, item := range r {
		fmt.Printf("%v -> %#v\n", index+1, item)
	}

}

func TestErgodicCard2(t *testing.T) {
	temp := calc.CardList{
		10, 11, 12, 13, 14,
	}
	r := calc.ErgodicCard2(temp, 3)
	for index, item := range r {
		fmt.Printf("%v -> %#v\n", index+1, item)
	}

}
