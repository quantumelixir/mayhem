package main

import (
	"fmt"
	"github.com/quantumelixir/mayhem/robot"
	"github.com/quantumelixir/mayhem/board"
)

const LIM = 16

// * Init
func init() {
	// initialize the board
	board.Board = make([][]robot.Color, LIM)
	for i := 0; i < LIM; i++ {
		board.Board[i] = make([]robot.Color, LIM)
	}

	for i := range board.Board {
		for j := range board.Board[i] {
			board.Board[i][j] = robot.Blue
		}
	}

	// making a 12 x 16 board
	board.Board[0][0] = robot.Red
	board.Board[4][11] = robot.Green
	board.Board[11][15] = robot.Red
	board.Board[5][8] = robot.Green
	board.Board[2][2] = robot.Red
	board.Board[2][5] = robot.Red
}

// * Main
func main() {
	var r robot.Robot

	r.Init(11, 0, robot.Up)
	r.DeclareFunctionList([]string{"F1", "F2"})
	r.DefineFunction("F1", []string{":F2", ":R", "Blue:F1"})
	r.DefineFunction("F2", []string{":F", "Red:R", "Green:L", "Blue:F2", ":F"})
	fmt.Println(r.FunctionMap["F1"])
	fmt.Println(r.FunctionMap["F2"])
	robot.Run(board.Board, "F1", []robot.Robot{r},true)
}
