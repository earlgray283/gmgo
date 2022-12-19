package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var (
	csvLabelFlag bool
)

func init() {
	_ = godotenv.Load(".env")
}

func main() {
	flag.BoolVar(&csvLabelFlag, "csv-label", false, "default is false")
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
	if os.Getenv("INPUT_LABEL_CSV") != "" {
		inputHeader = strings.Split("INPUT_LABEL_CSV", ",")
	}
	input, inputRemovedList, err := parseInputCsv(inputCsv)
	if err != nil {
		log.Fatal(err)
	}
	for _, index := range inputRemovedList {
		inputHeader = append(inputHeader[:index], inputHeader[index+1:]...)
	}
	outputHeader, outputCsv, err := openFileAsCsv(args[1])
	if err != nil {
		log.Fatal(err)
	}
	if os.Getenv("OUTPUT_LABEL_CSV") != "" {
		outputHeader = strings.Split("OUTPUT_LABEL_CSV", ",")
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
		if len(significantGroup) == 0 {
			fmt.Println()
			fmt.Printf("%s = 0\n", outputHeader[outputIndex])
			continue
		}

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
