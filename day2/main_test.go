package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePartOneDemo(t *testing.T) {
	result := solvePartOne("demo.txt")

	assert.Equal(t, 1227775554, result)
}

func TestIsInvalidIDPartOne(t *testing.T) {
	tests := []struct {
		name     string
		id       int
		expected bool
	}{
		{"valid: 12", 12, false},
		{"invalid: 11", 11, true},
		{"invalid: 1010", 1010, true},
		{"invalid: 1188511885", 1188511885, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isInvalidIDPartOne(tt.id)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestInvalidIDsInRangePartOne(t *testing.T) {
	tests := []struct {
		name      string
		rangeSpec string
		expected  []int
	}{
		{"range 11-22", "11-22", []int{11, 22}},
		{"range 95-115", "95-115", []int{99}},
		{"range 1188511880-1188511890", "1188511880-1188511890", []int{1188511885}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := invalidIDsInRange(tt.rangeSpec, isInvalidIDPartOne)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSolvePartTwoDemo(t *testing.T) {
	result := solvePartTwo("demo.txt")

	assert.Equal(t, 4174379265, result)
}

func TestIsInvalidIDPartTwo(t *testing.T) {
	tests := []struct {
		name     string
		id       int
		expected bool
	}{
		{"valid: 12", 12, false},
		{"invalid: 11", 11, true},
		{"invalid: 999", 999, true},
		{"invalid: 1010", 1010, true},
		{"invalid: 1188511885", 1188511885, true},
		{"invalid: 2121212121", 2121212121, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isInvalidIDPartTwo(tt.id)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestInvalidIDsInRangePartTwo(t *testing.T) {
	tests := []struct {
		name      string
		rangeSpec string
		expected  []int
	}{
		{"range 11-22", "11-22", []int{11, 22}},
		{"range 95-115", "95-115", []int{99, 111}},
		{"range 998-1012", "998-1012", []int{999, 1010}},
		{"range 1188511880-1188511890", "1188511880-1188511890", []int{1188511885}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := invalidIDsInRange(tt.rangeSpec, isInvalidIDPartTwo)
			assert.Equal(t, tt.expected, result)
		})
	}
}
