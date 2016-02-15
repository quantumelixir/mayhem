package main

import (
	"fmt"

	"github.com/quantumelixir/mayhem/robot"
	"github.com/quantumelixir/mayhem/types"
)

// Rotation matrices for Up, Right, Down, Left
var R = [4][2][2]int{
	{{1, 0}, {0, 1}},
	{{0, -1}, {1, 0}},
	{{-1, 0}, {0, -1}},
	{{0, 1}, {-1, 0}},
}

// * The Runner
// Run all the bots starting from the main function on the defined Board
func Run(board [][]types.Color, main string, bots []robot.Robot, say bool) {

	// a single stack location
	type Location struct {
		f        *types.Function
		position int
	}

	stacks := make([][]Location, len(bots))

	for i, r := range bots {
		stacks[i] = []Location{{r.FunctionMap[main], 0}}
	}

	type Position struct {
		X, Y int
	}
	PaintMap := make(map[Position]types.Color)

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

			// if say {
			// 	fmt.Println("Tick!")
			// }
		}

		if len(stacks[i]) == 0 &&
			(current[i].f == nil || len(*(current[i].f)) == current[i].position) {

			// nothing left to do for the i-th bot

			i++
			continue
		}

		someStackNotEmpty = true

		if current[i].f == nil || len(*(current[i].f)) == current[i].position {
			// pop the head of the i-th bot's stack
			current[i] = stacks[i][len(stacks[i])-1]
			stacks[i] = stacks[i][:len(stacks[i])-1]
		}

		// read a single statement from that location
		v := (*(current[i].f))[current[i].position]

		if !(v.Cond == types.WildCard || v.Cond == board[bots[i].X][bots[i].Y]) {
			// if say {
			// 	fmt.Println(v.Cond, types.WildCard, board[bots[i].X][bots[i].Y])
			// 	fmt.Println(i, "|", "Unsatisfied condition")
			// }

			current[i].position++
			i++
			continue
		}

		switch v.Which {

		case types.Stay:

			// do nothing
			if say {
				fmt.Println(i, "|", "Staying in place")
			}

		case types.Step:

			switch v.Step {

			case types.MoveForward:
				bots[i].X = bots[i].X + R[bots[i].D][0][0]*(-1) + R[bots[i].D][1][0]*(0)
				bots[i].Y = bots[i].Y + R[bots[i].D][0][1]*(-1) + R[bots[i].D][1][1]*(0)
				if say {
					fmt.Printf("%d | Moving Forward to (%d, %d)\n", i, bots[i].X, bots[i].Y)
				}

			case types.TurnRight:
				bots[i].D = (bots[i].D + 1) % 4
				if say {
					var whichway string
					for key, value := range types.DirectionMap {
						if value == bots[i].D {
							whichway = key
							break
						}
					}
					fmt.Println(i, "|", "Facing", whichway)
				}

			case types.TurnLeft:
				bots[i].D = (bots[i].D + 3) % 4
				if say {
					var whichway string
					for key, value := range types.DirectionMap {
						if value == bots[i].D {
							whichway = key
							break
						}
					}
					fmt.Println(i, "|", "Facing", whichway)
				}
			}

		case types.Jump:

			// save the current location (optimizing away tail calls)
			if current[i].position+1 < len(*(current[i].f)) {
				stacks[i] = append(stacks[i],
					Location{current[i].f, current[i].position + 1})
			}

			// jump to the new location
			current[i] = Location{v.Jump, 0}

			if say {
				var whereto string
				for key, value := range bots[i].FunctionMap {
					if value == v.Jump {
						whereto = key
						break
					}
				}
				fmt.Println(i, "|", "Jumping to", whereto)
			}
			continue

		case types.Paint:
			// don't make the updates immediately => check for race conditions
			if _, ok := PaintMap[Position{bots[i].X, bots[i].Y}]; ok {
				fmt.Println("ERROR: Painting same square twice in the same tick!")
				break ExecutionLoop
			} else {
				PaintMap[Position{bots[i].X, bots[i].Y}] = v.Paint
			}

			if say {
				var color string
				for key, value := range types.ColorMap {
					if value == v.Paint {
						color = key
						break
					}
				}
				fmt.Println(i, "|", "Painting", color)
			}

		case types.Return:
			current[i].position = len(*(current[i].f)) - 1
		}

		// run all the painting steps at the end of a global tick
		if i+1 == len(bots) {
			for pos, color := range PaintMap {
				board[pos.X][pos.Y] = color
			}
			// zero out the map after updating
			PaintMap = make(map[Position]types.Color)
		}

		current[i].position++
		i++
	} // outer for
}
