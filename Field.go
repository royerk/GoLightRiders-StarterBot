package main

import "strings"

const empty int = 0
const player1 int = 1
const player2 int = 2
const blocked int = 3

// Field class
type Field struct {
	width, height int
	cells         [][]int // I will use [row][col], not x,y
}

func (f *Field) initField(h, w int) {
	f.width = w
	f.height = h
	f.cells = make([][]int, f.height)
	for row := 0; row < f.height; row++ {
		f.cells[row] = make([]int, f.width) // default initialization with 0s
	}
}
func stringToInt(s string) int {
	switch s {
	case ".":
		return empty
	case "0":
		return player1
	case "1":
		return player2
	default:
		return blocked
	}
}
func (f *Field) parse(text string) {
	values := strings.Split(text, ",")
	for row := 0; row < f.height; row++ {
		for col := 0; col < f.width; col++ {
			f.cells[row][col] = stringToInt(values[row*f.height+col])
		}
	}
}
func (f Field) isValid(row, col int) bool {
	if row < 0 || row >= f.height {
		return false
	}
	if col < 0 || col >= f.width {
		return false
	}
	return true
}
