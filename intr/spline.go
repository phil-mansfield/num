package intr

type Spline struct { }

type Coeff struct { }

func NewSpline(xs, ys []float64, opts ...SplineOption) *Spline {
	panic("NYI")
}

func (s *Spline) LowerBound() float64 {
	panic("NYI")
}

func (s *Spline) UpperBound() float64 {
	panic("NYI")
}

func (s *Spline) Coeffs(out ...[]Coeff) []Coeff {
	panic("NYI")
}

func (s *Spline) Eval(x float64) float64 {
	panic("NYI")
}

func (s *Spline) EvalAll(xs []float64, out ...[]float64) []float64 {
	panic("NYI")
}

func (s *Spline) Deriv(x float64, order int) float64 {
	panic("NYI")
}

func (s *Spline) DerivAll(xs []float64, order int, out ...[]float64) []float64 {
	panic("NYI")
}

func (s *Spline) Int(low, high float64) float64 {
	panic("NYI")
}

func (s *Spline) IntAll(lows, high []float64, out ...[]float64) []float64 {
	panic("NYI")
}

type splineOption func(*Spline)
type SplineOption splineOption

func StrictRange(strictRange bool) SplineOption {
	panic("NYI")
}

func SplineBounds(lower, upper SplineBoundaryCondition) SplineOption {
	panic("NYI")
}

func MemoizeIntegrals(flag bool) SplineOption {
	panic("NYI")
}

func CopyInput(flag bool) SplineOption {
	panic("NYI")
}

type splineBoundaryCondition func(*Spline)
type SplineBoundaryCondition splineBoundaryCondition

var Natural SplineBoundaryCondition = natural
var FiniteDiff SplineBoundaryCondition = finiteDiff

func natural(s *Spline) {
	panic("NYI")
}

func finiteDiff(s *Spline){
	panic("NYI")
}
