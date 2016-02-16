package main

func NewBoard(rows int, cols int, fillvalue Color) (board [][]Color) {

	// initialize the board
	board = make([][]Color, rows)

	for i := 0; i < rows; i++ {
		board[i] = make([]Color, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			board[i][j] = fillvalue
		}
	}

	return board
}
