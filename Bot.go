package main

import (
	"fmt"
	"math/rand"
	"os"
)

// Bot here goes Hal
type Bot struct {
	row, col int
}

func (b Bot) play(g Game) {
	for row := 0; row < g.fieldHeight; row++ {
		for col := 0; col < g.fieldWidth; col++ {
			if g.field.cells[row][col] == g.myBotID {
				b.row = row
				b.col = col
			}
		}
	}
	fmt.Fprintln(os.Stderr, "bot position: row (y) = ", b.row, ", col (x) = ", b.col)
	// find a random legal move
	moves := b.getValidMoves(g.field)
	if len(moves) == 0 {
		fmt.Fprintln(os.Stderr, "blocked, print \"up\"")
		fmt.Println("up") // bot is blocked
	} else {
		fmt.Fprintln(os.Stderr, "available moves: ", moves)
		fmt.Println(moves[rand.Intn(len(moves))])
	}
}
func (b Bot) getValidMoves(f Field) map[int]string {
	result := make(map[int]string)
	index := 0
	if f.isValid(b.row+1, b.col) {
		result[index] = "down"
		index++
	}
	if f.isValid(b.row-1, b.col) {
		result[index] = "up"
		index++
	}
	if f.isValid(b.row, b.col+1) {
		result[index] = "right"
		index++
	}
	if f.isValid(b.row, b.col-1) {
		result[index] = "left"
	}
	return result
}
