package main

import (
	"fmt"
	"os"
	"slices"

	"../utils"
)

func drawRoom(room [][]byte, g Guard, o Coord) {
	for i := range len(room) {
		for j := range len(room[0]) {
			t := "."
			switch room[i][j] {
			case 0:
				t = "."
			case 0xff:
				t = "#"
			case byte(Up):
				t = "|"
			case byte(Down):
				t = "|"
			case byte(Left):
				t = "-"
			case byte(Right):
				t = "-"
			default:
				t = "+"
			}
			if g.x == j && g.y == i {
				switch g.dir {
				case Up:
					t = "^"
				case Down:
					t = "v"
				case Left:
					t = "<"
				case Right:
					t = ">"
				}
			}
			if o.x == j && o.y == i {
				t = "O"
			}
			fmt.Print(t)
		}
		fmt.Println()
	}
	fmt.Println("-----------------------")
}

func part2(room [][]byte, guard Guard) {
	visited := guard.track // save visited Coord{}
	loops := 0
	for coord := range visited {
		origRoom := room // save room [][]byte
		guard.reset()
		save := origRoom[coord.y][coord.x]
		origRoom[coord.y][coord.x] = 0xff
		if findLoop(origRoom, guard, visited) {
			//drawRoom(origRoom, guard, coord)
			loops++
		}
		origRoom[coord.y][coord.x] = save
	}
	fmt.Println("Part 2 result:", loops)
}

func part1(room [][]byte, guard Guard) Guard {
	for {
		res := guard.canStep(room)
		switch res {
		case -1:
			fmt.Println("Part 1 result:", guard.steps)
			return guard
		case 0:
			guard.turnRight()
		case 1:
			guard.step(room)
		}
	}
}

func findLoop(room [][]byte, guard Guard, visited utils.Set[Coord]) bool {
	for range 100_000 {
		res := guard.canStep(room)
		switch res {
		case -1: // reach exit
			return false
		case 0:
			guard.turnRight()
		case 1:
			//if visited.Contains(guard.getCoord()) {
			if visited.Contains(guard.getCoord()) {
				return true // cicled!
			}
			guard.step(room)
		}
	}
	return false
}


type Direction int

const (
	Zero  Direction = 0x00
	Up    Direction = 0x01
	Down  Direction = 0x02
	Left  Direction = 0x04
	Right Direction = 0x08
)

type Coord struct {
	x, y int
	dir Direction
}

type Guard struct {
	x, y, origX, origY int
	dir Direction
	steps int
	track utils.Set[Coord]
}

func newGuard(x, y int, dir Direction) Guard {
	return Guard{
		x: x,
		y: y,
		dir: dir,
		track: utils.SetOf[Coord](),
		origX: x,
		origY: y,
	}
}

func (g* Guard) reset() {
	g.x = g.origX
	g.y = g.origY
	g.steps = 0
	g.track = utils.SetOf[Coord]()
}

func (g* Guard) getCoord() Coord {
	return Coord{
		x: g.x,
		y: g.y,
		dir: g.dir,
	}
}

func (g* Guard) wasHere() bool {
	return g.track.Contains(g.getCoord())
}

// -1 => step out of room, 0 => can't step, 1 => can step
func (g* Guard) canStep(room [][]byte) int {
	height := len(room) - 1
	width := len(room[0]) - 1
	switch g.dir {
	case Up:
		if g.y == 0 { return -1 }
		if room[g.y - 1][g.x] != 0xff { return 1 }
	case Down:
		if g.y >= height { return -1 }
		if room[g.y + 1][g.x] != 0xff { return 1 }
	case Left:
		if g.x == 0 { return -1 }
		if room[g.y][g.x - 1] != 0xff { return 1 }
	case Right:
		if g.x >= width { return -1 }
		if room[g.y][g.x + 1] != 0xff { return 1 }
	}
	return 0
}

func (g* Guard) turnRight() {
	switch g.dir {
	case Up:
		g.dir = Right
	case Down:
		g.dir = Left
	case Left:
		g.dir = Up
	case Right:
		g.dir = Down
	}
}

func (g* Guard) step(room [][]byte) {
	switch g.dir {
	case Up:
		g.y--
	case Down:
		g.y++
	case Left:
		g.x--
	case Right:
		g.x++
	}
	if room[g.y][g.x] == 0 {
		g.steps++
	}
	room[g.y][g.x] |= byte(g.dir) // mark this coord with direction

	g.track.Add(Coord{x: g.x, y: g.y, dir: g.dir})
}

func main() {
	//isTest := false
	isTest := true
	data := loadData(isTest)
	roomMap, guard := parseMap(data)
	r := roomMap
	guard = part1(r, guard)

	part2(roomMap, guard)
}

func parseMap(data string) ([][]byte, Guard) {
	height := 0
	for _, s := range data {
		if s == '\n' {
			height++
		}
	}
	width := slices.IndexFunc([]rune(data), func(c rune) bool { return c == '\n' })
	result := make([][]byte, height)
	var guard Guard

	for i := range height {
		result[i] = make([]byte, width)
		for j := range(width + 1) { // take in account \n
			offs := i * (width+1) + j
			if data[offs] == '^' {
				guard = newGuard(j, i, Up)
			}
			if data[offs] == '#' {
				result[i][j] = 0xff
			}
		}
	}
	return result, guard
}

func loadData(wantTest bool) string {
	if wantTest {
		return `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
`
	} else {
		text, err := os.ReadFile("day06/input.txt")
		if err != nil {
			panic("Can't load input file!")
		}
		return string(text)
	}
}

/*
https://adventofcode.com/2024/day/6

Part 1 test result: 41
Part 1 result: 4602

Part 2 test result: 6
Part 2 result: 1703 ** not solved :()
*/