package main

import "log"

// Rotation matrices for Up, Right, Down, Left
var R = [4][2][2]int{
	{{1, 0}, {0, 1}},
	{{0, -1}, {1, 0}},
	{{-1, 0}, {0, -1}},
	{{0, 1}, {-1, 0}},
}

// * The Runner
// Run all the bots starting from the main function on the defined Board
func Run(board [][]Color, main string, bots []*Robot) {

	// a single stack location
	type Location struct {
		ifun int
		ipos int
	}

	stacks := make([][]Location, len(bots))

	current := make([]Location, len(bots))

	for i, r := range bots {
		stacks[i] = []Location{}
		current[i] = Location{r.FunctionIndex(main), 0}
	}

	type Position struct {
		X, Y int
	}
	PaintMap := make(map[Position]Color)

	i := 0
	someStackNotEmpty := true

ExecutionLoop:
	for { // every len(bots) iterations corresponds to one global tick

		// stopping condition: when all bots halt
		if len(bots) < 1 || (i == len(bots) && !someStackNotEmpty) {
			break ExecutionLoop
		}

		if i == len(bots) {
			i = 0
			someStackNotEmpty = false
		}

		if len(stacks[i]) == 0 && len(bots[i].FunctionList[current[i].ifun]) == current[i].ipos {

			// nothing left to do for the i-th bot
			i++
			continue
		}

		someStackNotEmpty = true

		if len(bots[i].FunctionList[current[i].ifun]) == current[i].ipos {
			// pop the head of the i-th bot's stack
			current[i] = stacks[i][len(stacks[i])-1]
			stacks[i] = stacks[i][:len(stacks[i])-1]
		}

		f := bots[i].FunctionList[current[i].ifun]

		// read a single statement from that location
		v := f[current[i].ipos]

		if !(v.Cond == WildCard || v.Cond == board[bots[i].X][bots[i].Y]) {
			current[i].ipos++
			i++
			continue
		}

		switch v.Which {

		case Stay:
			// do nothing

		case Step:

			switch v.Step {
			case MoveForward:
				bots[i].X = bots[i].X + R[bots[i].D][0][0]*(-1) + R[bots[i].D][1][0]*(0)
				bots[i].Y = bots[i].Y + R[bots[i].D][0][1]*(-1) + R[bots[i].D][1][1]*(0)

			case TurnRight:
				bots[i].D = (bots[i].D + 1) % 4

			case TurnLeft:
				bots[i].D = (bots[i].D + 3) % 4
			}

		case Jump:

			// save the current location (optimizing away tail calls)
			if current[i].ipos + 1 < len(f) {
				stacks[i] = append(stacks[i],
					Location{current[i].ifun, current[i].ipos + 1})
			}

			// jump to the new location
			current[i] = Location{v.Jump, 0}
			continue

		case Paint:
			// don't make the updates immediately => check for race conditions
			if _, ok := PaintMap[Position{bots[i].X, bots[i].Y}]; ok {
				log.Fatal("Painting same square twice in the same tick!")
				break ExecutionLoop
			} else {
				PaintMap[Position{bots[i].X, bots[i].Y}] = v.Paint
			}

		case Return:
			current[i].ipos = len(f) - 1
		}

		// run all the painting steps at the end of a global tick
		if i+1 == len(bots) {
			for pos, color := range PaintMap {
				board[pos.X][pos.Y] = color
			}
			// zero out the map after updating
			PaintMap = make(map[Position]Color)
		}

		current[i].ipos++
		i++
	} // outer for
}
