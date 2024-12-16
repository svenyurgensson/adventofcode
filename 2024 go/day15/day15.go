package main

import (
	"fmt"
	"os"
	"strings"
)

func hasRoom(g *Game, dir int) (int, int, bool) {
	dx, dy := 0, 0
	switch dir {
	case '<':
		dx = -1
	case '>':
		dx = 1
	case '^':
		dy = -1
	case 'v':
		dy = 1
	}
	x, y := g.rx, g.ry
	for {
		x, y = x+dx, y+dy
		if g.room[y][x] == '#' { return 0, 0, false }
		if g.room[y][x] == '.' { return x, y, true }
	}
}
// return bool finished?
func step(g *Game) bool {
	dir := g.path[g.step]
	g.room[g.ry][g.rx] = '.' // replace current robot pos with empty space
	switch dir {
	case '<':
		if emX, _, ok := hasRoom(g, '<'); ok { // can move!
			if emX == g.rx - 1 {
				g.rx-- // step robot
			} else {
				g.room[g.ry][emX] = g.room[g.ry][g.rx-1] // move box
				g.rx-- // step robot
			}
		}
	case '>':
		if emX, _, ok := hasRoom(g, '>'); ok { // can move!
			if emX == g.rx + 1 {
				g.rx++ // step robot
			} else {
				g.room[g.ry][emX] = g.room[g.ry][g.rx+1] // move box
				g.rx++ // step robot
			}
		}
	case '^':
		if _, emY, ok := hasRoom(g, '^'); ok { // can move!
			if emY == g.ry - 1 {
				g.ry-- // step robot
			} else {
				g.room[emY][g.rx] = g.room[g.ry-1][g.rx] // move box
				g.ry-- // step robot
			}
		}
	case 'v':
		if _, emY, ok := hasRoom(g, 'v'); ok { // can move!
			if emY == g.ry + 1 {
				g.ry++ // step robot
			} else {
				g.room[emY][g.rx] = g.room[g.ry+1][g.rx] // move box
				g.ry++ // step robot
			}
		}
	}
	g.room[g.ry][g.rx] = '@' // robot position

	g.step++
	return g.step >= len(g.path)
}

func draw(g Game) {
	for y := range len(g.room) {
		for x := range len(g.room[y]) {
			fmt.Printf("%c", g.room[y][x])
		}
		fmt.Println("")
	}
}

func part1(g Game) {
	//draw(g)
	for {
		if step(&g) { break }
	}
	sumGPS := 0
	for y, _ := range g.room {
		for x, el := range g.room[y] {
			if el == 'O' {
				sumGPS += 100 * y + x
			}
		}
	}
	//draw(g)
	fmt.Println("Part 1 result:", sumGPS)
}

func checkVertical(g *Game, px, py, vx, vy int) bool {
	lx, ly, rx, ry := 0,0,0,0
	switch g.room[py][px] {
	case '#':
		return false
	case '.':
		return true
	case '[':
		lx, ly = px, py + vy
		rx, ry = px + 1, py + vy
	case ']':
		lx, ly = px - 1, py + vy
		rx, ry = px, py + vy
	default:
		panic("check !")
	}
	return checkVertical(g, lx, ly, vx, vy) && checkVertical(g, rx, ry, vx, vy)
}

func shiftVertical(g *Game, px, py, vx, vy int) {
	lcx, lcy, rcx, rcy := 0,0,0,0 // current
	lnx, lny, rnx, rny := 0,0,0,0 // next
	switch g.room[py][px] {
	case '.':
		return
	case '[':
		lcx, lcy = px, py
		rcx, rcy = px + 1, py
		lnx, lny = px, py + vy
		rnx, rny = px + 1, py + vy
	case ']':
		lcx, lcy = px - 1, py
		rcx, rcy = px, py
		lnx, lny = px - 1, py + vy
		rnx, rny = px, py + vy
	default:
		panic("shift !")
	}
	shiftVertical(g, lnx, lny, vx, vy)
	shiftVertical(g, rnx, rny, vx, vy)
	g.room[lny][lnx] = g.room[lcy][lcx]
	g.room[rny][rnx] = g.room[rcy][rcx]
	g.room[lcy][lcx] = '.'
	g.room[rcy][rcx] = '.'
}

