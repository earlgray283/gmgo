package quinemccluskey

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
func Test_quineMcCluskeyWith1out(t *testing.T) {
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

	mustList, optionalList := quineMcCluskeyWith1out(in, GetColumnFrom2d(out, 0))

	fmt.Println("===must-significant-table===")
	printSignificantTable(mustList)
	fmt.Println()
	fmt.Println("===optional-significant-table===")
	printSignificantTable(optionalList)
	fmt.Println()

	assert.Equal(t, 2, len(mustList))
	assert.EqualValues(t, []*int{nil, &one, &zero, &zero}, mustList[0].Significant)
	assert.EqualValues(t, []int{4, 12}, mustList[0].IndexList)
	assert.EqualValues(t, []*int{&one, nil, &one, nil}, mustList[1].Significant)
	assert.EqualValues(t, []int{10, 11, 14, 15}, mustList[1].IndexList)

	assert.Equal(t, 2, len(optionalList))
	assert.EqualValues(t, []*int{&one, &zero, nil, nil}, optionalList[0].Significant)
	assert.EqualValues(t, []int{8, 9, 10, 11}, optionalList[0].IndexList)
	assert.EqualValues(t, []*int{&one, nil, nil, &zero}, optionalList[1].Significant)
	assert.EqualValues(t, []int{8, 10, 12, 14}, optionalList[1].IndexList)
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
