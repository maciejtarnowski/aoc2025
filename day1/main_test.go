package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDialApplyMove(t *testing.T) {
	tests := []struct {
		name     string
		move     string
		expected int
	}{
		{"Move left 10", "L10", 40},
		{"Move right 20", "R20", 70},
		{"Move left 60", "L60", 90},
		{"Move right 50", "R50", 0},
		{"Move left 100", "L100", 50},
		{"Move right 100", "R100", 50},
		{"Move left 110", "L110", 40},
		{"Move right 120", "R120", 70},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dial := NewDial()
			dial.ApplyMove(tt.move)
			if dial.position != int64(tt.expected) {
				t.Errorf("expected position %d, got %d", tt.expected, dial.position)
			}
		})
	}
}

func TestSolvePartOneDemo(t *testing.T) {
	result := solvePartOne("demo.txt")

	assert.Equal(t, 3, result)
}

func TestSolvePartTwoDemo(t *testing.T) {
	result := solvePartTwo("demo.txt")

	assert.Equal(t, 6, result)
}
