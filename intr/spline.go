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
//     
//
//     // Integrate.
//     area := sp.Int(sp.LowerBound(), sp.UpperBound())
type Spline struct { }

type coeff struct { }

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
func (s *Spline) LowerBound() float64 {
	panic("NYI")
}

// UpperrBound returns the highest value which is within range for the spline.
// Unless the StrictRange option is set to false, calls to Eval, Deriv, and
// Int which are higher than this value will panic.
func (s *Spline) UpperBound() float64 {
	panic("NYI")
}

// Intervals returns the number of intervals that the spline is divided into.
func (s *Spline) Intervals() int {
	panic("NYI")
}

// Coeffs returns the spline coefficients of the specified interval.
func (s *Spline) Coeffs(i int) (a, b, c, d float64) {
	panic("NYI")
}

// Range returns the lower and upper bounds of the specified interval.
func (s *Spline) Range(i int) (low, high float64) {
	panic("NYI")
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
	panic("NYI")
}

// SplineBounds sets the boundary conditions used by a spline. If not called,
// both sides of the spline will use the Natural flag.
func SplineBounds(lower, upper SplineBoundaryCondition) SplineOption {
	panic("NYI")
}

// AccelInt sets whether or not to aggressively optimizae integral calls at
// the expense of slightly increased memory consumption. By default it is set
// to true.
func AccelInt(flag bool) SplineOption {
	panic("NYI")
}

// CopyInput sets whether or not to explicitly allocate and copy the input
// slices when creating the spline. Only set this to true if memory consumption
// is a concern and the input slices are not modified over the lifetime of the
// spline.
func CopyInput(flag bool) SplineOption {
	panic("NYI")
}

type splineBoundaryCondition func(*Spline)
type SplineBoundaryCondition splineBoundaryCondition

// Natural forces the spline to use natural boundary condtions (the second
// derivative is set to 0).
var Natural SplineBoundaryCondition = natural

func natural(s *Spline) {
	panic("NYI")
}

func finiteDiff(s *Spline){
	panic("NYI")
}

// Deriv2 forces the spline to use boundary conditions such that d^2y/dx^2 is
// equal to the specified values.
func Deriv2(val float64) SplineBoundaryCondition {
	panic("NYI")
}
