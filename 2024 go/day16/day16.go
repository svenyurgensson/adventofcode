package main

import (
	"fmt"
	"os"
	"strings"

	"../utils"
)

func draw(s Scene) {
	for y := range len(s.maze) {
		for _, el := range s.maze[y] {
			fmt.Print(el)
		}
		fmt.Println()
	}
}

func findAndSavePossibilities(s Scene, r Robot, visited *utils.Set[Coord]) {
	var (
		steps [][]Rotation
		ok bool
	)
	if steps, ok = possibleSteps(s, r, visited); !ok {
		return
	}
	for _, rots := range steps {
		newRobot := r.clone()
		for _, rot := range rots {
			if rot == RT { newRobot.right() }
			if rot == LF { newRobot.left() }
		}
		coord := newRobot.frontCoord()
		if visited.Contains(coord) {
			continue
		}
		robots = robots.Push(newRobot)
	}
}

// Calcs rotations list for current coord and direction
// returns:
//   [[LF], [RT, RT]] array of rotations to possible ways
//   true if here any way to move, false
func possibleSteps(s Scene, r Robot, visited *utils.Set[Coord]) ([][]Rotation, bool) {
	possibleCoords := findNeighbors(s, r, visited)
	rotations := [][]Rotation{}
	if len(possibleCoords) == 0 {
		return rotations, false
	}
	for _, c := range possibleCoords {
		rots := getRotations(r, c)
		if len(rots) > 0 {
			rotations = append(rotations, rots)
		}
	}
	return rotations, len(rotations) > 0
}

// return array of rotations to reach c[Coord] where step possible
func getRotations(r Robot, c Coord) []Rotation {
	rotations := []Rotation{}
	for range 3 {
		rotations = append(rotations, LF)
		if c == r.left() {
			if len(rotations) == 3 {
				return []Rotation{RT}
			}
			return rotations
		}
	}

	return rotations
}

func findNeighbors(s Scene, r Robot, visited *utils.Set[Coord]) []Coord {
	res := []Coord{}
	if s.maze[r.coord.y-1][r.coord.x] == "." && !visited.Contains(Coord{x: r.coord.x, y: r.coord.y-1}) { res = append(res, Coord{x: r.coord.x, y: r.coord.y-1}) }
	if s.maze[r.coord.y+1][r.coord.x] == "." && !visited.Contains(Coord{x: r.coord.x, y: r.coord.y+1}) { res = append(res, Coord{x: r.coord.x, y: r.coord.y+1}) }
	if s.maze[r.coord.y][r.coord.x-1] == "." && !visited.Contains(Coord{x: r.coord.x-1, y: r.coord.y}) { res = append(res, Coord{x: r.coord.x-1, y: r.coord.y}) }
	if s.maze[r.coord.y][r.coord.x+1] == "." && !visited.Contains(Coord{x: r.coord.x+1, y: r.coord.y}) { res = append(res, Coord{x: r.coord.x+1, y: r.coord.y}) }
	return res
}

func findPath(s Scene, r Robot) (Robot, bool) {
	visited := utils.SetOf[Coord]()
	for {
		if r.isFinished(s) {
			draw(s)
			return r, true
		}
		for {
			saveCoord := Coord{x: r.coord.x, y: r.coord.y}
			if visited.Contains(saveCoord) {
				if !r.canStepForward(s) {
					//draw(s)
					return r, false
				}
			}
			if  r.canStepForward(s) {
				visited.Add(saveCoord)
				findAndSavePossibilities(s, r, &visited)
				r.step(&s)
			} else {
				break
			}
		}
		var (
			steps [][]Rotation
			ok bool
		)
		if steps, ok = possibleSteps(s, r, &visited); !ok {
			return r, false
		}
		rclone := r.clone()
		for _, rot := range steps[0] {
			if rot == RT { r.right() }
			if rot == LF { r.left() }
		}
		for _, rots := range steps[1:] {
			newRobot := rclone.clone()
			for _, rot := range rots {
				if rot == RT { newRobot.right() }
				if rot == LF { newRobot.left() }
			}
			robots = robots.Push(newRobot)
		}
	}
}

var robots = utils.StackOf[Robot]()

func part1(s Scene) {
	robots = utils.StackOf[Robot]()
	r := newRobot(Coord{x: s.start.x, y: s.start.y, dir: W})
	robots = robots.Push(r)
	draw(s)
	fmt.Println("")
	for !robots.IsEmpty() {
		//fmt.Println("robot counts:", len(robots))
		r, robots = robots.Pop()
		if robot, ok := findPath(s, r); ok {
			points := robot.rotsCnt*1000 + robot.stepsCnt
			fmt.Println("Found way! Points: ", points)
		}
	}
}

