package intr

func Smooth(vals []float64, k *Kernel, opts ...SmoothOption) []float64 {
	panic("NYI")
}

type smoothParams struct { }
func (p *smoothParams) load(opts []SmoothOption) {
	for _, opt := range opts { opt(p) }
}

type smoothOption func(*smoothParams)
type SmoothOption smoothOption

func Bounds(bc BoundaryCondition) SmoothOption {
	panic("NYI")
}

func LowerBound(bc BoundaryCondition) SmoothOption {
	panic("NYI")
}

func UpperBound(bc BoundaryCondition) SmoothOption {
	panic("NYI")
}

func Out(out []float64) SmoothOption {
	panic("NYI")
}

type Kernel struct { }

func Gaussian(width, dx float64) *Kernel {
	panic("NYI")
}

func Tophat(width, dx float64) *Kernel {
	panic("NYI")
}

func SavGol(window int, opts ...SavGolOption) *Kernel {
	panic("NYI")
}

type savGolOption func(*smoothParams)
type SavGolOption savGolOption

func Deriv(order int, dx float64) SavGolOption {
	panic("NYI")
}

func Order(order int) SavGolOption {
	panic("NYI")
}

type boundaryCondition func(lower, uppper bool, p *smoothParams)
type BoundaryCondition boundaryCondition

var (
	Mirror BoundaryCondition = mirror
	Extend BoundaryCondition = extend
	Periodic BoundaryCondition = periodic
)

func Constant(c float64) BoundaryCondition {
	panic("NYI")
}

func mirror(lower, upper bool, p *smoothParams) {
	panic("NYI")
}

func extend(lower, upper bool, p *smoothParams) {
	panic("NYI")
}

func periodic(lower, upper bool, p *smoothParams) {
	panic("NYI")
}
