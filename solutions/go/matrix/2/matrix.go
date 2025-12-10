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
	rowSize := 0
	columnSize := 0
	newMatrix := Matrix{}
	for _, v := range slice {
		if len(strings.TrimSpace(v)) == 0 {
			return newMatrix, errors.New("empty row")
		}
		newLine := strings.Fields(v)
		if rowSize == 0 {
			rowSize = len(newLine)
		} else if rowSize != len(newLine) {
			return newMatrix, errors.New("uneven rows")
		}
		columnSize++
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

	newMatrix.columns = make([][]int, rowSize)

	for i := 0; i < rowSize; i++ {
		newMatrix.columns[i] = make([]int, columnSize)
		for j := 0; j < columnSize; j++ {
			newMatrix.columns[i][j] = newMatrix.rows[j][i]
		}
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