func main() {
	isTest := true
	data := loadData(isTest)
	part1(parseData(data))
	//part2(parseData2(data))
}

type Rotation int
const (
	LF Rotation = iota
	RT
)

//    N
//  W X E
//    S
type Direction int
const (
	N Direction = iota
	S
	W
	E
)

type Coord struct {
	x, y int
	dir Direction
}

type Robot struct {
	path []Coord
	coord Coord // default Lf
	stepsCnt, rotsCnt int
}

func newRobot(c Coord) Robot {
	return Robot{
		path: []Coord{},
		coord: c,
	}
}

func (r Robot) clone() Robot {
	nr := newRobot(r.coord)
	nr.path = make([]Coord, 0, len(r.path))
	nr.path = append(nr.path, r.path...)
	nr.rotsCnt = r.rotsCnt
	nr.stepsCnt = r.stepsCnt
	return nr
}

func (r Robot) manhDist(s Scene) int {
	res := utils.Abs(s.finish.x - r.coord.x) +
	       utils.Abs(s.finish.y - r.coord.y)
    return res
}

func (r Robot) isFinished(s Scene) bool {
	return r.coord.x == s.finish.x && r.coord.y == s.finish.y
}

func (r Robot) canStepForward(s Scene) bool {
	c := r.frontCoord()
	return s.maze[c.y][c.x] != "#"
}

func (r Robot) frontCoord() Coord {
	switch r.coord.dir {
	case N:
		return Coord{x: r.coord.x, y: r.coord.y - 1}
	case S:
		return Coord{x: r.coord.x, y: r.coord.y + 1}
	case W:
		return Coord{x: r.coord.x - 1, y: r.coord.y}
	case E:
		return Coord{x: r.coord.x + 1, y: r.coord.y}
	}
	return Coord{}
}

// returns point Coord in front after step
func (r *Robot) step(s *Scene) Coord {
	s.maze[r.coord.y][r.coord.x] = "+"
	switch r.coord.dir {
	case N:
		r.coord.y--
	case S:
		r.coord.y++
	case W:
		r.coord.x--
	case E:
		r.coord.x++
	}
	r.stepsCnt++
	s.maze[r.coord.y][r.coord.x] = "S"
	return r.frontCoord()
}

// returns point Coord in front after rotation
func (r *Robot) left() Coord {
	switch r.coord.dir {
	case N:
		r.coord.dir = W
	case W:
		r.coord.dir = S
	case S:
		r.coord.dir = E
	case E:
		r.coord.dir = N
	}
	r.rotsCnt++
	return r.frontCoord()
}

func (r *Robot) right() Coord {
	switch r.coord.dir {
	case N:
		r.coord.dir = E
	case E:
		r.coord.dir = S
	case S:
		r.coord.dir = W
	case W:
		r.coord.dir = N
	}
	r.rotsCnt++
	return r.frontCoord()
}

type Scene struct {
	maze [][]string
	start Coord
	finish Coord
}

func parseData(data string) Scene {
	scene := Scene{}
	str_arr := strings.Split(strings.TrimSpace(data), "\n\n")
	room := strings.Split(str_arr[0], "\n")
	scene.maze = make([][]string, len(room))
	for y, line := range room {
		scene.maze[y] = make([]string, len(string(line)))
		for x, el := range string(line) {
			if el == '\n' { continue }
			scene.maze[y][x] = string(el)
			if el == 'S' {
				scene.start.x, scene.start.y = x, y
			}
			if el == 'E' {
				scene.finish.x, scene.finish.y = x, y
			}
		}
	}

	return scene
}

func loadData(wantTest bool) string {
	if wantTest {
		return `
###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############
`
	} else {
		text, err := os.ReadFile("day16/input.txt")
		if err != nil { panic(err) }
		return string(text)
	}
}

/*
 012345678901234
0###############
1#.......#....E#
2#.#.###.#.###.#
3#.....#.#...#.#
4#.###.#####.#.#
5#.#.#.......#.#
6#.#.#####.###.#
7#...........#.#
8###.#.#####.#.#
9#...#.....#.#.#
0#.#.#.###.#.#.#
1#.....#...#.#.#
2#.###.#.#.#.#.#
3#S..#.....#...#
4###############

part1: 85432
part2: 465

*/