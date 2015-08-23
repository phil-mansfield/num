package intr

type Spline struct { }

func NewSpline(xs, ys []float64, opts ...SplineOption) *Spline {
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

type splineBoundaryCondition func(*Spline)
type SplineBoundaryCondition splineBoundaryCondition

// Test Comment 1.
var natural SplineBoundaryCondition = Natural
// Test Comment 2
var finiteDiff SplineBoundaryCondition = FiniteDiff

func Natural(s *Spline) {
	panic("NYI")
}

func FiniteDiff(s *Spline){
	panic("NYI")
}
