package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getColumnFrom2d(t *testing.T) {
	a := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	colList := getColumnFrom2d(a, 3)
	assert.Equal(t, colList[0], 4)
	assert.Equal(t, colList[1], 8)
	assert.Equal(t, colList[2], 12)
	assert.Equal(t, colList[3], 16)
}
