package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix struct {
	rows    [][]int
	columns [][]int
}

func New(s string) (Matrix, error) {
	slice := strings.Split(s, "\n")
	newMatrix := Matrix{}
	for _, v := range slice {
		newLine := strings.Fields(v)
		row := make([]int, len(newLine))
		for j, num := range newLine {
			n, err := strconv.Atoi(num)
			if err != nil {
				return newMatrix, errors.New("invalid input")
			}
			row[j] = n
		}
		newMatrix.rows = append(newMatrix.rows, row)
	}
	for _, v := range newMatrix.rows {
		if len(v) != len(newMatrix.rows[0]) {
			return newMatrix, errors.New("not a valid matrix")
		}
	}
	for i := 0; i < len(newMatrix.rows[0]); i++ {
		column := make([]int, len(newMatrix.rows))
		for j := 0; j < len(newMatrix.rows); j++ {
			column[j] = newMatrix.rows[j][i]
		}
		newMatrix.columns = append(newMatrix.columns, column)
	}
	return newMatrix, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	col := make([][]int, len(m.columns))
	for i := range m.columns {
		col[i] = make([]int, len(m.columns[i]))
		copy(col[i], m.columns[i])
	}
	return col
}

func (m Matrix) Rows() [][]int {
	row := make([][]int, len(m.rows))
	for i := range m.rows {
		row[i] = make([]int, len(m.rows[i]))
		copy(row[i], m.rows[i])
	}
	return row
}

func (m Matrix) Set(row, col, val int) bool {
	if row < 0 || col < 0 {
		return false
	}
	if len(m.rows) <= row || len(m.rows[row]) <= col {
		return false
	}
	m.rows[row][col] = val
	m.columns[col][row] = val
	return true
}
