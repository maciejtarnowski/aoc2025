package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(solvePartOne("input.txt"))

	fmt.Println(solvePartTwo("input.txt"))
}

func solvePartOne(filename string) int {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	rangeSpecs := strings.Split(string(data), ",")

	invalidIDsSum := 0
	for _, rangeSpec := range rangeSpecs {
		invalidIDs := invalidIDsInRange(strings.TrimSpace(rangeSpec), isInvalidIDPartOne)
		for _, id := range invalidIDs {
			invalidIDsSum += id
		}
	}

	return invalidIDsSum
}

func isInvalidIDPartOne(id int) bool {
	idStr := fmt.Sprintf("%d", id)
	if len(idStr)%2 != 0 {
		return false
	}
	firstHalf := idStr[:len(idStr)/2]
	secondHalf := idStr[len(idStr)/2:]

	for i := 0; i < len(firstHalf); i++ {
		if firstHalf[i] != secondHalf[i] {
			return false
		}
	}

	return true
}

func solvePartTwo(filename string) int {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	rangeSpecs := strings.Split(string(data), ",")

	invalidIDsSum := 0
	for _, rangeSpec := range rangeSpecs {
		invalidIDs := invalidIDsInRange(strings.TrimSpace(rangeSpec), isInvalidIDPartTwo)
		for _, id := range invalidIDs {
			invalidIDsSum += id
		}
	}

	return invalidIDsSum
}

func isInvalidIDPartTwo(id int) bool {
	idStr := fmt.Sprintf("%d", id)

	for i := 0; i < len(idStr)/2; i++ {
		count := strings.Count(idStr, idStr[:i+1])
		if count*len(idStr[:i+1]) == len(idStr) && count > 1 {
			return true
		}
	}

	return false
}

func invalidIDsInRange(rangeSpec string, predicate func(id int) bool) []int {
	var invalidIDs []int

	var start, end int
	n, err := fmt.Sscanf(rangeSpec, "%d-%d", &start, &end)
	if err != nil {
		panic(err)
	}
	if n != 2 {
		panic("must scan 2 items")
	}

	for id := start; id <= end; id++ {
		if predicate(id) {
			invalidIDs = append(invalidIDs, id)
		}
	}

	return invalidIDs
}
