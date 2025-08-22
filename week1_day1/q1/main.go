package main

import "fmt"

type Matrix struct {
	rows int
	cols int
	Data [][]int
}

func (m Matrix) GetRows() {
	fmt.Println("Number of rows are: ", m.rows)
}

func (m Matrix) GetColumns() {
	fmt.Println("Number of columns are: ", m.cols)
}

func (m Matrix) SetElement(i, j, value) {
	if i >= 0 && i < m.rows && j >= 0 && j < m.cols {
		m.Data[i][j] = value
	}
}
