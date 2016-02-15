package main

import (
	"github.com/quantumelixir/mayhem/robot"
	"github.com/quantumelixir/mayhem/board"
)

const LIM = 100

// * Init
func init() {
	// initialize the board
	board.Board = make([][]robot.Color, LIM)
	for i := 0; i < LIM; i++ {
		board.Board[i] = make([]robot.Color, LIM)
	}
}

// * Main
func main() {
	var r robot.Robot

	r.Init(LIM/2, LIM/2, robot.Up)
	r.DeclareFunctionList([]string{"F1", "F2"})
	r.DefineFunction("F1", []string{"Red:Blue", "Red:F1", ":L"})
	r.DefineFunction("F2", []string{":F", ":R", ":F", ":L"})

	board.Board[r.X][r.Y] = robot.Red

	robot.Run(board.Board, "F1", []robot.Robot{r},true)
}
