package quinemccluskey

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)



func spreadDontCare(in, out [][]*int) ([][]int, [][]*int) {
	for {
		newIn := [][]*int{}
		newOut := [][]*int{}
		containsDontCare := false
		for rowIndex, row := range in {
			appended := false
			for colIndex, col := range row {
				if col == nil {
					newRow0, newRow1 := slices.Clone(row), slices.Clone(row)
					newRow0[colIndex], newRow1[colIndex] = pointerOfValue(0), pointerOfValue(1)
					newIn = append(newIn, newRow0, newRow1)
					newOut = append(newOut, out[rowIndex], out[rowIndex])
					containsDontCare = true
					appended = true
				}
			}
			if !appended {
				newIn = append(newIn, slices.Clone(row))
				newOut = append(newOut, out[rowIndex])
			}
		}
		if !containsDontCare {
			return lo.Map(newIn, func(row []*int, _ int) []int {
				return lo.Map(row, func(col *int, _ int) int { return *col })
			}), newOut
		}
		in, out = newIn, newOut
	}
}

// nil は don't care として扱う
func QuineMcCluskey(rawIn [][]*int, rawOut [][]*int) ([][][]SignificantGroup, error) {
	in, out := spreadDontCare(rawIn, rawOut)

	if len(in) != len(out) {
		return nil, errors.New("len(in) must be equal to len(out)")
	}
	n := len(in)
	if n == 0 {
		return nil, nil
	}
	if !All(in, func(t []int) bool { return len(t) == len(in[0]) }) || !All(out, func(t []*int) bool { return len(t) == len(out[0]) }) {
		return nil, errors.New("")
	}
	_, outM := len(in[0]), len(out[0])

	table := make([][][]SignificantGroup, outM)
	// out の列ごとに Quine-McCluskey をやる
	for i := 0; i < outM; i++ {
		mustSignificantList, optionalSignificantList := quineMcCluskeyWith1out(in, GetColumnFrom2d(out, i))
		var optimizedSignificantList [][]SignificantGroup
		if len(optionalSignificantList) != 0 {
			optimizedSignificantList = make([][]SignificantGroup, len(optionalSignificantList))
			for j, optionalSignificant := range optionalSignificantList {
				optimizedSignificantList[j] = append(optimizedSignificantList[j], mustSignificantList...)
				optimizedSignificantList[j] = append(optimizedSignificantList[j], optionalSignificant)
			}
		} else if len(mustSignificantList) != 0 {
			optimizedSignificantList = append(optimizedSignificantList, []SignificantGroup{})
			optimizedSignificantList[0] = append(optimizedSignificantList[0], mustSignificantList...)
		} else {
			optimizedSignificantList = append(optimizedSignificantList, []SignificantGroup{})
			for j, row := range in {
				if out[j][i] == nil {
					continue
				}
				if *out[j][i] == 1 {
					optimizedSignificantList[0] = append(optimizedSignificantList[0],
						SignificantGroup{
							Significant: lo.Map(row, func(item int, _ int) *int { return &item }),
						})
				}
			}
		}
		table[i] = optimizedSignificantList
	}

	return table, nil
}

var dfsCount int

type SignificantGroup struct {
	Significant []*int
	IndexList   []int
}

func compareSignificantGroup(a, b SignificantGroup) bool {
	aID := strings.Join(lo.Map(a.IndexList, func(item, _ int) string { return strconv.Itoa(item) }), ",")
	bID := strings.Join(lo.Map(b.IndexList, func(item, _ int) string { return strconv.Itoa(item) }), ",")
	if len(aID) != len(bID) {
		return len(aID) < len(bID)
	} else {
		return aID < bID
	}
}

