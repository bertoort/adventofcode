package solution

import (
	"fmt"
	"strings"
)

type Table struct {
	columns [][]int
	operators []string
}

func newTable(input []string) *Table {
	table := &Table{}
	for i, r := range input {
		row := strings.Fields(string(r))
		if i == 0 {
			table.operators = make([]string, len(row))
			table.columns = make([][]int, len(row))
		} 
		for j, cell := range row {
			if i == len(input) - 1 {
				table.operators[j] = string(cell)
			} else {
				if (table.columns[j] == nil) {
					table.columns[j] = make([]int, len(input)-1)
				}
				table.columns[j][i] = parseInt(string(cell))
			}
		}
	}
	return table
}

func performOperation(a, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b != 0 {
			return a / b
		}
		fmt.Println("Error: Division by zero")
		return 0
	default:
		fmt.Printf("Unknown operator: %s\n", operator)
		return 0
	}
}

func (t *Table) totalColumns() []int {
	totals := make([]int, len(t.columns))
	for i, column := range t.columns {
		for j, cell := range column {
			if j == 0 {
				totals[i] = cell
				continue
			}
			totals[i] = performOperation(totals[i], cell, t.operators[i])
		}
	}
	return totals
}

func sumTotals(totals []int) int {
	sum := 0
	for _, total := range totals {
		sum += total
	}
	return sum
}

// Solution 2

type Operators struct {
	indices []int
	operators []string
}

func parseOperators(row string) *Operators {	
	operators := []string{}
	indices := make([]int, len(operators))
	for i, r := range row {
		if r == ' ' {
			continue
		}
		operators = append(operators, string(r))
		indices = append(indices, i)
	}
	return &Operators{indices: indices, operators: operators}
}

func convertColumns(columns [][]string) [][]int {
	newColumns := make([][]int, len(columns))
	for i, column := range columns {
		newColumns[i] = []int{}
		for _, cell := range column {
			cell = strings.TrimSpace(cell)
			if cell == "" {
				continue
			}
			newColumns[i] = append(newColumns[i], parseInt(cell))
		}
	}
	return newColumns
}

func parseColumns(input []string, indices []int) [][]string {
	columns := make([][]string, len(indices))
	for _, r := range input {
		columnIndex := 0
		rowIndex := 0
		for j, cell := range r {
			if columns[columnIndex] == nil {
				columns[columnIndex] = make([]string, len(input))
			}
			// Size is the remaining length of the row
			size := len(r)-j-1
			// Unless we're not at the last column, use the next index as the size
			if columnIndex != len(indices) - 1 {
				size = indices[columnIndex+1]
			}
			// The end of the column is blank, move on to the next
			if j == size -1 {
				columnIndex++
				rowIndex = 0
				continue
			}
			columns[columnIndex][rowIndex] = columns[columnIndex][rowIndex] + string(cell)
			rowIndex++
		}
	}
	return columns
}

func newTable2(input []string) *Table {
	// Grab operators and spread indices for each column
	o := parseOperators(input[len(input)-1])
	// Parse columns based on column sizes
	c := parseColumns(input[:len(input)-1], o.indices)
	fmt.Println(c)
	// Trim and convert columns to ints
	columns := convertColumns(c)
	return &Table{ columns: columns, operators: o.operators}
}