package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() {
	data, err := os.ReadFile("day02/input.txt")
	if err != nil {
		panic("Can't load input file!")
	}
/*
	data := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`
*/
	safe_cnt := 0
	for _, line := range strings.Split(string(data), "\n") {
		results := strings.Fields(line)
		if len(results) < 1 {
			break
		}
		ary := make([] int, len(results))
		for i, el := range(results) {
			r, err := strconv.Atoi(el)
			if err != nil {
				panic("Wrong integer value!")
			}
			ary[i] = r
		}
		a, b := ary[0], ary[1]
		if a == b {
			continue
		}
		if a < b { // up
			for i, el := range(ary) {
				if i != 0 && ((el - a) < 1 || (el - a) > 3) {
					goto out
				}
				a = el
			}
			safe_cnt += 1
		} else { // down
			for i, el := range(ary) {
				if i != 0 && ((a - el) < 1 || (a - el) > 3) {
					goto out
				}
				a = el
			}
			safe_cnt += 1
		}
		out:
	}
	fmt.Printf("Part 1 response: %d\n", safe_cnt)
}

func removeIndex(s []int, index int) []int {
    ret := make([]int, 0)
    ret = append(ret, s[:index]...)
    return append(ret, s[index+1:]...)
}

func checkValidWithSkip(ary []int, idxToSkip int) bool {
	if idxToSkip >= 0 {
		ary = removeIndex(ary, idxToSkip)
	}

	a, b := ary[0], ary[1]
	if a == b {
		return false
	}

	if a < b { // up
		for i, el := range(ary) {
			if i != 0 && ((el - a) < 1 || (el - a) > 3) {
				return false
			}
			a = el
		}
	} else { // down
		for i, el := range(ary) {
			if i != 0 && ((a - el) < 1 || (a - el) > 3) {
				return false
			}
			a = el
		}
	}

	return true
}

func checkValid(ary []int) bool {
	if checkValidWithSkip(ary, -1) {
		return true
	}
	for i := range(len(ary)) {
		if checkValidWithSkip(ary, i) {
			return true
		}
	}
	return false
}

func part2() {
	data, err := os.ReadFile("day02/input.txt")
	if err != nil {
		panic("Can't load input file!")
	}
/*
	data := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`
*/
	safe_cnt := 0
	for _, line := range strings.Split(string(data), "\n") {
		results := strings.Fields(line)
		if len(results) < 1 {
			break
		}
		ary := make([] int, len(results))
		for i, el := range(results) {
			r, err := strconv.Atoi(el)
			if err != nil {
				panic("Wrong integer value!")
			}
			ary[i] = r
		}

		if checkValid(ary) {
			safe_cnt += 1
		}
	}
	fmt.Printf("Part 2 response: %d\n", safe_cnt)
}

func main() {
	part1()
	part2()
}

// Part 1 response: 306
// Part 2 response: 366