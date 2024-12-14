package main

import (
	"fmt"
	"os"
	"strings"
)

var maxX, maxY int

func inGrid(x, y int) bool {
	return x >= 0 && x <= maxX &&
		   y >= 0 && y <= maxY
}

func calcCorners(x, y, el int, input[][]int) int {
	corners := 0
	dUL, dUR, dLL := [2]int{-1, -1}, [2]int{-1,  0}, [2]int{ 0, -1}

	for range 4 {
		UL := [2]int{x + dUL[0], y + dUL[1]}
		UR := [2]int{x + dUR[0], y + dUR[1]}
		LL := [2]int{x + dLL[0], y + dLL[1]}
		ul := inGrid((UL[0]), UL[1]) && (input[UL[1]][UL[0]] == el || input[UL[1]][UL[0]] == -el)
		ur := inGrid((UR[0]), UR[1]) && (input[UR[1]][UR[0]] == el || input[UR[1]][UR[0]] == -el)
		ll := inGrid((LL[0]), LL[1]) && (input[LL[1]][LL[0]] == el || input[LL[1]][LL[0]] == -el)
		if !ur && !ll { corners++ }
		if ur && ll && !ul { corners++ }
		dUL[0], dUL[1] = dUL[1], -dUL[0] // rotate
		dUR[0], dUR[1] = dUR[1], -dUR[0]
		dLL[0], dLL[1] = dLL[1], -dLL[0]
	}

	return corners
}

func markCalcArea(x, y, el int, input[][]int, area *int, perimetr *int, corners *int) {
	if x < 0 || y < 0 || x > maxX || y > maxY || input[y][x] != el {
		return
	}

	input[y][x] = -el // mark as visited
	(*corners) += calcCorners(x, y, el, input) // calc corners
	(*area)++ // update area size

	if y == 0 || (input[y-1][x] != -el && input[y-1][x] != el) { // up
		(*perimetr)++
	}
	if y == maxY || (input[y+1][x] != -el && input[y+1][x] != el) { // down
		(*perimetr)++
	}
	if x == 0 || ( input[y][x-1] != -el &&  input[y][x-1] != el) { // left
		(*perimetr)++
	}
	if x == maxX || (input[y][x+1] != -el && input[y][x+1] != el) { // right
		(*perimetr)++
	}

	markCalcArea(x + 1, y, el, input, area, perimetr, corners)
	markCalcArea(x - 1, y, el, input, area, perimetr, corners)
	markCalcArea(x, y + 1, el, input, area, perimetr, corners)
	markCalcArea(x, y - 1, el, input, area, perimetr, corners)
}

func part2(input [][]int) {
	price := 0
	for y, row := range input {
		for x, el := range row {
			if el < 0 { continue }
			area, perimeter, corners := 0, 0, 0
			markCalcArea(x, y, el, input, &area, &perimeter, &corners)
			price += corners * area
		}
	}

	fmt.Println("Part 2 response:", price)
}

func part1(input [][]int) {
	price := 0
	for y, row := range input {
		for x, el := range row {
			if el < 0 { continue }
			area, perimeter, corners := 0, 0, 0
			markCalcArea(x, y, el, input, &area, &perimeter, &corners)
			price += perimeter * area
		}
	}

	fmt.Println("Part 1 response:", price)
}

func main() {
	isTest := false
	data := loadData(isTest)
	part1(parseData(data))
	part2(parseData(data))
}

func parseData(data string) [][]int {
	str_arr := strings.Split(strings.TrimSpace(data), "\n")
	plants := make([][]int, len(str_arr))
	for y, str_elm := range str_arr {
		s_arr := strings.Trim(str_elm, "\n")
		plants[y] = make([]int, len(s_arr))
		for x, el := range s_arr {
			plants[y][x] = int(el)
		}
	}

	maxX = len(plants[0]) - 1
	maxY = len(plants) - 1
	return plants
}

func loadData(wantTest bool) string {
	if wantTest {
		return `
RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE
`
	} else {
		text, err := os.ReadFile("day12/input.txt")
		if err != nil { panic(err) }
		return string(text)
	}
}

/*
Part 1 response: 1446042 input
Part 2 response:  input
*/