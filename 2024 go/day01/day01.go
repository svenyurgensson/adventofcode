package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var lhs = []int{}
var rhs = []int{}

func prepare_input() {
	data, err := os.ReadFile("day01/input.txt")
	if err != nil {
		panic("Can't load input file!")
	}

	for _, line := range strings.Split(string(data), "\n") {
		results := strings.Fields(line)
		if len(results) < 1 {
			break
		}
		res, err := strconv.Atoi(results[0])
		if err != nil {
			panic("Wrong integer value!")
		}
		lhs = append(lhs, res)
		res, err = strconv.Atoi(results[1])
		if err != nil {
			panic("Wrong integer value!")
		}
		rhs = append(rhs, res)
	}
	sort.Ints(lhs)
	sort.Ints(rhs)
}

func part1() {
	diff := 0
	for i, el := range(lhs) {
		if el >= rhs[i] {
			diff += el - rhs[i]
		} else {
			diff += rhs[i] - el
		}
	}

	fmt.Printf("Part 1 response: %d\n", diff)
}

func similarity_score(n int) int {
	score := 0
	for _, el := range(rhs) {
		if el < n {
			continue
		}
		if el == n {
			score += 1
		}
		if el > n {
			break
		}
	}
	return n * score
}

func part2() {
	total := 0
	for _, el := range(lhs) {
		total += similarity_score(el)
	}
	fmt.Printf("Part 2 response: %d\n", total)
}


func main() {
	prepare_input()
	part1()
	part2()
}

/*
Part 1 response: 1189304
Part 2 response: 24349736
*/