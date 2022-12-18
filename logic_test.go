package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	zero = 0
	one  = 1
)

func Test_QuineMcCluskey(t *testing.T) {
	in := [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 1},
		{0, 0, 1, 0},
		{0, 0, 1, 1},
		{0, 1, 0, 0},
		{0, 1, 0, 1},
		{0, 1, 1, 0},
		{0, 1, 1, 1},
		{1, 0, 0, 0},
		{1, 0, 0, 1},
		{1, 0, 1, 0},
		{1, 0, 1, 1},
		{1, 1, 0, 0},
		{1, 1, 0, 1},
		{1, 1, 1, 0},
		{1, 1, 1, 1},
	}
	out := [][]*int{
		{&zero},
		{&zero},
		{&zero},
		{&zero},
		{&one},
		{&zero},
		{&zero},
		{&zero},
		{&one},
		{nil},
		{&one},
		{&one},
		{&one},
		{&zero},
		{nil},
		{&one},
	}

	fmt.Println("===in===")
	printTruthTable2(in)
	fmt.Println("===out===")
	printTruthTable(out)

	ans, err := QuineMcCluskey(in, out)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("===optimized-significant-table===")
	printSignificantTable(ans[0])
}

func Test_calcHammingDistance(t *testing.T) {
	dist := calcHammingDistance([]*int{&one, &zero, &zero, nil}, []*int{&one, &zero, &one, nil})
	assert.Equal(t, 1, dist)

	dist = calcHammingDistance([]*int{&one, &zero, nil, &zero}, []*int{&one, &zero, &one, nil})
	assert.Equal(t, 2, dist)
}

func Test_combination(t *testing.T) {
	var intNil *int = nil

	combi := combination([]*int{&one, &zero, &zero, nil}, []*int{&one, &zero, &one, nil})
	assert.Equal(t, *combi[0], 1)
	assert.Equal(t, *combi[1], 0)
	assert.Equal(t, combi[2], intNil)
	assert.Equal(t, combi[3], intNil)

	combi = combination([]*int{&zero, &one, &zero, &zero}, []*int{&one, &one, &zero, &zero})
	assert.Equal(t, combi[0], intNil)
	assert.Equal(t, combi[1], &one)
	assert.Equal(t, combi[2], &zero)
	assert.Equal(t, combi[3], &zero)
}
