package num

import (
	"math"
)

func CloseEnough(val, diff float64) bool {
	return math.Abs(diff) <= math.Abs(val * ConvergenceEpsilon)
}

// This is somewhat slower due to the extra Abs calls, but has a nicer interface.
func AlmostEqual(val1, val2 float64) bool {
	if math.Abs(val1) > math.Abs(val2) {
		return CloseEnough(val1, val1 - val2)
	} else {
		return CloseEnough(val2, val2 - val1)
	}
}
