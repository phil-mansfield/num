package intr

import (
	"fmt"
)

// Spline represents a cubic interpolating spline curve. The curve passes
// through a given set of points and allows for the evaluation of integrals,
// derivatives, and integrals across the entire x-domain of the original data
// set.
//
// Example:
//
//     // Create spline.
//     xs, ys := getMyData()
//     sp := NewSpline(xs, ys)
//
//     // Evaluate.
//     dx := (sp.LowerBound - sp.UpperBound()) / 100.0
//     for x := sp.LowerBound(); x < fp.UpperBound; x += dx {
//          y := sp.Eval(x)
//          // Do things with y.
//     }
//
//     // Second-order derivative.
//     val := sp.Deriv(sp.LowerBound(), 2)
//
//     // Integrate.
//     area := sp.Int(sp.LowerBound(), sp.UpperBound())
type Spline struct {
	xs, dy2s []float64
	coeffs []coeff
	dy2Low, dy2High float64

	cache int
	
	// Optional params:
	reuse bool
	copyInput, unif, strict, accelInt bool
	dx float64

	// intSum[i] = \int_xs[0]^xs[i] dx S(x)
	intSum []float64
}

type coeff struct {
	a, b, c, d float64
}

// NewSpline creates a new interpolating spline corresponding to the given
// points. If the x values are not in strictly increasing order, the function
// will panic.
//
// Additional customization options can be provided as variadic arguments.
// Example:
//
//     xs, ys := getMyData()
//     sp := NewSpline(
//         xs, ys, SplineBounds(-7, +1), AccelInt(false),
//     )
func NewSpline(xs, ys []float64, opts ...SplineOption) *Spline {
	if len(xs) != len(ys) {
		panic(fmt.Sprintf("len(xs) = %d, but len(ys) = %d", len(xs), len(ys)))
	} else if len(xs) <= 1 {
		panic(fmt.Sprintf("len(xs) <= 1"))
	} else if !isIncr(xs) {
		panic("xs is not strictly increasing.")
	}

	// Setting params
	s := &Spline{
		copyInput: true, strict: true, unif: false,
		accelInt: true,
	}
	for _, opt := range opts { opt(s) }
	if s.unif { s.dx = (xs[len(xs) - 1] - xs[0]) / float64(len(xs) - 1) }

	// Allocations
	n := len(xs)
	if s.copyInput {
		s.xs = make([]float64, n)
		copy(s.xs, xs)
	} else {
		s.xs = xs
	}
	s.dy2s = make([]float64, n)
	s.coeffs = make([]coeff, n - 1)
	if s.accelInt {
		s.intSum = make([]float64, n)
	}

	s.calcY2s(ys)
	s.calcCoeffs(ys)
	if s.accelInt { s.precomputeInt() }
	return s
}

// Reuse reuses Spline for a new set of data without making any more permanent
// heap allocations.
//
// SplineOptions which change the the underlying allocation scheme of the
// Spline cannot be used:
//
//     AccelInt
//     CopyInput
//
// All other options reset to their default values if not used.
func (s *Spline) Reuse(xs, ys []float64, opts ...SplineOption) {
	if len(xs) != len(ys) {
		panic(fmt.Sprintf("len(xs) = %d, but len(ys) = %d", len(xs), len(ys)))
	} else if len(xs) != len(s.xs) {
		panic(fmt.Sprintf(
			"len(xs) = %d, but spline's internals have length %d.",
			len(xs), len(s.xs),
		))
	} else if !isIncr(xs) {
		panic("xs is not sctrictly increasing.")
	}


	s.reuse = true
	s.strict = true
	s.unif = false
	for _, opt := range opts { opt(s) }
	panic("NYI")

	// Clear buffers.
	if s.copyInput {
		copy(s.xs, xs)
	} else {
		s.xs = xs
	}
	for i := range s.dy2s { s.dy2s[i] = 0 }
	for i := range s.coeffs { s.coeffs[i] = coeff{} }
	if s.accelInt {
		for i := range s.intSum { s.intSum[i] = 0 }
	}

	s.calcY2s(ys)
	s.calcCoeffs(ys)
	if s.accelInt { s.precomputeInt() }
}

func isIncr(xs []float64) bool {
	for i := 1; i < len(xs); i++ {
		if xs[i] <= xs[i + 1] { return false }
	}
	return true
}

func (s *Spline) calcY2s(ys []float64) {
	panic("NYI")
}

func (s *Spline) calcCoeffs(ys []float64) {
	panic("NYI")
}

func (s *Spline) precomputeInt() {
	s.intSum[0] = 0
	for i := 1; i < len(s.intSum); i++ {
		s.intSum[i]  = s.intSum[i - 1] +
			intTerm(&s.coeffs[i - 1], s.xs[i - 1], s.xs[i])
	}
}

