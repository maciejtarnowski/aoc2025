package main

import (
	"aoc2025"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(solvePartOne("input.txt"))

	fmt.Println(solvePartTwo("input.txt"))
}

type Dial struct {
	position         int64
	touchedZeroCount int
}

func NewDial() *Dial {
	return &Dial{position: 50, touchedZeroCount: 0}
}

func (d *Dial) ApplyMove(move string) {
	step, err := strconv.ParseInt(move[1:], 10, 64)
	if err != nil {
		panic(err)
	}
	switch move[0] {
	case 'L':
		d.Move(-step)
	case 'R':
		d.Move(step)
	}
}

func (d *Dial) Move(step int64) {
	for range abs(step) {
		if step > 0 {
			d.position++
			if d.position > 99 {
				d.position = 0
			}
			if d.IsAtZero() {
				d.touchedZeroCount++
			}
		} else {
			d.position--
			if d.position < 0 {
				d.position = 99
			}
			if d.IsAtZero() {
				d.touchedZeroCount++
			}
		}
	}
}

func (d *Dial) IsAtZero() bool {
	return d.position == 0
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func solvePartOne(filename string) int {
	scanner, file := aoc2025.OpenFileScanner(filename)
	defer file.Close()

	dial := NewDial()
	atZeroCount := 0

	for scanner.Scan() {
		line := scanner.Text()

		dial.ApplyMove(line)
		if dial.IsAtZero() {
			atZeroCount++
		}
	}

	return atZeroCount
}

func solvePartTwo(filename string) int {
	scanner, file := aoc2025.OpenFileScanner(filename)
	defer file.Close()

	dial := NewDial()

	for scanner.Scan() {
		line := scanner.Text()

		dial.ApplyMove(line)
	}

	return dial.touchedZeroCount
}
