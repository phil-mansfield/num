package gorecipies

import (
	"math"
)

const (
	ConvergenceIters = 100
	ConvergenceEpsilon = 3.0e-7
)

func CloseEnough(val, diff float64) bool {
	return math.Abs(diff) <= math.Abs(val * ConvergenceEpsilon)
}
