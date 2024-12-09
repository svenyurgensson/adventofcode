package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"../utils"
)

func draw(input []int) {
	for _, el := range input {
		if el == -1 {
			fmt.Print(".")
		} else {
			fmt.Print(el)
		}
	}
	fmt.Println()
}

// idx index of -1 block with length or -1 if not found
func detectSuitableGap(input []int, length int) int {
	lastIdx := len(input) - 1
	for i, el := range input {
		if el != -1 { continue }
		ii := i
		for range length {
			el = input[ii]
			if el != -1 {
				goto bad
			}
			if ii >= lastIdx {
				return -1
			}
			ii++
		} // here we passed whole length
		return i
		bad:
	}

	return -1
}

// (idx, length) not -1 block from end
func detectLatestBlock(input []int, maxIdx int) (int, int) {
	var idx int
	length := 0
	el := -1

	for idx = maxIdx; idx >= 0; idx-- {
		if input[idx] == el { continue }
		break
	}

	if idx == 0 {
		return 0, 0
	}

	length = 1
	el = input[idx]
	for i := idx - 1; i >= 0; i-- {
		if input[i] != el {
			break
		}
		idx--
		length++
	}

	return idx, length
}

// [0 0 -1 -1 -1 1 1 1 -1 -1 -1 2 -1 -1 -1 3 3 3 -1 4 4 -1 5 5 5 5 -1 6 6 6 6 -1 7 7 7 -1 8 8 8 8 9 9]
func defragment2(input []int) {
	maxIdx := len(input) - 1

	for maxIdx > 0 {
		if input[maxIdx] < 0 {
			maxIdx--
			continue
		}
		idxBlock, length := detectLatestBlock(input, maxIdx)
		idxGap := detectSuitableGap(input, length)
		if idxGap >= 0 && idxGap < (idxBlock - length) {
			for i := range length {
				input[idxGap + i] = input[idxBlock + i]
				input[idxBlock + i] = -1
			}
		}
		maxIdx -= length
	}
}

func part2(input []int) {
	//draw(input)
	defragment2(input)

	total := 0
	for i, el := range input {
		if el == -1 { continue }
		total += i * el
	}
	//draw(input)
	fmt.Println("Part 2 response:", total)
}

// [0 0 -1 -1 -1 1 1 1 -1 -1 -1 2 -1 -1 -1 3 3 3 -1 4 4 -1 5 5 5 5 -1 6 6 6 6 -1 7 7 7 -1 8 8 8 8 9 9]
// flag true: array was changed false: all data defragmented
func defragment1(input []int) bool {
	flag := false
	firstFreeIdx := slices.Index(input, -1)
	rest := utils.Reverse(input)
	lastBlockIdx := len(input) - slices.IndexFunc(rest, func(e int) bool { return e != -1 })
	if lastBlockIdx <= firstFreeIdx { return flag }
	for i, el := range rest {
		if el == -1 { continue }
		elIdx := len(input) - i - 1
		input[firstFreeIdx] = el
		input[elIdx] = -1
		flag = true
		break
	}
	return flag
}

func part1(input []int) {
	for {
		fragm := defragment1(input)
		if !fragm { break }
	}
	total := 0
	for i, el := range input {
		if el == -1 { break }
		total += i * el
	}
	fmt.Println("Part 1 response:", total)
}

func main() {
	isTest := false
	data := loadData(isTest)
	part1(parseData(data))
	part2(parseData(data))
}

func parseData(data string) []int {
	str_arr := strings.Split(strings.TrimSpace(data), "")
	result := make([]int, 0, len(str_arr))
	for i, el := range str_arr {
		r, err := strconv.Atoi(el)
		if err != nil {	panic(err) }
		if i % 2 == 0 {
			for range r {
				result = append(result, i / 2)
			}
		} else {
			if r == 0 { continue }
			for range r {
				result = append(result, -1)
			}
		}
	}
	return result
}

func loadData(wantTest bool) string {
	if wantTest {
		return "2333133121414131402"
	} else {
		text, err := os.ReadFile("day09/input.txt")
		if err != nil { panic(err) }
		return string(text)
	}
}

/*
Part 1 response: 1928 test
Part 1 response: 6291146824486 input
Part 1 response: 2858 test
Part 1 response: 6307279963620 input*
*/