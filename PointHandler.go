package main

/*
import (
	"fmt"
	"math/rand"
	"time"
)

const (
	rows    = 10
	columns = 10
)

func main() {
	// Seed the random number generator to get different initial patterns
	rand.Seed(time.Now().UnixNano())

	// Create a 2D grid to represent the game board
	board := make([][]bool, rows)
	for i := range board {
		board[i] = make([]bool, columns)
	}

	// Initialize the board with random initial state
	initializeBoard(board)

	// Run the game for a number of generations
	generations := 10
	for generation := 0; generation < generations; generation++ {
		fmt.Printf("Generation %d:\n", generation+1)
		printBoard(board)
		board = evolve(board)
		time.Sleep(1 * time.Second)
	}
}

// Initialize the board with random initial state
func initializeBoard(board [][]bool) {
	for i := range board {
		for j := range board[i] {
			// Randomly assign cells as alive (true) or dead (false)
			board[i][j] = rand.Intn(2) == 1
		}
	}
}

// Evolve the board to the next generation based on the rules of the game
func evolve(board [][]bool) [][]bool {
	newBoard := make([][]bool, rows)
	for i := range newBoard {
		newBoard[i] = make([]bool, columns)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			liveNeighbors := countLiveNeighbors(board, i, j)
			if board[i][j] {
				// Any live cell with fewer than two live neighbors dies
				// Any live cell with more than three live neighbors dies
				newBoard[i][j] = liveNeighbors == 2 || liveNeighbors == 3
			} else {
				// Any dead cell with exactly three live neighbors becomes alive
				newBoard[i][j] = liveNeighbors == 3
			}
		}
	}

	return newBoard
}

// Count the number of live neighbors of a cell
func countLiveNeighbors(board [][]bool, x, y int) int {
	liveNeighbors := 0
	dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	for i := 0; i < 8; i++ {
		nx, ny := x+dx[i], y+dy[i]
		if nx >= 0 && nx < rows && ny >= 0 && ny < columns && board[nx][ny] {
			liveNeighbors++
		}
	}

	return liveNeighbors
}

// Print the current state of the board
func printBoard(board [][]bool) {
	for i := range board {
		for j := range board[i] {
			if board[i][j] {
				fmt.Print("■ ")
			} else {
				fmt.Print("□ ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
*/
