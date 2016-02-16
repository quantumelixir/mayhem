package main

// * Types: Color, Direction, Movement, Action
// ** Color
type Color byte

const WildCard, Red, Green, Blue, Invalid Color = 0, 1, 2, 3, 4

var ColorMap = map[string]Color{
	"":      WildCard,
	"Red":   Red,
	"Green": Green,
	"Blue":  Blue,
}

// ** Direction
type Direction byte

const Up, Right, Down, Left Direction = 0, 1, 2, 3

var DirectionMap = map[string]Direction{
	"Up":    Up,
	"Right": Right,
	"Down":  Down,
	"Left":  Left,
}

// ** Movement
type Movement byte

const MoveForward, TurnRight, TurnLeft Movement = 0, 1, 2

var MovementMap = map[string]Movement{
	"F": MoveForward,
	"R": TurnRight,
	"L": TurnLeft,
}

// ** Action
type Action byte

const Stay, Step, Jump, Paint, Return Action = 0, 1, 2, 3, 4

// * Statements and Functions
type Function []Statement

type Statement struct {
	Cond Color

	Which Action
	Step  Movement
	Jump  int // index of the Function to jump to
	Paint Color
}
