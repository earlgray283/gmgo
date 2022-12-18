package main

import (
	"strconv"

	"github.com/pkg/errors"
)

func parseInputCsv(records [][]string) ([][]int, error) {
	inputTable := [][]int{}
	for _, record := range records {
		row := []int{}
		for _, cell := range record {
			v, err := strconv.Atoi(cell)
			if err != nil {
				return nil, errors.Wrap(err, "don't care is currently not allowed:(")
			}
			if v != 0 && v != 1 {
				v = 1
			}
			row = append(row, v)
		}
		inputTable = append(inputTable, row)
	}
	return inputTable, nil
}

func parseOutputCsv(records [][]string) ([][]*int, error) {
	outputTable := [][]*int{}
	for _, record := range records {
		row := []*int{}
		for _, cell := range record {
			v, err := strconv.Atoi(cell)
			if err != nil {
				row = append(row, nil)
				continue
			}
			if v != 0 && v != 1 {
				v = 1
			}
			row = append(row, &v)
		}
		outputTable = append(outputTable, row)
	}
	return outputTable, nil
}
