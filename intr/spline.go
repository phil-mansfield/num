package intr


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

	findCache int

	incr bool
	
	// Optional params:
	copyInput, unif, strict, accelInt bool
	dx float64

	intSum []float64
}

type coeff struct {
	a, b, c, d float64
}

// NewSpline creates a new interpolating spline corresponding to the given
// points. If these points are not sorted in either ascending or descending
// order, the function will panic.
//
// Additional customization options can be provided as variadic arguments.
// Example:
//
//     xs, ys := getMyData()
//     sp := NewSpline(
//         xs, ys, SplineBounds(Deriv(-7), Deriv2(+1)), AccelInt(false),
//     )
func NewSpline(xs, ys []float64, opts ...SplineOption) *Spline {
	panic("NYI")
}

// Reuse reuses Spline for a new set of data without doing any additional
// allocations.
//
// SplineOptions which change the the underlying allocation scheme of the Spline
// cannot be used:
//
//     AccelInt
//     CopyInput
func Reuse(xs, ys []float64, opts ...SplineOption) {
	panic("NYI")
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
	c := s.coeffs[i]
	return c.a, c.b, c.c, c.d
}

// Range returns the lower and upper bounds of the specified interval.
func (s *Spline) Range(i int) (low, high float64) {
	if i < 0 || i >= len(s.xs) - 2 {
		panic(fmt.Sprintf(
			"Index %d out of range for spline with %d intervals.",
			i, len(s.xs) - 1))
	}

	return 
}

// Eval returns the value of the spline at the given point.
func (s *Spline) Eval(x float64) float64 {
	panic("NYI")
}

// EvalAll returns the value of the spline at all the given points. Points must
// be given in ascending order.
//
// If any output arrays are given, the output is written to those arrays with
// no allocation.
func (s *Spline) EvalAll(xs []float64, out ...[]float64) []float64 {
	panic("NYI")
}

// Deriv calculates the derivative of the spline at the given point to the 
// specified order.
func (s *Spline) Deriv(x float64, order int) float64 {
	panic("NYI")
}

// DerivAll returns the derivative of the spline at all the given points calculated to the specified order. Points must be given in ascending order.
//
// If any output arrays are given, the output is written to those arrays with
// no allocation.
func (s *Spline) DerivAll(xs []float64, order int, out ...[]float64) []float64 {
	panic("NYI")
}

// Int calculates the integral of the spline across the given interval.
func (s *Spline) Int(low, high float64) float64 {
	panic("NYI")
}

// IntAll returns the integral of the spline across all the given intervals.
// Lower bounds must be given in ascending order.
//
// If any output arrays are given, the output is written to those arrays with
// no allocation.
func (s *Spline) IntAll(lows, high []float64, out ...[]float64) []float64 {
	panic("NYI")
}

type splineOption func(*Spline)

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

// AccelInt sets whether or not to aggressively optimizae integral calls at
// the expense of slightly increased memory consumption. By default it is set
// to true.
func AccelInt(accelInt bool) SplineOption {
	return func(s *Spline) { s.accelInt = accelInt }
}

// Unif tells the spline that the input x values are uniformly spaces with
// distance dx.
func Unif(dx float64) SplineOption {
	return func(s *Spline) { s.unif, s.dx = true, dx }
}

// CopyInput sets whether or not to explicitly allocate and copy the input
// slices when creating the spline. Only set this to true if memory consumption
// is a concern and the input slices are not modified over the lifetime of the
// spline.
func CopyInput(copyInput bool) SplineOption {
	return func(s *Spline) { s.copyInput = copyInput }
}
