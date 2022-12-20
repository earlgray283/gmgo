package quinemccluskey

import (
	"fmt"
	"strconv"
	"strings"
)

func All[T any](a []T, f func(t T) bool) bool {
	for _, elem := range a {
		if !f(elem) {
			return false
		}
	}
	return true
}

func GetColumnFrom2d[T any](a [][]T, rowPos int) []T {
	column := make([]T, len(a))
	for i := range a {
		column[i] = a[i][rowPos]
	}
	return column
}

func printTruthTable(a [][]*int) {
	for _, row := range a {
		dispList := []string{}
		for _, col := range row {
			if col == nil {
				dispList = append(dispList, "-")
			} else {
				dispList = append(dispList, strconv.Itoa(*col))
			}
		}
		fmt.Println(strings.Join(dispList, " | "))
	}
}

func printTruthTable2(a [][]int) {
	for _, row := range a {
		dispList := []string{}
		for _, col := range row {
			dispList = append(dispList, strconv.Itoa(col))
		}
		fmt.Println(strings.Join(dispList, " | "))
	}
}
