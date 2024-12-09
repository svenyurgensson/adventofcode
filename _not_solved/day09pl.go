package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

const space = -1

func SolveP1() {
	lines := ReadFile("day09/input.txt")

	puzzle := lines[0]
	fmt.Println("puzzle input: ", puzzle)

	hardDisk, _ := mapHardDisk(puzzle)
	fmt.Println("hardDisk before: ", hardDisk)

	hardDisk = emptySpace(hardDisk)
	fmt.Println("hardDisk after: ", hardDisk)

	checksum := calculateFileSystemChecksum(hardDisk)
	fmt.Println("checksum: ", checksum)
}

func main() {
	SolveP2()
}

func SolveP2() {
	lines := ReadFile("day09/input.txt")

	puzzle := lines[0]
	fmt.Println("puzzle input: ", puzzle)

	hardDisk, hardDiskMap := mapHardDisk(puzzle)
	fmt.Println("hardDisk before: ", hardDisk)
	fmt.Println("hardDiskMap: ", hardDiskMap)

	hardDisk = emptySpaceWithoutFragmentation(hardDisk, hardDiskMap)
	fmt.Println("hardDisk after: ", hardDisk)

	checksum := calculateFileSystemChecksum(hardDisk)
	fmt.Println("checksum: ", checksum)
}

func getLastNonEmptyValue(disk []int) (value, idxOfvalue int) {
	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i] != space {
			return disk[i], i
		}
	}

	return -1, -1
}

func emptySpace(disk []int) []int {

	newDiskLen := 0
	for _, value := range disk {
		if value != space {
			newDiskLen++
		}
	}

	emptiedDisk := []int{}
	for _, value := range disk {

		if newDiskLen == len(emptiedDisk) {
			break
		}

		if value == -1 {
			lastValue, indexOfLastValue := getLastNonEmptyValue(disk)

			// remove last value from original disk
			disk = disk[:indexOfLastValue]

			emptiedDisk = append(emptiedDisk, lastValue)
		} else {
			emptiedDisk = append(emptiedDisk, value)
		}
	}
	return emptiedDisk
}

type hardDiskMap map[int]hardDiskValue

type hardDiskValue struct {
	value int
	idx   int
}

func mapHardDisk(puzzle string) ([]int, hardDiskMap) {
	hardDisk := []int{}
	hardDiskMap := make(hardDiskMap)

	for idx, value := range puzzle {
		numValue, _ := strconv.Atoi(string(value))

		isSpace := idx%2 != 0

		var valToAppend int
		if isSpace {
			valToAppend = space
		} else {
			valToAppend = idx / 2
		}

		if !isSpace {
			hardDiskMap[valToAppend] = hardDiskValue{value: numValue, idx: -1}
		}
		for i := 0; i < numValue; i++ {
			hardDisk = append(hardDisk, valToAppend)
		}

	}

	for idx, value := range hardDisk {
		if value != space {
			if hardDiskMap[value].idx == -1 {
				val := hardDiskMap[value]
				val.idx = idx
				hardDiskMap[value] = val
			}
		}
	}

	return hardDisk, hardDiskMap
}

func getNextEmptyValue(idx int, disk []int, byteLength int, byteCurrentPosition int) (pos int, posFound bool) {

	if idx == len(disk) {
		return 0, false
	}

	foundPos := false
	position := 0
	spaceLength := 0
	for i := idx; i < len(disk); i++ {
		if disk[i] == space {
			if !foundPos {
				position = i
				foundPos = true
			}

			spaceLength++
		} else if foundPos {
			break
		}
	}

	if spaceLength < byteLength || position > byteCurrentPosition {
		return getNextEmptyValue(position+spaceLength, disk, byteLength, byteCurrentPosition)
	}

	return position, true
}

func removeByteFromDisk(disk *[]int, byteToRemove int) {
	for idx, value := range *disk {

		if value == byteToRemove {
			(*disk)[idx] = space
		} else {
			continue
		}
	}
}

// without fragmenting the disk
func emptySpaceWithoutFragmentation(disk []int, diskMap hardDiskMap) []int {

	type bytes struct {
		idx   int
		len   int
		value int
	}

	diskMapList := []bytes{}
	for key, value := range diskMap {
		diskMapList = append(diskMapList, bytes{idx: value.idx, len: value.value, value: key})
	}

	// reverse sort by value
	slices.SortFunc(diskMapList, func(a, b bytes) int {
		return b.value - a.value
	})

	for _, byteToMove := range diskMapList {
		pos, posFound := getNextEmptyValue(0, disk, byteToMove.len, byteToMove.idx)
		if !posFound {
			continue
		}

		removeByteFromDisk(&disk, byteToMove.value)

		for p := range byteToMove.len {
			disk[pos+p] = byteToMove.value
		}
	}

	//6636608781232
	return disk
}

func calculateFileSystemChecksum(disk []int) int {
	sum := 0
	for idx, value := range disk {
		if value != space {
			sum += value * idx
		}
	}

	return sum
}

func ReadFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}