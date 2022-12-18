package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

var (
	csvLabelFlag bool
)

func main() {
	flag.BoolVar(&csvLabelFlag, "csv-label", true, "default is true")
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("USAGE: qmgo <input path> <output path> [...options]")
		return
	}
	inputHeader, inputCsv, err := openFileAsCsv(args[0])
	if err != nil {
		log.Fatal(err)
	}
	input, err := parseInputCsv(inputCsv)
	if err != nil {
		log.Fatal(err)
	}
	outputHeader, outputCsv, err := openFileAsCsv(args[1])
	if err != nil {
		log.Fatal(err)
	}
	output, err := parseOutputCsv(outputCsv)
	if err != nil {
		log.Fatal(err)
	}

	significantGroupEachOutput, err := QuineMcCluskey(input, output)
	if err != nil {
		log.Fatal(err)
	}
	for outputIndex, significantGroup := range significantGroupEachOutput {
		expr := []string{}
		notList := []string{}
		for _, optimizedSignificant := range significantGroup[0] {
			significantStr := []string{}
			notToken := []string{}
			for inputIndex, item := range optimizedSignificant.Significant {
				if item != nil {
					significantStr = append(significantStr, inputHeader[inputIndex])
					if *item == 1 {
						notToken = append(notToken, joinNtimes(" ", len(inputHeader[inputIndex])))
					} else {
						notToken = append(notToken, joinNtimes("_", len(inputHeader[inputIndex])))
					}
				}
			}
			expr = append(expr, strings.Join(significantStr, " * "))
			notList = append(notList, strings.Join(notToken, "   "))
		}
		fmt.Printf("    %s\n", strings.Join(notList, "   "))
		fmt.Printf("%s = %s\n", outputHeader[outputIndex], strings.Join(expr, " + "))
	}
}
