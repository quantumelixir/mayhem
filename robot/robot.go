package robot

import (
	"fmt"
	"strings"
	"log"
)

// * Types: Color, Direction, Step, Selector
// ** Color
type Color int
const (
	None =  iota
	Red
	Green
	Blue
)
var ColorMap = map[string]Color{
	"": None,
	"Red": Red,
	"Green": Green,
	"Blue": Blue,
}

// ** Direction
type Direction int
const (
	Up = iota
	Right
	Down
	Left
)
var DirectionMap = map[string]Direction{
	"Up": Up,
	"Right": Right,
	"Down": Down,
	"Left": Left,
}

// ** Step
type Step int
const (
	MoveForward = iota
	TurnRight
	TurnLeft
)
var StepMap = map[string]Step{
	"F": MoveForward,
	"R": TurnRight,
	"L": TurnLeft,
}

// ** Selector
type Selector int
const (
	Stay = iota
	Move
	Jump
	Paint
)
// * Statements and Functions
type Function []Statement

type Statement struct {

	cond Color

	which Selector
	step Step
	jump *Function
	paint Color

}

// * Robot
type Robot struct {
	X, Y int
	D Direction
	FunctionMap map[string]*Function
}

// ** Init
func (r *Robot) Init(X int, Y int, D Direction) {
	r.X, r.Y = X, Y
	r.D = D
	r.FunctionMap = make(map[string]*Function)
}

// ** DeclareFunctionList
func (r *Robot) DeclareFunctionList(list []string) {
	for _, s := range list {
		if _, ok := r.FunctionMap[s]; !ok {
			r.FunctionMap[s] = &Function{}
		}
	}
}

// ** DefineFunction

// Each element of parse is a string of the form "X:Y" where X is a
// Color and Y can be one of the following:
// - a Step (F/L/R),
// - the name of a Function, or
// - a Color (in which case it is interpreted as a painting move).

func (r *Robot) DefineFunction(name string, parse []string) {

	if _, ok := r.FunctionMap[name]; !ok {
		log.Fatal("Trying to define undeclared function", name)
	}

	// zero out the previous Function definition
	*r.FunctionMap[name] = (*r.FunctionMap[name])[:0]

	for _, s := range parse {
		split := strings.Split(s, ":")

		var t Statement

		if v, ok := ColorMap[split[0]]; ok {
			t.cond = v
		} else {
			t.cond = None
		}

		// avoid function names that are colors to prevent
		// confusion between the last two cases

		if v, ok := StepMap[split[1]]; ok {
			t.which = Move
			t.step = v
		} else if v, ok := r.FunctionMap[split[1]]; ok {
			t.which = Jump
			t.jump = v
		} else if v, ok := ColorMap[split[1]]; ok {
			t.which = Paint
			t.paint = v
		}

		*r.FunctionMap[name] = append(*r.FunctionMap[name], t)
	}
}

// * The Runner

// Run all the bots starting from the main function on the defined Board

func Run(board [][]Color, main string, bots []Robot, say bool) {

	// Rotation matrices for Up, Right, Down, Left
	var R = [4][2][2]int{
		[2][2]int{[2]int{1, 0}, [2]int{0, 1}},
		[2][2]int{[2]int{0, -1}, [2]int{1, 0}},
		[2][2]int{[2]int{-1, 0}, [2]int{0, -1}},
		[2][2]int{[2]int{0, 1}, [2]int{-1, 0}},
	}

	// a single stack location
	type Location struct {
		f *Function
		position int
	}

	stacks := make([][]Location, len(bots))

	for i, r := range bots {
		stacks[i] = []Location{{r.FunctionMap[main], 0}}
	}

	type Position struct {
		X, Y int
	}
	PaintMap := make(map[Position]Color)

	// holds the currently processing stack item for each bot
	current := make([]Location, len(bots))

	i := 0
	someStackNotEmpty := true
ExecutionLoop:
	for { // every len(bots) iterations corresponds to one global tick

		// stopping condition: when all bots halt
		if i == len(bots) && !someStackNotEmpty {
			break
		}

		if i == len(bots) {

			i = 0
			someStackNotEmpty = false
			
			if say {
				fmt.Println("Tick!")
			}
		}

		if len(stacks[i]) == 0 && len(*(current[i].f)) == current[i].position {

			// nothing left to do for the i-th bot

			i++
			continue
		}
		
		someStackNotEmpty = true

		if current[i].f == nil {
			// pop the head of the i-th bot's stack
			current[i] = stacks[i][len(stacks[i]) - 1]
			stacks[i] = stacks[i][:len(stacks[i]) - 1]
		}

		// read a single statement from that location
		v := (*(current[i].f))[current[i].position]

		r := bots[i]

		if !(v.cond == None || v.cond == board[r.X][r.Y]) {
			if say {
				fmt.Println(i, "|", "Unsatisfied condition")
			}

			current[i].position++
			i++
			continue
		}

		switch v.which {

		case Stay:
			// do nothing
			if say {
				fmt.Println(i, "|", "Staying in place")
			}

		case Move:
			switch v.step {

			case MoveForward:
				r.X = r.X + R[r.D][0][0] * (-1) + R[r.D][1][0] * (0)
				r.Y = r.Y + R[r.D][0][1] * (-1) + R[r.D][1][1] * (0)
				if say {
					fmt.Printf("%d | Moving forward to (%d, %d)\n", i, r.X, r.Y)
				}

			case TurnRight:
				r.D = (r.D + 1) % 4
				if say {
					var whichway string
					for key, value := range DirectionMap {
						if value == r.D {
							whichway = key
							break
						}
					}
					fmt.Println(i, "|", "Facing", whichway)
				}

			case TurnLeft:
				r.D = (r.D + 3) % 4
				if say {
					var whichway string
					for key, value := range DirectionMap {
						if value == r.D {
							whichway = key
							break
						}
					}
					fmt.Println(i, "|", "Facing", whichway)
				}
			}

		case Jump:
			// save the current location (optimizing away tail calls)
			if current[i].position + 1 < len(*(current[i].f)) {
				stacks[i] = append(stacks[i],
					Location{current[i].f, current[i].position + 1})
			}
			// jump to the new location
			current[i] = Location{v.jump, 0}

			if say {
				var whereto string
				for key, value := range r.FunctionMap {
					if value == v.jump {
						whereto = key
						break
					}
				}
				fmt.Println(i, "|", "Jumping to", whereto)
			}

		case Paint:
			// don't make the updates immediately => check for race conditions
			if _, ok := PaintMap[Position{r.X, r.Y}]; ok {
				fmt.Println("ERROR: Painting same square twice in the same tick!")
				break ExecutionLoop
			} else {
				PaintMap[Position{r.X, r.Y}] = v.paint
			}

			if say {
				var color string
				for key, value := range ColorMap {
					if value == v.paint {
						color = key
						break
					}
				}
				fmt.Println(i, "|", "Painting", color)
			}

		}

		// run all the painting steps at the end of a global tick
		if i + 1 == len(bots) {
			for pos, color := range PaintMap {
				board[pos.X][pos.Y] = color
			}
			// zero out the map after updating
			PaintMap = make(map[Position]Color)
		}

		current[i].position++
		i++
	} // outer for
}
