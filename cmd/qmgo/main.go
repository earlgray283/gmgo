package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	qm "github.com/earlgray283/quine-mccluskey"
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
	rawInputHeader, inputCsv, err := openFileAsCsv(args[0])
	if err != nil {
		log.Fatal(err)
	}
	if os.Getenv("INPUT_LABEL_CSV") != "" {
		rawInputHeader = strings.Split(os.Getenv("INPUT_LABEL_CSV"), ",")
	}
	input, inputRemovedSet, err := parseInputCsv(inputCsv)
	if err != nil {
		log.Fatal(err)
	}
	inputHeader := []string{}
	for index, field := range rawInputHeader {
		if _, ok := inputRemovedSet[index]; !ok {
			inputHeader = append(inputHeader, field)
		}
	}
	outputHeader, outputCsv, err := openFileAsCsv(args[1])
	if err != nil {
		log.Fatal(err)
	}
	if os.Getenv("OUTPUT_LABEL_CSV") != "" {
		outputHeader = strings.Split(os.Getenv("OUTPUT_LABEL_CSV"), ",")
	}
	output, err := parseOutputCsv(outputCsv)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(inputHeader, outputHeader)

	significantGroupEachOutput, err := qm.QuineMcCluskey(input, output)
	if err != nil {
		log.Fatal(err)
	}

	for outputIndex, significantGroup := range significantGroupEachOutput {
		column := qm.GetColumnFrom2d(output, outputIndex)
		first := column[0]
		if qm.All(column, func(item *int) bool {
			if first == nil {
				return item == first
			} else {
				if item == nil {
					return false
				} else {
					return *item == *first
				}
			}
		}) {
			fmt.Println()
			if first == nil {
				fmt.Printf("%s = 0\n", outputHeader[outputIndex])
			} else if *first == 0 {
				fmt.Printf("%s = 0\n", outputHeader[outputIndex])
			} else {
				fmt.Printf("%s = 1\n", outputHeader[outputIndex])
			}
			continue
		}

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

func joinNtimes(s string, n int) string {
	tmp := strings.Builder{}
	for i := 0; i < n; i++ {
		_, _ = tmp.WriteString(s)
	}
	return tmp.String()
}
