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

func SmoothBounds(lower, upper SmoothBoundaryCondition) SmoothOption {
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

type smoothBoundaryCondition func(lower, uppper bool, p *smoothParams)
type SmoothBoundaryCondition smoothBoundaryCondition

var Mirror SmoothBoundaryCondition = mirror
var Extend SmoothBoundaryCondition = extend
var Periodic SmoothBoundaryCondition = periodic

func Constant(c float64) SmoothBoundaryCondition {
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