// LowerBound returns the lowest value which is within range for the spline.
// Unless the StrictRange option is set to false, calls to Eval, Deriv, and
// Int which are lower than this value will panic.
func (s *Spline) LowerBound() float64 { return s.xs[0] }

// UpperrBound returns the highest value which is within range for the spline.
// Unless the StrictRange option is set to false, calls to Eval, Deriv, and
// Int which are higher than this value will panic.
func (s *Spline) UpperBound() float64 { return s.xs[len(s.xs) - 1] }

// Intervals returns the number of intervals that the spline is divided into.
func (s *Spline) Intervals() int { return len(s.xs) - 1 }

// Coeffs returns the spline coefficients of the specified interval.
func (s *Spline) Coeffs(i int) (a, b, c, d float64) {
	if i < 0 || i > len(s.xs) - 2 {
		panic(fmt.Sprintf(
			"Index %d out of range for spline with %d intervals.",
			i, len(s.xs) - 1))
	}
	term := s.coeffs[i]
	return term.a, term.b, term.c, term.d
}

// Range returns the lower and upper bounds of the specified interval.
func (s *Spline) Range(i int) (low, high float64) {
	if i < 0 || i >= len(s.xs) - 2 {
		panic(fmt.Sprintf(
			"Index %d out of range for spline with %d intervals.",
			i, len(s.xs) - 1))
	}

	return s.xs[i], s.xs[i+1]
}

// Eval returns the value of the spline at the given point.
func (s *Spline) Eval(x float64) float64 {
	i := s.bsearch(x)
	dx := x - s.xs[i]
	a, b, c, d := s.coeffs[i].a, s.coeffs[i].b, s.coeffs[i].c, s.coeffs[i].d
	return a*dx*dx*dx + b*dx*dx + c*dx + d
}

// EvalAll returns the value of the spline at all the given points.
//
// If any output arrays are given, the output is written to those arrays with
// no allocation.
func (s *Spline) EvalAll(xs []float64, out ...[]float64) []float64 {
	var evalOut []float64
	if len(out) == 0 {
		evalOut = make([]float64, len(xs))
	} else {
		for i := range out {
			if len(out[i]) != len(xs) {
				panic(fmt.Sprintf("len(xs) = %d, but len(out[%d]) = %d",
					len(xs), i, len(out[i])))
			}
		}
		evalOut = out[0]
	}

	for i := range evalOut { evalOut[i] = s.Eval(xs[i]) }

	if len(out) > 1 {
		for i := range out[1:] { copy(out[i], evalOut) }
	}
	return evalOut
}

// Deriv calculates the derivative of the spline at the given point to the 
// specified order.
func (s *Spline) Deriv(x float64, order int) float64 {
	i := s.bsearch(x)
	dx := x - s.xs[i]
	a, b, c, d := s.coeffs[i].a, s.coeffs[i].b, s.coeffs[i].c, s.coeffs[i].d
	switch order {
	case 0:
		return a*dx*dx*dx + b*dx*dx + c*dx + d
	case 1:
		return 3*a*dx*dx + 2*b*dx + c
	case 2:
		return 6*a*dx + 2*b
	case 3:
		return 6*a
	default:
		return 0
	}
}

// DerivAll returns the derivative of the spline at all the given points
// calculated to the specified order.
//
// If any output arrays are given, the output is written to those arrays with
// no allocation.
func (s *Spline) DerivAll(xs []float64, order int, out ...[]float64) []float64 {
	var derivOut []float64
	if len(out) == 0 {
		derivOut = make([]float64, len(xs))
	} else {
		for i := range out {
			if len(out[i]) != len(xs) {
				panic(fmt.Sprintf("len(xs) = %d, but len(out[%d]) = %d",
					len(xs), i, len(out[i])))
			}
		}
		derivOut = out[0]
	}

	for i := range derivOut { derivOut[i] = s.Deriv(xs[i], order) }

	if len(out) > 1 {
		for i := range out[1:] { copy(out[i], derivOut) }
	}
	return derivOut
}

// Int calculates the integral of the spline across the given interval.
func (s *Spline) Int(low, high float64) float64 {
	iLow, iHigh := s.bsearch(low), s.bsearch(high)
	if iLow == iHigh {
		return intTerm(&s.coeffs[iLow], low, high)
	}

	sum := intTerm(&s.coeffs[iLow], low, s.xs[iLow+1]) +
		intTerm(&s.coeffs[iHigh], s.xs[iHigh], high)

	if s.accelInt {
		sum += s.intSum[iHigh - 1] - s.intSum[iLow]
	} else {
		for i := iLow + 1; i < iHigh; i++ {
			sum += intTerm(&s.coeffs[i], s.xs[i], s.xs[i+1])
		}
	}

	return sum
}

