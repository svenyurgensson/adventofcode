package main

import (
	"fmt"
	"os"
	"slices"
)

func countNl(data string) int {
	count := 0
	for _, s := range data {
		if s == '\n' {
			count++
		}
	}
	return count
}

func part1() {
	text, err := os.ReadFile("day04/input.txt")
	if err != nil {
		panic("Can't load input file!")
	}
	data := string(text)
	/* data := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
` */
	width := slices.IndexFunc([]rune(data), func(c rune) bool { return c == '\n' }) + 1
	height := countNl(data)
	dataLen := len(data)
	fmt.Println("len:", dataLen, "width:", width, "height:", height)
	total_xmas := 0
	total_samx := 0

	// calc XMAS
	for i := range data {
		// horizontal
		if (i <= dataLen - 4) && (data[i:i+4] == "XMAS") {
				total_xmas++
		}
		// vertical
		if (i <= dataLen - 3 * width - 2) &&
			(data[i]           == 'X') &&
			(data[i + 1*width] == 'M') &&
			(data[i + 2*width] == 'A') &&
			(data[i + 3*width] == 'S') {
				total_xmas++
		}
		// vertical diagonal top to left
		if (i <= dataLen - 3 * width - 2) &&
			(data[i]               == 'X') &&
			(data[i + 1*width - 1] == 'M') &&
			(data[i + 2*width - 2] == 'A') &&
			(data[i + 3*width - 3] == 'S') {
				total_xmas++
		}
		// vertical diagonal top to right
		if (i <= dataLen - 3 * width - 5) &&
			(data[i]               == 'X') &&
			(data[i + 1*width + 1] == 'M') &&
			(data[i + 2*width + 2] == 'A') &&
			(data[i + 3*width + 3] == 'S') {
				total_xmas++
		}

		// calc SAMX
		// horizontal
		if (i <= dataLen - 4) && (data[i:i+4] == "SAMX") {
				total_samx++
		}
		// vertical
		if (i <= dataLen - 3 * width - 2) &&
			(data[i]           == 'S') &&
			(data[i + 1*width] == 'A') &&
			(data[i + 2*width] == 'M') &&
			(data[i + 3*width] == 'X') {
				total_samx++
		}
		// vertical diagonal to left
		if (i <= dataLen - 3 * width - 2) &&
			(data[i]               == 'S') &&
			(data[i + 1*width - 1] == 'A') &&
			(data[i + 2*width - 2] == 'M') &&
			(data[i + 3*width - 3] == 'X') {
				total_samx++
		}
		// vertical diagonal to right
		if (i <= dataLen - 3 * width - 5) &&
			(data[i]               == 'S') &&
			(data[i + 1*width + 1] == 'A') &&
			(data[i + 2*width + 2] == 'M') &&
			(data[i + 3*width + 3] == 'X') {
				total_samx++
		}
	}
	total := total_xmas + total_samx
	fmt.Println("Part 1 result:", total, "/ xmas:", total_xmas, " / samx:",  total_samx)
}

func part2() {
	/**/
	text, err := os.ReadFile("day04/input.txt")
	if err != nil {
		panic("Can't load input file!")
	}
	data := string(text)
	/**/
	/*
	data := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
`
	/**/
	width := slices.IndexFunc([]rune(data), func(c rune) bool { return c == '\n' }) + 1
	height := countNl(data)
	total := 0

	for i := range height - 1 {
		for j := range width - 3 {
			s := i + j * width
			// tl . tr
			//  . c .
			// bl . br
			tl, tr, c, bl, br := data[s], data[s + 2], data[s + width + 1], data[s + 2 * width], data[s + 2 * width + 2]
			if (c == 'A') && (
				(tl == 'M' && tr == 'M' && bl == 'S' && br == 'S') ||
				(tl == 'M' && tr == 'S' && bl == 'M' && br == 'S') ||
				(tl == 'S' && tr == 'M' && bl == 'S' && br == 'M') ||
				(tl == 'S' && tr == 'S' && bl == 'M' && br == 'M')) {
					total++
			}
		}
	}
	fmt.Println("Part 2 result:", total)
}

func main() {
	part1()
	part2()
}

// Part 1 result: 2685
// Part 2 result: 2048