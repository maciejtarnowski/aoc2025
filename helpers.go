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

func ParseListOfInts(line string, sep string) []int64 {
	items := strings.Split(line, sep)

	ints := make([]int64, 0)

	i := 0
	for _, item := range items {
		if strings.TrimSpace(item) == "" {
			continue
		}
		num, err := strconv.ParseInt(item, 10, 64)
		if err != nil {
			panic(err)
		}
		ints = append(ints, num)
		i++
	}

	return ints
}

func IntSliceToInt64(ints []int64) int64 {
	var result int64 = 0
	for pos, v := range ints {
		result += v * powInt64(10, int64(len(ints)-pos-1))
	}
	return result
}

func powInt64(base, exp int64) int64 {
	if exp == 0 {
		return 1
	}
	result := base
	for i := int64(1); i < exp; i++ {
		result *= base
	}
	return result
}

func StrToInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}
