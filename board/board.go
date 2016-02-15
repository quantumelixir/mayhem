package board

import "github.com/quantumelixir/mayhem/types"

func NewBoard(rows int, cols int, fillvalue types.Color) (board [][]types.Color) {

	// initialize the board
	board = make([][]types.Color, rows)

	for i := 0; i < rows; i++ {
		board[i] = make([]types.Color, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			board[i][j] = fillvalue
		}
	}

	return board
}
