package aoc2025

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func OpenFileScanner(path string) (scanner *bufio.Scanner, file *os.File) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	scan := bufio.NewScanner(file)

	return scan, file
}

func CountLinesInFile(path string) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := 0
	for scanner.Scan() {
		lines++
	}

	return lines
}

func ParseSpaceSeparatedListOfInts(line string) []int64 {
	items := strings.Split(line, " ")

	ints := make([]int64, len(items))

	for i, item := range items {
		num, err := strconv.ParseInt(item, 10, 64)
		if err != nil {
			panic(err)
		}
		ints[i] = num
	}

	return ints
}
