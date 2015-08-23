package intr

// Smooth smooths a uniformly-spaced sequence of data points by applying some
// kernel to a moving window of points.
//
// There are numerious customization options, but the average user will usually
// only need to specify a kernel type and a kernel width.
//
// Example:
// 
//     noisyData := myFunction()
//     smoothData := Smooth(vals, SavGol(11))
//
// If you don't know which filter to use, SavGol is recommended as a default.
//
// Additional options can be provided as variadic arguments. For example:
//
//     Smooth(noisyData, SavGol(9), Out(smoothData))
//
// Supported options are:
//
//     Out(out)
//     SmoothBounds(lower, upper)
func Smooth(vals []float64, k *Kernel, opts ...SmoothOption) []float64 {
	panic("NYI")
}

type smoothParams struct { }
func (p *smoothParams) load(opts []SmoothOption) {
	for _, opt := range opts { opt(p) }
}

type smoothOption func(*smoothParams)

// SmoothOptions can be used to provide additional information to the Smooth
// function which is not neccessary for the average user. They perform a similar
// function to keyword arguments.
type SmoothOption smoothOption

// SmoothBounds specifies the boundary conditions used at the upper and lower
// bounds of the data seqeunce.
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

func Median(window) *Kernel {
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
var Fit SmoothBoundaryCondition = fit

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
