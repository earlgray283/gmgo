package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func all[T any](a []T, f func(t T) bool) bool {
	for _, elem := range a {
		if !f(elem) {
			return false
		}
	}
	return true
}

func getColumnFrom2d[T any](a [][]T, rowPos int) []T {
	column := make([]T, len(a))
	for i := range a {
		column[i] = a[i][rowPos]
	}
	return column
}

func openFileAsCsv(path string) ([]string, [][]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return nil, nil, err
	}
	var header []string
	if csvLabelFlag {
		header = records[0]
		records = records[1:]
	}
	return header, records, nil
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
