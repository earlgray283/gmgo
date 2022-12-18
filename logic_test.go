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

// example from https://ja.wikipedia.org/wiki/%E3%82%AF%E3%83%AF%E3%82%A4%E3%83%B3%E3%83%BB%E3%83%9E%E3%82%AF%E3%83%A9%E3%82%B9%E3%82%AD%E3%83%BC%E6%B3%95
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
	fmt.Println()

	fmt.Println("===out===")
	printTruthTable(out)
	fmt.Println()

	ans, err := QuineMcCluskey(in, out)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("===optimized-significant-table===")
	printSignificantTable(ans[0])
	fmt.Println()

	assert.Equal(t, len(ans), 1)
	assert.EqualValues(t, []*int{nil, &one, &zero, &zero}, ans[0][0].Significant)
	assert.EqualValues(t, []int{4, 12}, ans[0][0].IndexList)
	assert.EqualValues(t, []*int{&one, &zero, nil, nil}, ans[0][1].Significant)
	assert.EqualValues(t, []int{8, 9, 10, 11}, ans[0][1].IndexList)
	assert.EqualValues(t, []*int{&one, nil, nil, &zero}, ans[0][2].Significant)
	assert.EqualValues(t, []int{8, 10, 12, 14}, ans[0][2].IndexList)
	assert.EqualValues(t, []*int{&one, nil, &one, nil}, ans[0][3].Significant)
	assert.EqualValues(t, []int{10, 11, 14, 15}, ans[0][3].IndexList)
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
