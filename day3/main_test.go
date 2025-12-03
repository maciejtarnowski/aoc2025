package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePartOneDemo(t *testing.T) {
	result := solvePartOne("demo.txt")

	assert.Equal(t, int64(357), result)
}

func TestGetMaxJoltageInBank(t *testing.T) {
	tests := []struct {
		name     string
		bank     string
		expected int64
	}{
		{"Bank 987654321111111", "987654321111111", 98},
		{"Bank 811111111111119", "811111111111119", 89},
		{"Bank 234234234234278", "234234234234278", 78},
		{"Bank 818181911112111", "818181911112111", 92},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getMaxJoltageInBank(tt.bank)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSolvePartTwoDemo(t *testing.T) {
	result := solvePartTwo("demo.txt")

	assert.Equal(t, int64(3121910778619), result)
}

func TestGetMaxJoltageInBankWithOverride(t *testing.T) {
	tests := []struct {
		name     string
		bank     string
		expected int64
	}{
		{"Bank 987654321111111", "987654321111111", 987654321111},
		{"Bank 811111111111119", "811111111111119", 811111111119},
		{"Bank 234234234234278", "234234234234278", 434234234278},
		{"Bank 818181911112111", "818181911112111", 888911112111},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getMaxJoltageInBankWithOverride(tt.bank)
			assert.Equal(t, tt.expected, result)
		})
	}
}
