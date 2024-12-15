package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// uncomment for test run
//const maxX int = 11
//const maxY int = 7
//const midX int = 5
//const midY int = 3
const maxX int = 101
const maxY int = 103
const midX int = 50
const midY int = 51

// 0 1
// 2 3
func robotsCount(quadr int, robots []Robot) int {
	x, y := 0, 0
	switch quadr {
	case 0:
		x, y = 0, 0
	case 1:
		x, y = midX+1, 0
	case 2:
		x, y = 0,midY+1
	default:
		x, y = midX+1,midY+1
	}
	count := 0
	for sx := x; sx < x+midX; sx++ {
		for sy := y; sy < y+midY; sy++ {
			for _, r := range robots {
				if r.x == sx && r.y == sy { count++ }
			}
		}
	}
	return count
}

func draw(robots []Robot) {
	t := [maxY][maxX]string{}
	for y, _ := range t {
		for x, _ := range t[0] {
			t[y][x] = "."
		}
	}
	for _, r := range robots {
		rs := t[r.y][r.x]
		switch rs {
		case ".":
			t[r.y][r.x] = "1"
		case "1":
			t[r.y][r.x] = "2"
		case "2":
			t[r.y][r.x] = "3"
		default:
			t[r.y][r.x] = "X"
		}
	}
	fmt.Println("")
	for y := range len(t) {
		for x := range len(t[0]) {
			fmt.Print(t[y][x])
		}
		fmt.Print("\n")
	}
}

func part1(input []Robot) {
	//draw(input)
	for range 100 {
		for i := range input {
			input[i].step()
		}

	}
	//draw(input)
	total := robotsCount(0, input) *
			 robotsCount(1, input) *
			 robotsCount(2, input) *
			 robotsCount(3, input)

	fmt.Println("Part 1 result:", total)
}

func detectTree(input []Robot) bool {
	for i := range maxY {
		xx := []int{}
		for _, el := range input {
			if el.y != i { continue }
			xx = append(xx, el.x)
		}
		if len(xx) > 30 {
			// detect asc numbers in xx
			sort.Ints(xx)
			success := 0
			for k := range len(xx) - 1 {
				if xx[k+1] - xx[k] == 1 {
					success++
				}
			}
			if (float32(success) / float32(len(xx))) * 100 > 90 {
				return true
			}
		}
	}

	return false
}

func part2(input []Robot) {
	//draw(input)
	for j := range 10000 {
		for i := range input {
			input[i].step()
		}
		if detectTree(input) {
			fmt.Println("Part 2 result:", j+1)
			draw(input)
			return
		}
	}
	//draw(input)
	fmt.Println("Part 2 result: tree not found")
}

func main() {
	isTest := false
	data := loadData(isTest)
	parsed := parseData(data)
	part1(parsed)
	parsed = parseData(data)
	part2(parsed)
}

type Robot struct {
	x, y int
	vx, vy int
}

func (r *Robot) step() {
	switch {
	case (r.x + r.vx) >= maxX:
		r.x = r.x + r.vx - maxX
	case (r.x + r.vx) < 0:
		r.x = maxX + (r.x + r.vx)
	default:
		r.x += r.vx
	}
	switch {
	case (r.y + r.vy) >= maxY:
		r.y = r.y + r.vy - maxY
	case  (r.y + r.vy) < 0:
		r.y = maxY + (r.y + r.vy)
	default:
		r.y += r.vy
	}
}

func parseData(data string) []Robot {
	str_arr := strings.Split(strings.TrimSpace(data), "\n")
	robots := make([]Robot, 0, len(str_arr))
	for _, block := range str_arr {
		re := regexp.MustCompile(`[^-\d]+(-?\d+)[^-\d]+(-?\d+)[^-\d]+(-?\d+)[^-\d]+(-?\d+)`)
		split := re.FindAllStringSubmatch(block, -1)
		a := [4]int{}
		for i := range 4 {
			k, err := strconv.Atoi(split[0][i+1])
			if err != nil { panic(err) }
			a[i] = k
		}
		robots = append(robots, Robot{x: a[0], y: a[1], vx: a[2], vy: a[3]})
	}
	return robots
}

func loadData(wantTest bool) string {
	if wantTest {
		return `
p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3
`
	} else {
		text, err := os.ReadFile("day14/input.txt")
		if err != nil { panic(err) }
		return string(text)
	}
}

/*
Part 1 result: 221655456
Part 2 result: 7857
*/