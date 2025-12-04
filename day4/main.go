package main

import (
	"aoc2025"
	"fmt"
)

func main() {
	fmt.Println(solvePartOne("input.txt"))

	fmt.Println(solvePartTwo("input.txt"))
}

func solvePartOne(filename string) int {
	f := NewFloorMapFromFile(filename)

	paperRollsAccessible := 0
	for y := 0; y < f.ySize; y++ {
		for x := 0; x < f.xSize; x++ {
			if f.GetObjectAt(x, y) == PaperRoll {
				if f.IsAccessiblePaperRoll(x, y) {
					paperRollsAccessible++
				}
			}
		}
	}

	return paperRollsAccessible
}

func solvePartTwo(filename string) int {
	f := NewFloorMapFromFile(filename)

	paperRollsRemoved := 0

	for {
		initialPaperRollsRemoved := paperRollsRemoved
		for y := 0; y < f.ySize; y++ {
			for x := 0; x < f.xSize; x++ {
				if f.GetObjectAt(x, y) == PaperRoll {
					if f.IsAccessiblePaperRoll(x, y) {
						f.SetObjectAt(x, y, RemovedPaperRoll)
						paperRollsRemoved++
					}
				}
			}
		}
		if initialPaperRollsRemoved == paperRollsRemoved {
			break
		}
	}

	return paperRollsRemoved
}

type Object int

const (
	Empty Object = iota
	PaperRoll
	RemovedPaperRoll
)

type FloorMap struct {
	floor map[int]Object
	xSize int
	ySize int
}

func (f *FloorMap) GetObjectAt(x, y int) Object {
	return f.floor[y*f.xSize+x]
}

func (f *FloorMap) SetObjectAt(x, y int, obj Object) {
	f.floor[y*f.xSize+x] = obj
}

func (f *FloorMap) InitXSize(xSize int) {
	if f.xSize == 0 {
		f.xSize = xSize
	}
}

func (f *FloorMap) InitYSize(ySize int) {
	if f.ySize == 0 {
		f.ySize = ySize
	}
}

func (f *FloorMap) IsAccessiblePaperRoll(x, y int) bool {
	tilesToCheck := [][2]int{
		{x, y - 1},
		{x, y + 1},
		{x - 1, y},
		{x + 1, y},
		{x - 1, y - 1},
		{x + 1, y - 1},
		{x - 1, y + 1},
		{x + 1, y + 1},
	}

	paperRollsFound := 0
	for _, tile := range tilesToCheck {
		tx, ty := tile[0], tile[1]
		if tx >= 0 && tx < f.xSize && ty >= 0 && ty < f.ySize {
			if f.GetObjectAt(tx, ty) == PaperRoll {
				paperRollsFound++
			}
		}
	}

	return paperRollsFound < 4
}

func NewFloorMapFromFile(filename string) *FloorMap {
	scanner, file := aoc2025.OpenFileScanner(filename)
	defer file.Close()

	f := &FloorMap{}
	f.floor = make(map[int]Object)
	f.InitYSize(aoc2025.CountLinesInFile(filename))

	lineC := 0
	for scanner.Scan() {
		line := scanner.Text()
		f.InitXSize(len(line))
		for x, char := range line {
			if char == '@' {
				f.SetObjectAt(x, lineC, PaperRoll)
			} else {
				f.SetObjectAt(x, lineC, Empty)
			}
		}
		lineC++
	}

	return f
}
