package main

import (
	"errors"
	"fmt"
)

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

type sparseMatrix[T Number] struct {
	rows, cols int
	data       map[[2]int]T
}

func NewSparseMatrix[T Number](rows, cols int) *sparseMatrix[T] {
	return &sparseMatrix[T]{
		rows: rows,
		cols: cols,
		data: make(map[[2]int]T),
	}
}

func (sm *sparseMatrix[T]) validate(row, col int) {
	if sm == nil {
		panic("Matrix is nil")
	}
	if row < 0 || row >= sm.rows || col < 0 || col >= sm.cols {
		panic("Invalid indices: row and col must be within matrix dimensions")
	}
}
func (sm *sparseMatrix[T]) Set(row, col int, value T) {
	sm.validate(row, col)
	sm.data[[2]int{row, col}] = value
}

func (sm *sparseMatrix[T]) Get(row, col int) (T, error) {
	if sm.IsEmpty() {
		return *new(T), errors.New("matrix is empty")
	}
	sm.validate(row, col)

	if value, ok := sm.data[[2]int{row, col}]; ok {
		return value, nil
	}
	return *new(T), nil
}

func (sm *sparseMatrix[T]) IsEmpty() bool {
	return len(sm.data) == 0
}

func main() {
	sm := NewSparseMatrix[float64](100, 200)

	sm.Set(0, 0, 5.0)
	sm.Set(0, 199, 8.123456)
	sm.Set(49, 99, 777)
	sm.Set(99, 0, 9)
	sm.Set(99, 199, 10.0)

	fmt.Println(sm.Get(0, 0))
	fmt.Println(sm.Get(1, 2))
	fmt.Println(sm.Get(0, 199))

}
