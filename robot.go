package main

import (
	"log"
	"strings"
)

// * Robot
type Robot struct {
	X, Y         int
	D            Direction
	FunctionList []Function
	FunctionName []string
}

// ** Init
func (r *Robot) Init(X int, Y int, direction string) {
	r.X, r.Y = X, Y
	r.D = DirectionMap[direction]
	r.FunctionList = []Function{}
	r.FunctionName = []string{}
}

// ** FunctionIndex

// Return the index where Function named "name" is found
// Returns the length of the FunctionList otherwise

func (r *Robot) FunctionIndex(name string) int {
	for i, v := range r.FunctionName {
		if v == name {
			return i
		}
	}

	return len(r.FunctionName)
}

// ** DeclareFunction

// Declare a new function or rename an existing one

func (r *Robot) DeclareFunction(name string) {

	index, length := r.FunctionIndex(name), len(r.FunctionList)

	if index != length {
		r.FunctionName[index] = name
	} else {
		r.FunctionList = append(r.FunctionList, Function{})
		r.FunctionName = append(r.FunctionName, name)
	}
}

// ** DeclareFunctionList
func (r *Robot) DeclareFunctionList(list []string) {
	for _, name := range list {
		r.DeclareFunction(name)
	}
}

// ** DefineFunction

// Each element of parse is a string of the form "X:Y" where X is a
// Color and Y can be one of the following:
// - a Step (F/L/R),
// - the name of a Function, or
// - a Color (in which case it is interpreted as a painting move), or
// - the string "Return"
func (r *Robot) DefineFunction(name string, parse []string) {

	index, length := r.FunctionIndex(name), len(r.FunctionList)
	
	if index == length {
		log.Fatal("Trying to define undeclared function", name)
	}

	// zero out the previous Function definition
	r.FunctionList[index] = Function{}

	for _, s := range parse {
		split := strings.Split(s, ":")

		var t Statement

		if v, ok := ColorMap[split[0]]; ok {
			t.Cond = v
		} else {
			t.Cond = WildCard
		}

		// avoid function names that are colors to prevent
		// confusion between the last two cases

		if v, ok := MovementMap[split[1]]; ok {
			t.Which = Step
			t.Step = v
		} else if v := r.FunctionIndex(split[1]); v != length {
			t.Which = Jump
			t.Jump = v
		} else if v, ok := ColorMap[split[1]]; ok {
			t.Which = Paint
			t.Paint = v
		} else if split[1] == "Return" {
			t.Which = Return
		}

		r.FunctionList[index] = append(r.FunctionList[index], t)
	}
}
