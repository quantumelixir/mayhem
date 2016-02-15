package main

import (
	"fmt"

	"github.com/quantumelixir/mayhem/board"
	"github.com/quantumelixir/mayhem/robot"
	"github.com/quantumelixir/mayhem/types"
)

// * Main
func main() {
	var r robot.Robot

	// making a 12 x 16 board
	board := board.NewBoard(12, 16, types.Blue)
	board[0][0] = types.Red
	board[4][11] = types.Green
	board[11][15] = types.Red
	board[5][8] = types.Green
	board[2][2] = types.Red
	board[2][5] = types.Red

	// program the robot
	r.Init(11, 0, "U")
	r.DeclareFunctionList([]string{"F1", "F2"})
	r.DefineFunction("F1", []string{":F2", ":R", "Blue:F1"})
	r.DefineFunction("F2", []string{":F", "Red:R", "Green:L", "Blue:F2", ":F"})

	// check the function specification
	fmt.Println(r.FunctionMap["F1"])
	fmt.Println(r.FunctionMap["F2"])

	// run!
	Run(board, "F1", []robot.Robot{r}, true)
}
