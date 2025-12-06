package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePartOneDemo(t *testing.T) {
	result := solvePartOne("demo.txt")

	assert.Equal(t, int64(4277556), result)
}

func TestSolvePartTwoDemo(t *testing.T) {
	result := solvePartTwo("demo.txt")

	assert.Equal(t, int64(3263827), result)
}
