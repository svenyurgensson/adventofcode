package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ldiv(x, y int) (int, int) {
	return x % y, x / y
}

func solution1(g Game) int {
	da_rem, da_quot := ldiv(g.prizeX*g.dyB - g.prizeY*g.dxB, g.dxA*g.dyB - g.dyA*g.dxB)
	db_rem, db_quot := ldiv(g.prizeY*g.dxA - g.prizeX*g.dyA, g.dxA*g.dyB - g.dyA*g.dxB)
	if da_rem == 0 && db_rem == 0 {
		return 3 * da_quot + db_quot
	}
	return 0
}

func part(inputs []Game, numb int) {
	total := 0

	for _, el := range inputs {
		total += solution1(el)
	}

	fmt.Printf("Part %d result: %d\n", numb, total)
}

func main() {
	isTest := false
	data := loadData(isTest)
	part(parseData(data, 0), 1)
	part(parseData(data, 10000000000000), 2)
}

type Game struct {
	dxA, dxB int
	dyA, dyB int
	prizeX, prizeY int
}

func parseData(data string, correct int) []Game {
	str_arr := strings.Split(strings.TrimSpace(data), "\n\n")
	games := make([]Game, 0, len(str_arr))
	for _, block := range str_arr {
		re := regexp.MustCompile(`(?ms)X\+(\d+),\sY\+(\d+).+X\+(\d+),\sY\+(\d+).+X=(\d+), Y=(\d+)`)
		split := re.FindAllStringSubmatch(block, -1)
		a := [6]int{}
		for i := range 6 {
			k, err := strconv.Atoi(split[0][i+1])
			if err != nil { panic(err) }
			a[i] = k
		}
		g := Game{dxA: a[0], dyA: a[1], dxB: a[2], dyB: a[3], prizeX: a[4] + correct, prizeY: a[5] + correct}
		games = append(games, g)
	}
	return games
}

func loadData(wantTest bool) string {
	if wantTest {
		return `
Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279
`
	} else {
		text, err := os.ReadFile("day13/input.txt")
		if err != nil { panic(err) }
		return string(text)
	}
}
/*
Part 1 result: 36954 input*
Part 2 result: 79352015273424 input*
*/