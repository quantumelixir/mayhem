package robot

import (
	"log"
	"strings"

	"github.com/quantumelixir/mayhem/types"
)

// * Robot
type Robot struct {
	X, Y        int
	D           types.Direction
	FunctionMap map[string]*types.Function
}

// ** Init
func (r *Robot) Init(X int, Y int, direction string) {
	r.X, r.Y = X, Y
	r.D = types.DirectionMap[direction]
	r.FunctionMap = make(map[string]*types.Function)
}

// ** DeclareFunctionList
func (r *Robot) DeclareFunctionList(list []string) {
	for _, s := range list {
		if _, ok := r.FunctionMap[s]; !ok {
			r.FunctionMap[s] = &types.Function{}
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

		var t types.Statement

		if v, ok := types.ColorMap[split[0]]; ok {
			t.Cond = v
		} else {
			t.Cond = types.WildCard
		}

		// avoid function names that are colors to prevent
		// confusion between the last two cases

		if v, ok := types.MovementMap[split[1]]; ok {
			t.Which = types.Step
			t.Step = v
		} else if v, ok := r.FunctionMap[split[1]]; ok {
			t.Which = types.Jump
			t.Jump = v
		} else if v, ok := types.ColorMap[split[1]]; ok {
			t.Which = types.Paint
			t.Paint = v
		}

		*r.FunctionMap[name] = append(*r.FunctionMap[name], t)
	}
}
