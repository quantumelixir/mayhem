package main

import (
	"fmt"
)

// * Main
func main() {

	var r Robot

	// making a 12 x 16 board
	board := NewBoard(12, 16, Blue)
	board[0][0] = Red
	board[4][11] = Green
	board[11][15] = Red
	board[5][8] = Green
	board[2][2] = Red
	board[2][5] = Red

	// program the robot
	r.Init(11, 0, "U")
	r.DeclareFunctionList([]string{"F1", "F2"})
	r.DefineFunction("F1", []string{":F2", ":R", "Blue:F1"})
	r.DefineFunction("F2", []string{":F", "Red:R", "Green:L", "Blue:F2", ":F"})

	// check the function specification
	fmt.Println(r.FunctionList[r.FunctionIndex("F1")])
	fmt.Println(r.FunctionList[r.FunctionIndex("F2")])

	// run!
	Run(board, "F1", []*Robot{&r})

	// print the final destatination
	fmt.Println(r.X, r.Y, r.D)
}
