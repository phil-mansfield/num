package num

import (
	"fmt"
	"math"
)

const (
	convergenceDigits = 1e6
	convergenceLimit  = 1000
)

// This function is just wrong.
func almostEq(x1, x2 float64) bool {
	var min, max float64
	a1, a2 := math.Abs(x1), math.Abs(x2)

	if a1 < a2 {
		min, max = a1, a2
	} else {
		min, max = a2, a1
	}

	return (max - min) <= (min / convergenceDigits)
}

func scaledAlmostEq(x1, x2, scale float64) bool {
	return math.Abs(x1 - x2) <= (scale / convergenceDigits)
}

func sameSign(x, y float64) bool {
	return (x < 0 && y < 0) || (x > 0 && y > 0) || (x == 0 && y == 0)
}

// Find returns an x value such that f(x) ~ target, with min < x max. f
// must be a strictly increasing or a strictly decreasing function.
func Find(f Func1D, target, min, max float64) float64 {
	if min > max {
		panic(fmt.Sprintf("min = %g is greater than max = %g.", min, max))
	}

	mid := min + (max-min)/2.0
	funcDir := Derivative(f, max-min)(mid)

	if funcDir == 0.0 {
		panic(fmt.Sprintf("Function is flat at x = %g", mid))
	}

	fmax, fmin := f(max), f(min)

	if math.IsNaN(fmax) || math.IsNaN(fmin) ||
		math.IsInf(fmax, 0) || math.IsInf(fmin, 0) {
		panic(fmt.Sprintf("function bounds are invalid x: (%.5g, %.5g), "+
			"y: (%.5g, %.5g)", min, max, fmin, fmax))
	}

	if (fmax > target && fmin > target) || (fmax < target && fmin < target) {
		panic(fmt.Sprintf("target %.6g not in function bounds x: (%.6g, %.6g)"+
			"y: (%.6g, %.6g)", target, min, max, fmin, fmax))
	}

	startMin, startMax := min, max

	fx := f(mid)
	i := 0
	for !almostEq(min, max) {
		if sameSign(target-fx, funcDir) {
			min = mid
		} else {
			max = mid
		}

		mid = min + (max-min)/2.0
		fx = f(mid)

		i++
		if i > convergenceLimit {
			panic(fmt.Sprintf("Function failed to converge. startMin = %.6g"+
				" startMax = %.6g, min = %.6g, max = %.6g, f(startMin) = %.6g"+
				" f(startMax) = %.6g, f(min) = %.6g,"+
				" f(max) = %.6g, target = %.6g",
				startMin, startMax, min, max,
				f(startMin), f(startMax), f(min), f(max), target))
		}
	}

	return mid
}

// FindZero finds an x value at which f is roughly zero using Newton's
// method. The scale of  derivatives taken during the computation is
// given by scale.
//
// This function is roughly three times as slow as Find() for low-weight
// functions, but is faster if you have more than three calls to big
// functions in the math module.
//
// Do not call this routine on functions which have a zero derivative
// very close to either the solution or the initial guess. This function
// may not recognize convergence if the solution is x = 0. I choose not to
// manually check for this case because external funcitons are not
// garuanteed to converge or return valid outputs for arbitrary x-values.
func FindZero(f Func1D, guess, scale float64) float64 {
	firstGuess := guess
	prevAtGuess := math.Inf(0)
	prevGuess := math.Inf(0)

	atGuess := f(guess)

	i := 0
	for !almostEq(guess, prevGuess) {
		prevAtGuess = atGuess
		prevGuess = guess

		dfdx := Derivative(f, scale)
		guess = guess - atGuess/dfdx(guess)
		if math.IsNaN(guess) || math.IsInf(guess, 0) {
			panic(fmt.Sprintf("df/dx = %.5g, guess = %.5g, f(guess) = %.g",
				dfdx(prevGuess), prevGuess, f(prevGuess)))
		}

		atGuess = f(guess)

		i++
		if i > convergenceLimit {
			panic(fmt.Sprintf("Function failed to converge. guess = %.5g,"+
				" f(guess) = %.5g, prevGuess = %.5g, f(prevGuess) = %.5g,"+
				" firstGuess = %.5g, f(firstGuess) = %.5g", guess, atGuess,
				prevGuess, prevAtGuess, firstGuess, f(firstGuess)))
		}
	}

	return guess
}

// FindEqual finds an x value at which f1 and f2 are equal.  The size of
// "interesting" features in f1 and f2 is given by scale and an x-value
// in the same region as the solution is given by guess.
func FindEqual(f1, f2 Func1D, guess, scale float64) float64 {
	diff := func(x float64) float64 { return f1(x) - f2(x) }
	return FindZero(diff, guess, scale)
}

// FindEqualConst finds an x value at which f is equal to c.  The size of
// "interesting" features in f1 and f2 is given by scale and an x-value
// in the same region as the solution is given by guess.
func FindEqualConst(f Func1D, c, guess, scale float64) float64 {
	diff := func(x float64) float64 { return f(x) - c }
	return FindZero(diff, guess, scale)
}

// This is non-optimal so that the user does not need to provide
// a scale for the derivatives.
func Maximum(f Func1D, minX, maxX float64) float64 {
	if minX > maxX {
		panic(fmt.Sprintf("minX: %g is larger than maxX: %g", minX, maxX))
	}

	scale := maxX - minX
	dfdx := Derivative(f, scale)

	var midPoint float64
	for (maxX - minX) * convergenceDigits > scale {
		midPoint = (maxX + minX) / 2.0
		dfdxMid := dfdx(midPoint)

		if dfdxMid < 0 {
			maxX = midPoint
		} else if dfdxMid > 0 {
			minX = midPoint
		} else { 
			// This shouldn't really ever happen.
			return midPoint
		}
	}

	return midPoint
}
