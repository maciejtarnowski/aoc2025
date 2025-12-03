package main

import (
	"aoc2025"
	"fmt"
	"slices"
)

func main() {
	fmt.Println(solvePartOne("input.txt"))

	fmt.Println(solvePartTwo("input.txt"))
}

func solvePartOne(filename string) int64 {
	scanner, file := aoc2025.OpenFileScanner(filename)
	defer file.Close()

	maxJoltageSum := int64(0)

	for scanner.Scan() {
		bank := scanner.Text()
		maxJoltage := getMaxJoltageInBank(bank)
		maxJoltageSum += maxJoltage
	}

	return maxJoltageSum
}

func getMaxJoltageInBank(bank string) int64 {
	batteries := aoc2025.ParseListOfInts(bank, "")

	options := make([]int64, 0)
	for i := 0; i < len(batteries); i++ {
		for j := i + 1; j < len(batteries); j++ {
			options = append(options, (batteries[i]*10)+batteries[j])
		}
	}

	return slices.Max(options)
}

func solvePartTwo(filename string) int64 {
	scanner, file := aoc2025.OpenFileScanner(filename)
	defer file.Close()

	maxJoltageSum := int64(0)

	for scanner.Scan() {
		bank := scanner.Text()
		maxJoltage := getMaxJoltageInBankWithOverride(bank)
		maxJoltageSum += maxJoltage
	}

	return maxJoltageSum
}

func getMaxJoltageInBankWithOverride(bank string) int64 {
	batteries := aoc2025.ParseListOfInts(bank, "")

	solution := batteries[len(batteries)-12:]

	for i := len(batteries) - 13; i > -1; i-- {
		battery := batteries[i]

		for j := 0; j < len(solution); j++ {
			if battery >= solution[j] {
				solution[j], battery = battery, solution[j]
			} else {
				break
			}
		}
	}

	return aoc2025.IntSliceToInt64(solution)
}