func dfs(significantGroupList []SignificantGroup) []SignificantGroup {
	dfsCount++
	if dfsCount == 100000 {
		panic("too many call of dfs")
	}

	newSignificantGroupListByID := map[string]SignificantGroup{} // 項
	for i := 0; i < len(significantGroupList)-1; i++ {
		for j := i + 1; j < len(significantGroupList); j++ {
			if calcHammingDistance(significantGroupList[i].Significant, significantGroupList[j].Significant) != 1 {
				continue
			}
			significantGroup := SignificantGroup{
				Significant: combination(significantGroupList[i].Significant, significantGroupList[j].Significant),
				IndexList:   append(significantGroupList[i].IndexList, significantGroupList[j].IndexList...),
			}
			slices.Sort(significantGroup.IndexList)
			id := strings.Join(lo.Map(significantGroup.IndexList, func(item, _ int) string { return strconv.Itoa(item) }), ",")
			newSignificantGroupListByID[id] = significantGroup
		}
	}
	newSignificantGroupList := lo.Values(newSignificantGroupListByID)

	// 単一項(index)が所属している組み合わせ項の index の list
	belongListByIndex := map[int][]int{}
	for i, significantGroup := range newSignificantGroupList {
		for _, index := range significantGroup.IndexList {
			belongListByIndex[index] = append(belongListByIndex[index], i)
		}
	}

	mustSignificantGroupListByIndex := map[int]SignificantGroup{}
	for index, belongList := range belongListByIndex {
		_, exists := mustSignificantGroupListByIndex[index]
		if len(belongList) == 1 && !exists {
			mustSignificantGroupListByIndex[index] = newSignificantGroupList[belongList[0]]
		}
	}
	mustSignificantGroupList := lo.Values(mustSignificantGroupListByIndex)
	slices.SortFunc(mustSignificantGroupList, compareSignificantGroup)

	// 必須項のみだったらこれ以上再帰をする必要がないので return
	if len(mustSignificantGroupList) == len(newSignificantGroupList) {
		return mustSignificantGroupList
	}

	mustSignificantList2 := dfs(newSignificantGroupList)

	return append(mustSignificantGroupList, mustSignificantList2...)
}

// return must, optional
func quineMcCluskeyWith1out(in [][]int, out []*int) ([]SignificantGroup, []SignificantGroup) {
	// out が1となるような input の行を追加していく
	significantList := make([]SignificantGroup, 0)
	dontCareFlagList := map[int]bool{}
	for i := 0; i < len(in); i++ {
		if out[i] == nil {
			significantList = append(significantList, SignificantGroup{
				Significant: lo.Map(in[i], func(a int, _ int) *int {
					return &a
				}),
				IndexList: []int{i},
			})
			dontCareFlagList[i] = true
		} else if *out[i] == 1 {
			significantList = append(significantList, SignificantGroup{
				Significant: lo.Map(in[i], func(a int, _ int) *int {
					return &a
				}),
				IndexList: []int{i},
			})
		}
	}

	significantGroupList := dfs(significantList)
	belongListByIndex := map[int][]int{}
	for i, significantGroup := range significantGroupList {
		for _, index := range significantGroup.IndexList {
			if dontCareFlagList[index] {
				continue
			}
			belongListByIndex[index] = append(belongListByIndex[index], i)
		}
	}
	mustSignificantGroupList := make([]SignificantGroup, 0)
	mustFlagList := map[int]bool{}
	for _, belongList := range belongListByIndex {
		if len(belongList) == 1 {
			mustSignificantGroupList = append(mustSignificantGroupList, significantGroupList[belongList[0]])
			mustFlagList[belongList[0]] = true
		}
	}
	optionalSignificantGroupList := make([]SignificantGroup, 0)
	for index, significantGroup := range significantGroupList {
		if !mustFlagList[index] {
			optionalSignificantGroupList = append(optionalSignificantGroupList, significantGroup)
		}
	}

	slices.SortFunc(mustSignificantGroupList, compareSignificantGroup)
	slices.SortFunc(optionalSignificantGroupList, compareSignificantGroup)

	return mustSignificantGroupList, optionalSignificantGroupList
}

func calcHammingDistance(a, b []*int) int {
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] == nil && b[i] == nil {
			continue
		} else if (a[i] == nil && b[i] != nil) || (a[i] != nil && b[i] == nil) {
			count++
		} else if *a[i] != *b[i] {
			count++
		}
	}
	return count
}

// 0101, 0111 → 01-1
func combination(a, b []*int) []*int {
	combi := make([]*int, len(a))
	for i := 0; i < len(a); i++ {
		if (a[i] == nil && b[i] != nil) || (a[i] != nil && b[i] == nil) {
			panic("pattern -0, -1, 0-, 1- are not allowed")
		}
		if a[i] == nil && b[i] == nil {
			combi[i] = nil
		} else if *a[i] != *b[i] {
			combi[i] = nil
		} else {
			combi[i] = a[i]
		}
	}
	return combi
}

func (s SignificantGroup) String() string {
	dispList := lo.Map(s.Significant, func(item *int, _ int) string {
		if item == nil {
			return "-"
		} else {
			return strconv.Itoa(*item)
		}
	})
	indexList := lo.Map(s.IndexList, func(index int, _ int) string {
		return strconv.Itoa(index)
	})

	return fmt.Sprintln(strings.Join(dispList, " | "), "("+strings.Join(indexList, ", ")+")")
}
