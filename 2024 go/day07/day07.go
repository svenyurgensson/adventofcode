package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"../utils"
)

func validate2(expected int, array []int) bool {
	if expected < 0 {
		return false
	}
	if len(array) == 0 {
		return expected == 0
	}
	current := array[0]
	if expected % current == 0 && validate2(expected / current, array[1:]) {
		return true
	}
	if next, found := strings.CutSuffix(strconv.Itoa(expected), strconv.Itoa(current)); found && len(next) != 0 {
   		n, err := strconv.Atoi(next)
   		if err != nil { panic(err) }
		if validate2(n, array[1:]) {
			return true
		}
	}
	return validate2(expected - current, array[1:])
}

func part2(input [][]int) {
	result := 0
	for _, arr := range input {
		if validate2(arr[0], utils.Reverse(arr[1:])) {
			result += arr[0]
		}
	}

	fmt.Println("Part 2 response:", result)
}

func validate(expected int, array []int) bool {
	if len(array) == 1 {
		return expected == array[0]
	} else if len(array) == 0 || expected <= 0{
		return false
	} else {
		return (expected % array[0] == 0 && validate(expected / array[0], array[1:])) ||
				validate(expected - array[0], array[1:])
	}
}

func part1(input [][]int) {
	result := 0
	for _, arr := range input {
		if validate(arr[0], utils.Reverse(arr[1:])) {
			result += arr[0]
		}
	}
	fmt.Println("Part 1 response:", result)
}

func main() {
	test := false
	data := loadData(test)
	part1(parseData(data))
	part2(parseData(data))
}

func parseData(s string) [][]int {
	str_arr := strings.Split(strings.TrimSpace(s), "\n")
	result := make([][]int, len(str_arr))
	re, _ := regexp.Compile(`(\d+)\D?`)
	for i, l := range str_arr {
		t := re.FindAllStringSubmatch(l, -1)
		result[i] = make([]int, len(t))
		for j, k := range t {
			r, err := strconv.Atoi(k[1])
			if err != nil {	panic("Wrong integer value!") }
			result[i][j] = r
		}
	}
	return result
}

func loadData(wantTest bool) string {
	if wantTest {
		return `
190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
`
	} else {
		text, err := os.ReadFile("day07/input.txt")
		if err != nil {
			panic(err)
		}
		return string(text)
	}
}

/*
Part 1 response: 3749 test
Part 1 response: 5512534574980  input.txt
Part 2 response: 11387 test
Part 2 response: 328790210468594  input.txt
*/