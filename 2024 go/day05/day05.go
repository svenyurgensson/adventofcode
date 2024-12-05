package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var ordRules = make(map[int][]int)
var updRules [][]int
var correct, incorrect []int

func parseUpdateRules(data string) {
	re := regexp.MustCompile("(?m)^\n")
	split := re.Split(data, -1)
	r := strings.Split(split[1], "\n")
	updRules = make([][]int, len(r))
	for i, el := range r {
		row := strings.Split(el, ",")
		updRules[i] = make([]int, len(row))
		for j, elm := range row {
			k, err := strconv.Atoi(elm)
			if err != nil { break }
			updRules[i][j] = k
		}
	}
}

// creates map[29:[13] 47:[53 13 61 29] .... ]
func parsePageOrderingRules(data string) {
	re := regexp.MustCompile("(?m)^\n")
	split := re.Split(data, -1)
	r1 := strings.Split(split[0], "\n")
	ordRules = make(map[int][]int)
	for _, el := range r1 {
		parts := strings.Split(el, "|")
		k, err := strconv.Atoi(parts[0])
		if err != nil { break }
		v, err := strconv.Atoi(parts[1])
		if err != nil { panic("Wrong integer value!") }
		ordRules[k] = append(ordRules[k], v)
	}
}

func isOrderRulesIncludes(k int, pages []int) (int, bool) {
	orig, ok := ordRules[k]
	if !ok { return 0, false }

	for idx, el := range pages {
		if !slices.Contains(orig, el) {
			return idx, false
		}
	}
	return 0, true
}

func idxSplitByCorrectUpdateRows() {
	correct = make([]int, 0)
	incorrect = make([]int, 0)
	for i, row := range updRules {
		flag := true
		for j, el := range row {
			rest := row[j+1:]
			if len(rest) == 0 { break }
			if _, res := isOrderRulesIncludes(el, rest); !res {
				flag = false
				break
			}
		}
		if flag {
			correct = append(correct, i)
		} else {
			incorrect = append(incorrect, i)
		}

	}
}

func fixIncorrectGetMid(row []int) int {
	midIdx := len(row) / 2
	restart:
	for i, el := range row {
		rest := row[i+1:]
		if len(rest) == 0 || i > midIdx {
			break
		}
		if badIdx, res := isOrderRulesIncludes(el, rest); !res {
			row[badIdx+i+1], row[i] = row[i], row[badIdx+i+1]
			goto restart
		}
	}
	return row[midIdx]
}

func part1() {
	total := 0
	for _, idx := range correct {
		l := len(updRules[idx])
		total += updRules[idx][l/2]
	}
	fmt.Println("Part 1 result:", total)
}

func part2() {
	total := 0
	for _, idx := range incorrect {
		total += fixIncorrectGetMid(updRules[idx])
	}
	fmt.Println("Part 2 result:", total)
}


func main() {
	data := getFileData()
	parsePageOrderingRules(data)
	parseUpdateRules(data)
	idxSplitByCorrectUpdateRows()

	part1()
	part2()
}

func getFileData() string {
	text, err := os.ReadFile("day05/input.txt")
	if err != nil {
		panic("Can't load input file!")
	}
	return string(text)
}

func getData() string {
	return `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47
`
}

/*
Part 1 result: 6949
Part 2 result: 4145
*/