package main

import (
	"fmt"
	"strconv"
	"strings"
)

func splitNumberByHalfs(number uint64, digits int) (uint64, uint64) {
	half := digits / 2
	scale := uint64(1)
	for range half {
		scale *= 10
	}
	lhs := number / scale
	rhs := number % scale
	return lhs, rhs
}

func digiCount(i uint64) int {
	d := 1
	v := uint64(10)
	for i >= v {
		v *= 10
		d++
	}
	return d
}

var memo map[uint64]uint64

func calcBlink(number uint64, current int) uint64 {
	if current == 0 {
		//fmt.Println(number)
		return 1
	}
	key := (number << 8) | uint64(current & 0xff)
	if res, ok := memo[key]; ok {
		return res
	}

	digits := digiCount(number)
	isEven := digits & 0x01 == 0
	var result uint64

	switch {
	case number == 0:
		result = calcBlink(1, current - 1)
	case isEven:
		first, second := splitNumberByHalfs(number, digits)
		result = calcBlink(first, current - 1) + calcBlink(second, current - 1)
	default:
		result = calcBlink(number * 2024, current - 1)
	}

	memo[key] = result
	return result
}

func solve(inputs []uint64, nblink int, part int) {
	memo = make(map[uint64]uint64)
	total := uint64(0)
	for i := range inputs {
		total += calcBlink(inputs[i], nblink)
	}

	fmt.Printf("Part %d response: %d\n", part, total)
}

func main() {
	isTest := false
	data := loadData(isTest)
	solve(parseData(data), 25, 1)
	solve(parseData(data), 75, 2)
}

func parseData(data string) []uint64 {
	str := strings.Split(strings.TrimSpace(data), " ")
	numbers := []uint64{}
	for _, el := range str {
		r, err := strconv.Atoi(string(el))
		if err != nil {	panic(err) }
		numbers = append(numbers, uint64(r))
	}

	return numbers
}

func loadData(wantTest bool) string {
	if wantTest {
		return `
125 17
`
	} else {
		return `
0 37551 469 63 1 791606 2065 9983586
`
	}
}

/*
Part 1 response: 55312 test
Part 1 response: 204022 input
Part 2 response: 241651071960597 input
*/