func step2(g *Game) bool {
	dir := g.path[g.step]
	switch dir {
	case '<':
		if emX, _, ok := hasRoom(g, '<'); ok { // can move!
			if emX != g.rx - 1 { // some further
				for x := emX; x < g.rx; x++ {
					g.room[g.ry][x] = g.room[g.ry][x+1] // move box
				}
			}
			g.room[g.ry][g.rx] = '.' // replace current robot pos with empty space
			g.rx-- // step robot
		}
	case '>':
		if emX, _, ok := hasRoom(g, '>'); ok { // can move!
			if emX != g.rx + 1 {
				if emX != g.rx - 1 { // some further
					for x := emX; x > g.rx; x-- {
						g.room[g.ry][x] = g.room[g.ry][x-1] // move box
					}
				}
			}
			g.room[g.ry][g.rx] = '.' // replace current robot pos with empty space
			g.rx++ // step robot
		}
	case '^':
		if  checkVertical(g, g.rx, g.ry-1, 0, -1) { // can move!
			shiftVertical(g, g.rx, g.ry-1, 0, -1) // shift up
			g.room[g.ry][g.rx] = '.' // replace current robot pos with empty space
			g.ry-- // step robot
		}
	case 'v':
		if  checkVertical(g, g.rx, g.ry+1, 0, 1) { // can move!
			shiftVertical(g, g.rx, g.ry+1, 0, 1) // shift down
			g.room[g.ry][g.rx] = '.' // replace current robot pos with empty space
			g.ry++ // step robot
		}
	}
	g.room[g.ry][g.rx] = '@' // robot position

	g.step++
	return g.step >= len(g.path)
}

func part2(g Game) {
	//draw(g)
	for {
		if step2(&g) { break }
	}
	//draw(g)
	sumGPS := 0
	for y, _ := range g.room {
		for x, el := range g.room[y] {
			if el == '[' {
				sumGPS += 100 * y + x
			}
		}
	}
	fmt.Println("Part 2 result:", sumGPS)
}

func main() {
	isTest := false
	data := loadData(isTest)
	part1(parseData(data))
	part2(parseData2(data))
}

type Game struct {
	room [][]int
	rx, ry int
	step int
	path []int
}

func parseData(data string) Game {
	str_arr := strings.Split(strings.TrimSpace(data), "\n\n")
	room := strings.Split(str_arr[0], "\n")
	game := Game{}
	game.room = make([][]int, len(room))
	for y, line := range room {
		game.room[y] = make([]int, len(string(line)))
		for x, el := range string(line) {
			if el == '\n' { continue }
			game.room[y][x] = int(el)
			if el == '@' {
				game.rx, game.ry = x, y
			}
		}
	}
	game.path = make([]int, 0, len(str_arr[1]))
	for _, el := range str_arr[1] {
		if el == '\n' { continue }
		game.path = append(game.path, int(el))
 	}
	return game
}

func parseData2(data string) Game {
	str_arr := strings.Split(strings.TrimSpace(data), "\n\n")
	room := strings.Split(str_arr[0], "\n")
	game := Game{}
	game.room = make([][]int, len(room))
	for y, line := range room {
		game.room[y] = make([]int, 2*len(string(line)))
		for x, el := range string(line) {
			xx := 2 * x
			switch el {
			case '\n':
				continue
			case '@':
				game.room[y][xx] = '@'
				game.room[y][xx+1] = '.'
				game.rx, game.ry = xx, y
			case 'O':
				game.room[y][xx] = '['
				game.room[y][xx+1] = ']'
			default:
				game.room[y][xx] = int(el)
				game.room[y][xx+1] = int(el)
			}

		}
	}
	game.path = make([]int, 0, len(str_arr[1]))
	for _, el := range str_arr[1] {
		if el == '\n' { continue }
		game.path = append(game.path, int(el))
 	}
	return game
}

func loadData(wantTest bool) string {
	if wantTest {
		return `
##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^
`
	} else {
		text, err := os.ReadFile("day15/input.txt")
		if err != nil { panic(err) }
		return string(text)
	}
}

/*
Part 1 result: 10092 test
Part 1 result: 1486930 input
Part 2 result: 9021 test
Part 2 result: 1492011 input
*/