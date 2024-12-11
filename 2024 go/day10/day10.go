package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"../utils"
)

type Coord struct {
	x, y int
}

var endPositions utils.Set[Coord]

func tryMove2(inputs [][]int, x, y int) int {
	height := inputs[y][x]
	if height == 9 {
		return 1
	}
	score := 0
	// left
    if x > 0 && inputs[y][x - 1] == height + 1 {
        score += tryMove2(inputs, x - 1, y)
	}
    // up
    if y > 0 && inputs[y - 1][x] == height + 1 {
        score += tryMove2(inputs, x, y - 1)
	}
    // right
    if x < len(inputs[y]) - 1 && inputs[y][x + 1] == height + 1 {
        score += tryMove2(inputs, x + 1, y)
	}
    // down
    if y < len(inputs) - 1 && inputs[y + 1][x] == height + 1 {
        score += tryMove2(inputs, x, y + 1)
	}

    return score
}

func part2(inputs [][]int) {
	total := 0
	for y := range len(inputs) {
		for x := range len(inputs[y]) {
			if inputs[y][x] == 0 {
				total += tryMove2(inputs, x, y)
			}
		}
	}
	fmt.Println("Part 2 response:", total)
}

func tryMove(inputs [][]int, x, y int) {
	height := inputs[y][x]
	if height == 9 {
		endPositions.Add(Coord{x: x, y: y})
		return
	}
	// left
	if x > 0 && inputs[y][x - 1] == height + 1 {
		tryMove(inputs, x - 1, y)
	}
	// up
	if y > 0 && inputs[y - 1][x] == height + 1 {
		tryMove(inputs, x, y - 1)
	}
	// right
	if x < (len(inputs[y]) - 1) && inputs[y][x + 1] == height + 1 {
		tryMove(inputs, x + 1, y)
	}
	// dowm
	if y < (len(inputs) - 1) && inputs[y + 1][x] == height + 1 {
		tryMove(inputs, x, y + 1)
	}
}

func part1(inputs [][]int) {
	total := 0
	for y := range len(inputs) {
		for x := range len(inputs[y]) {
			if inputs[y][x] == 0 {
				endPositions = utils.SetOf[Coord]()
				tryMove(inputs, x, y)
				total += len(endPositions)
			}
		}
	}
	fmt.Println("Part 1 response:", total)
}

func main() {
	isTest := false
	data := loadData(isTest)
	inputs := parseData(data)
	part1(inputs)
	part2(inputs)
}

func parseData(data string) [][]int {
	str_arr := strings.Split(strings.TrimSpace(data), "\n")
	topoMap := make([][]int, len(str_arr))
	for y, str_elm := range str_arr {
		s_arr := strings.Trim(str_elm, "\n")
		topoMap[y] = make([]int, len(s_arr))
		for x, el := range s_arr {
			if el == '.' { topoMap[y][x] = -1; continue }
			r, err := strconv.Atoi(string(el))
			if err != nil {	panic(err) }
			topoMap[y][x] = r
		}
	}

	return topoMap
}

func loadData(wantTest bool) string {
	if wantTest {
		return `
89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732
`
	} else {
		text, err := os.ReadFile("day10/input.txt")
		if err != nil { panic(err) }
		return string(text)
	}
}
/*
Part 1 response: 36 test
Part 1 response: 582 input
Part 2 response: 81 test
Part 2 response: 1302 input
*/