package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func part2() {
	data, err := os.ReadFile("day03/input.txt")
	if err != nil {
		panic("Can't load input file!")
	}

	re, _ := regexp.Compile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	res := re.FindAllStringSubmatch(string(data), -1)
	fmt.Println("len =", len(res))
	total := 0
	calcFlag := true
	for _, el := range res {
		//fmt.Println(el)
		switch {
		case el[0] == "do()":
			calcFlag = true
			continue
		case el[0] == "don't()":
			calcFlag = false
			continue
		}
		if !calcFlag {
			continue
		}
		r1, err := strconv.Atoi(el[1])
		if err != nil {
			panic("Wrong integer value!")
		}
		r2, err := strconv.Atoi(el[2])
		if err != nil {
			panic("Wrong integer value!")
		}
		total += r1 * r2
	}
	fmt.Println("Part 2 response:", total)
}

func part1() {
	data, err := os.ReadFile("day03/input.txt")
	if err != nil {
		panic("Can't load input file!")
	}

	re, _ := regexp.Compile(`mul\((\d+),(\d+)\)`)
	res := re.FindAllStringSubmatch(string(data), -1)
	fmt.Println("len =", len(res))
	total := 0
	for _, el := range res {
		r1, err := strconv.Atoi(el[1])
		if err != nil {
			panic("Wrong integer value!")
		}
		r2, err := strconv.Atoi(el[2])
		if err != nil {
			panic("Wrong integer value!")
		}
		total += r1 * r2
	}
	fmt.Println("Part 1 response:", total)
}

func main() {
	part1()
	part2()
}

/*
len = 747
Part 1 response: 181345830
len = 814
Part 2 response: 98729041
*/