func intTerm(coeff *coeff, lo, hi float64) float64 {
	a, b, c, d := coeff.a, coeff.b, coeff.c, coeff.d
	dx := hi - lo
	return a*dx*dx*dx*dx/4 + b*dx*dx*dx/3 + c*dx*dx/2 + d*dx
}

// IntAll returns the integral of the spline across all the given intervals.
//
// If any output arrays are given, the output is written to those arrays with
// no allocation.
func (s *Spline) IntAll(lows, highs []float64, out ...[]float64) []float64 {
	if len(lows) != len(highs) {
		panic(fmt.Sprintf("len(lows) = %d, but len(highs) = %d",
			len(lows), len(highs)))
	}

	var intOut []float64
	if len(out) == 0 {
		intOut = make([]float64, len(lows))
	} else {
		for i := range out {
			if len(out[i]) != len(lows) {
				panic(fmt.Sprintf("len(lows) = %d, but len(out[%d]) = %d",
					len(lows), i, len(out[i])))
			}
		}
		intOut = out[0]
	}

	for i := range intOut { intOut[i] = s.Int(lows[i], highs[i]) }

	if len(out) > 1 {
		for i := range out[1:] { copy(out[i], intOut) }
	}
	return intOut
}

type splineOption func(*Spline)

// A cached binary search.
func (s *Spline) bsearch(x float64) int {
	// TODO: benchmark whether this should go before or after the cache check.
	if s.unif {
		if s.strict {
			if s.xs[0] > x || s.xs[len(s.xs) - 1] < x {
				panic(fmt.Sprintf(
					"%g is out of bounds for Spline with bounds [%g, %g]",
					x, s.LowerBound(), s.UpperBound(),
				))
			}
		}
		return int((x - s.xs[0]) / s.dx)
	}

	// Check the bsearch cache.
	// TODO: benchmark whether checking the adjacent intervals is worth it.
	// In the average case, the user is just iterating across a range, so
	// we could actually avoid all but the first bianry searches.
	var low, high int
	n := s.Intervals()
	if s.xs[s.cache] <= x {
		if s.xs[s.cache + 1] >= x { return s.cache }
		low, high = s.cache + 1, n
	} else {
		low, high = 0, s.cache
	}

	for high - low > 1 {
		mid := (low + high) / 2
		if x >= s.xs[mid] {
			low = mid
		} else {
			high = mid
		}
	}

	if s.strict {
		if s.xs[low] > x || s.xs[high] < x {
			panic(fmt.Sprintf(
				"%g is out of bounds for Spline with bounds [%g, %g]",
				x, s.LowerBound(), s.UpperBound(),
			))
		}
	}

	s.cache = low
	return low
}

// SplineOptions are passed to NewSpline as variadic arguments to customize its
// behavior. 
type SplineOption splineOption

// StrictRange sets the behavior of the spline outside the range of supplied
// points. If set to true, calls to Eval, EvalAll, Deriv, DerivAll, Int, and
// IntAll outside this range will panic. Otherwise the spline coefficients of
// the outermost intervals are used.
func StrictRange(strictRange bool) SplineOption {
	return func(s *Spline) { s.strict = strictRange }
}

// SplineBounds sets the boundary conditions for the second derivatives of
// the spline. If not called, both sides of the spline will use "natural"
// boundary conditions, setting the second derivatives to 0.
func SplineBounds(lower, upper float64) SplineOption {
	return func(s *Spline) { s.dy2Low, s.dy2High = lower, upper	}
}

// AccelInt sets whether or not to perform certain optimizations on integral
// calls at the expense of increased memory consumption and Spline
// initialization time. With this option set, calls to Int are O(1). Without
// this option set, calls to Int that span N intervals take O(N) time. By
// default, it is set to true.
func AccelInt(accelInt bool) SplineOption {
	return func(s *Spline) {
		if s.reuse {
			panic("Cannot set AccelInt option for calls to Spline.Reuse().")
		}
		s.accelInt = accelInt
	}
}

// Unif tells the spline that the input x values are uniformly spaced. If set
// to true, random access time is reduced.
func Unif(unif bool) SplineOption {
	return func(s *Spline) { s.unif = unif }
}

// CopyInput sets whether or not to explicitly allocate and copy the input
// slices when creating the spline. Only set this to true if memory consumption
// is a concern and the input slices are not modified over the lifetime of the
// spline.
func CopyInput(copyInput bool) SplineOption {
	return func(s *Spline) {
		if s.reuse {
			panic("Cannot set CopyInput option for calls to Spline.Reuse().")
		}
		s.copyInput = copyInput
	}
}
