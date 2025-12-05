package main

import (
	"aoc2025"
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fmt.Println(solvePartOne("input.txt"))

	fmt.Println(solvePartTwo("input.txt"))
}

func solvePartOne(filename string) int {
	db := NewDatabaseFromFile(filename)

	scanner, file := aoc2025.OpenFileScanner(filename)
	defer file.Close()

	rangesSkipped := false
	freshIDsCount := 0
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			rangesSkipped = true
			continue
		}
		if !rangesSkipped {
			continue
		}

		if db.isIDFresh(aoc2025.StrToInt64(line)) {
			freshIDsCount++
		}
	}

	return freshIDsCount
}

func solvePartTwo(filename string) int {
	db := NewDatabaseFromFile(filename)
	db.foldOverlappingRanges()

	freshIDCount := 0
	for _, r := range db.freshRanges {
		freshIDCount += int(r.end - r.start + 1)
	}

	return freshIDCount
}

type IDRange struct {
	start int64
	end   int64
}

func (r *IDRange) contains(id int64) bool {
	return id >= r.start && id <= r.end
}

type Database struct {
	freshRanges []*IDRange
}

func (db *Database) addFreshRange(start, end int64) {
	newRange := &IDRange{start: start, end: end}

	db.freshRanges = append(db.freshRanges, newRange)
}

func (db *Database) isIDFresh(id int64) bool {
	for _, r := range db.freshRanges {
		if r.contains(id) {
			return true
		}
	}
	return false
}

func (db *Database) foldOverlappingRanges() {
	sortedRangesByStart := make([]*IDRange, len(db.freshRanges))
	copy(sortedRangesByStart, db.freshRanges)

	slices.SortFunc(sortedRangesByStart, func(a, b *IDRange) int {
		return cmp.Compare(a.start, b.start)
	})

	foldedRanges := make([]*IDRange, 0)
	currentRange := sortedRangesByStart[0]

	for i := 1; i < len(sortedRangesByStart); i++ {
		nextRange := sortedRangesByStart[i]

		if currentRange.end >= nextRange.start {
			if nextRange.end > currentRange.end {
				currentRange.end = nextRange.end
			}
		} else {
			foldedRanges = append(foldedRanges, currentRange)
			currentRange = nextRange
		}
	}

	foldedRanges = append(foldedRanges, currentRange)

	db.freshRanges = foldedRanges
}

func NewDatabaseFromFile(filename string) *Database {
	scanner, file := aoc2025.OpenFileScanner(filename)
	defer file.Close()

	db := &Database{
		freshRanges: make([]*IDRange, 0),
	}

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			break
		}

		var start, end int64
		n, err := fmt.Sscanf(line, "%d-%d", &start, &end)
		if err != nil {
			panic(err)
		}
		if n != 2 {
			panic("invalid range, must scan 2 items")
		}

		db.addFreshRange(start, end)
	}

	return db
}
