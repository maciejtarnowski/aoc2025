package main

import (
	"aoc2025"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(solvePartOne("input.txt"))

	fmt.Println(solvePartTwo("input.txt"))
}

func solvePartOne(filename string) int64 {
	scanner, file := aoc2025.OpenFileScanner(filename)
	defer file.Close()

	lines := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	operations := make([]Calculation, 0)

	operands := strings.Split(lines[len(lines)-1], " ")

	opIndex := 0
	for _, operand := range operands {
		if strings.TrimSpace(operand) == "" {
			continue
		}
		if operand == "+" {
			operations = append(operations, &Addition{operands: make([]int64, 0)})
		}
		if operand == "*" {
			operations = append(operations, &Multiplication{operands: make([]int64, 0)})
		}
		opIndex++
	}

	for _, line := range lines[:len(lines)-1] {
		nums := aoc2025.ParseListOfInts(line, " ")
		for i, num := range nums {
			operations[i].AddOperand(num)
		}
	}

	total := int64(0)

	for _, operation := range operations {
		total += operation.Compute()
	}

	return total
}

func solvePartTwo(filename string) int64 {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lineWrapIndex := 0
	for i, b := range content {
		if b == '\n' {
			lineWrapIndex = i
			break
		}
	}

	operations := make([]Calculation, 0)
	operands := make([]int64, 0)

	column := lineWrapIndex - 1
	for column >= 0 {
		digits := make([]string, 0)
		for row := 0; row < len(content)/lineWrapIndex; row++ {
			index := row*lineWrapIndex + column + row
			if content[index] == ' ' {
				continue
			} else if content[index] == '+' {
				operands = append(operands, digitsToOperand(digits))
				op := &Addition{operands: operands}
				operations = append(operations, op)
				operands = make([]int64, 0)
				digits = make([]string, 0)
				// sign is the last char in the last column (r-to-l), so move one extra column to the left (skipping spaces)
				column--
				break
			} else if content[index] == '*' {
				operands = append(operands, digitsToOperand(digits))
				op := &Multiplication{operands: operands}
				operations = append(operations, op)
				operands = make([]int64, 0)
				digits = make([]string, 0)
				// sign is the last char in the last column (r-to-l), so move one extra columns to the left (skipping spaces)
				column--
				break
			} else {
				digits = append(digits, string(content[index]))
			}
		}
		if len(digits) > 0 {
			operands = append(operands, digitsToOperand(digits))
		}
		column--
	}

	total := int64(0)
	for _, operation := range operations {
		total += operation.Compute()
	}
	return total
}

type Calculation interface {
	Compute() int64
	AddOperand(operand int64)
}

type Addition struct {
	operands []int64
}

type Multiplication struct {
	operands []int64
}

func (a *Addition) Compute() int64 {
	sum := int64(0)
	for _, op := range a.operands {
		sum += op
	}
	return sum
}

func (m *Multiplication) Compute() int64 {
	product := int64(1)
	for _, op := range m.operands {
		product *= op
	}
	return product
}

func (a *Addition) AddOperand(operand int64) {
	a.operands = append(a.operands, operand)
}

func (m *Multiplication) AddOperand(operand int64) {
	m.operands = append(m.operands, operand)
}

func digitsToOperand(digits []string) int64 {
	str := strings.Join(digits, "")
	return aoc2025.StrToInt64(str)
}
