package main

import (
	"fmt"
	"os"
	"strings"

	"../utils"
)

func part2(input [][]rune) {
	//draw(input)
	antMap := mapAntennas(input)
	total := 0
	uniqAntinode := utils.SetOf[Antenna]()
	maxY, maxX := len(input), len(input[0])
	for _, arr := range antMap {
		for _, current := range arr {
			for _, other := range arr {
				if current == other { continue }
				x, y, dx, dy := current.antinode(other)
				if x < 0 || y < 0 || x >= maxX || y >= maxY { continue}
				a := Antenna{x: x, y: y, freq: 'a'}
				if !uniqAntinode.Contains(a) && input[y][x] == '.' {
					uniqAntinode.Add(a)
					total++
				}
				input[y][x] = '#'
				for {
					x += dx
					y += dy
					if x < 0 || y < 0 || x >= maxX || y >= maxY { break}
					a := Antenna{x: x, y: y, freq: 'a'}
					if !uniqAntinode.Contains(a) && input[y][x] == '.'  {
						uniqAntinode.Add(a)
						total++
					}
					if input[y][x] == '.' { input[y][x] = '#' }
				}
			}
		}
		if len(arr) > 1 {
			total += len(arr)
		}
	}
	fmt.Println("")
	//draw(input)

	fmt.Println("Part 2 response:", total)
}

func part1(input [][]rune) {
	//draw(input)
	antMap := mapAntennas(input)
	total := 0
	uniqAntinode := utils.SetOf[Antenna]()
	maxY, maxX := len(input), len(input[0])
	for _, arr := range antMap {
		for _, current := range arr {
			for _, other := range arr {
				if current == other { continue }
				x, y, _, _ := current.antinode(other)
				if x < 0 || y < 0 || x >= maxX || y >= maxY { continue}
				a := Antenna{x: x, y: y, freq: 'a'}
				if !uniqAntinode.Contains(a) {
					uniqAntinode.Add(a)
					total++
				}
				input[y][x] = '#'
			}
		}
	}
	fmt.Println("")
	//draw(input)
	fmt.Println("Part 1 response:", total)
}

func draw(input [][]rune) {
	for _, arr := range input {
		for _, r := range arr {
			fmt.Print(string(r))
		}
		fmt.Println()
	}
}

type Antenna struct {
	x, y int
	freq rune
}

func (my *Antenna) antinode(other Antenna) (x, y, dx, dy int) {
	newx, newy := 0, 0
	switch {
	case my.x > other.x:
		diff := my.x - other.x
		newx = my.x - 2*diff
	case my.x < other.x:
		newx = other.x + other.x - my.x
	}
	switch {
	case my.y > other.y:
		diff := my.y - other.y
		newy = my.y - 2*diff
	case my.y < other.y:
		newy = other.y + other.y - my.y
 	}
	return newx, newy, other.x - my.x, other.y - my.y
}

func mapAntennas(input [][]rune) map[rune][]Antenna {
	antennas := make(map[rune][]Antenna)

	for y, ar := range input {
		for x, elm := range ar {
			if elm == rune('.') {  continue }
			a := Antenna{x: x, y: y, freq: elm}
			antennas[elm] = append(antennas[elm], a)
		}
	}
	return antennas
}

func main() {
	test := false
	data := loadData(test)
	part1(parseData(data))
	part2(parseData(data))
}

func parseData(s string) [][]rune {
	str_arr := strings.Split(strings.TrimSpace(s), "\n")
	result := make([][]rune, len(str_arr))
	for i, l := range str_arr {
		result[i] = make([]rune, len(l))
		for j, k := range l {
			result[i][j] = k
		}
	}
	return result
}

func loadData(wantTest bool) string {
	if wantTest {
		return `
............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............
`
	} else {
		text, err := os.ReadFile("day08/input.txt")
		if err != nil { panic(err) }
		return string(text)
	}
}

/*
Part 1 response: 14 test
Part 1 response: 379 input
Part 1 response: 34 test
Part 1 response: 1339 input
*/