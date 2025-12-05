package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePartOneDemo(t *testing.T) {
	result := solvePartOne("demo.txt")

	assert.Equal(t, 3, result)
}

func TestBuildsDatabaseFromFile(t *testing.T) {
	db := NewDatabaseFromFile("demo.txt")

	assert.Equal(t, 4, len(db.freshRanges))

	assert.Equal(t, int64(3), db.freshRanges[0].start)
	assert.Equal(t, int64(5), db.freshRanges[0].end)

	assert.Equal(t, int64(10), db.freshRanges[1].start)
	assert.Equal(t, int64(14), db.freshRanges[1].end)

	assert.Equal(t, int64(16), db.freshRanges[2].start)
	assert.Equal(t, int64(20), db.freshRanges[2].end)

	assert.Equal(t, int64(12), db.freshRanges[3].start)
	assert.Equal(t, int64(18), db.freshRanges[3].end)
}

func TestSolvePartTwoDemo(t *testing.T) {
	result := solvePartTwo("demo.txt")

	assert.Equal(t, 14, result)
}
