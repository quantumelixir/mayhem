package types

// * Types: Color, Direction, Movement, Action
// ** Color
type Color int
const WildCard, Red, Green, Blue Color = 0, 1, 2, 3
var ColorMap = map[string]Color{
	"": WildCard,
	"Red": Red,
	"Green": Green,
	"Blue": Blue,
}

// ** Direction
type Direction int
const Up, Right, Down, Left Direction = 0, 1, 2, 3
var DirectionMap = map[string]Direction{
	"Up": Up,
	"Right": Right,
	"Down": Down,
	"Left": Left,
}

// ** Movement
type Movement int
const MoveForward, TurnRight, TurnLeft Movement = 0, 1, 2
var MovementMap = map[string]Movement{
	"F": MoveForward,
	"R": TurnRight,
	"L": TurnLeft,
}

// ** Action
type Action int
const Stay, Step, Jump, Paint Action = 0, 1, 2, 3

// * Statements and Functions
type Function []Statement

type Statement struct {

	Cond Color

	Which Action
	Step Movement
	Jump *Function
	Paint Color
}
