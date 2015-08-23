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
var Natural SplineBoundaryCondition = natural
// Test Comment 2
var FiniteDiff SplineBoundaryCondition = finiteDiff

func natural(s *Spline) {
	panic("NYI")
}

func finiteDiff(s *Spline){
	panic("NYI")
}
