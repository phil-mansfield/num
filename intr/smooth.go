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
// Currently the supported kernels are
//
//     SavGol(windowSize, opts...)
//     Gaussian(width, dx)
//     Tophat(width, dx)
//     Median(windowSize)
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

// SmoothOption can be used to provide additional information to the Smooth
// function which is not neccessary for the average user. They perform a similar
// function to keyword arguments.
type SmoothOption smoothOption

// SmoothBounds specifies the boundary conditions used at the upper and lower
// bounds for calls to Smooth().
//
// The supported boundary conditions are:
//
//     Fit
//     Extend
//     Const(c)
//     Pediodic
//     Reflect
//
// If the SmoothBounds option is not supplied to a call to Smooth(), Fit will be
// used on both ends.
func SmoothBounds(lower, upper SmoothBoundaryCondition) SmoothOption {
	panic("NYI")
}

// Out provides an output slice for Smooth() so that no allocations need to
// be done.
func Out(out []float64) SmoothOption {
	panic("NYI")
}

// Kernel 
type Kernel struct { }

// Gaussian creates a gaussian kernel with sigma = width / 2 for a seqeunce of
// data points which are separated by a distance of dx.
func Gaussian(width, dx float64) *Kernel {
	panic("NYI")
}

// Gaussian creates a tophat kernel with the given width for a seqeunce of
// data points which are separated by a distance of dx.
func Tophat(width, dx float64) *Kernel {
	panic("NYI")
}

// Median creates a "kernel" which places the median value of the moving window
// at its center.
//
// windowSize must be odd.
func Median(windowSize int) *Kernel {
	panic("NYI")
}

// SavGol creates a kernel implementing a Savitzky-Golay filter. Applying this
// kernel to noisy data is equivalent to fitting a low-order polynomial to
// all the points within the moving window. It it designed to conserver
// higher moments in the in input data.
//
// Computing this kernel can be expensive, but applying it is no more expensive
// than applying any other kernel of the same size.
//
// windowSize must be odd and must be larger than the order of the fitted
// polynomials (which is 4 by default).
//
// SavGol supports the following options:
//
//     Deriv(order, dx)
//     Order(order)
func SavGol(windowSize int, opts ...SavGolOption) *Kernel {
	panic("NYI")
}

type savGolOption func(*smoothParams)
// SavGolOption can be used to provide additional information to the SavGol
// function which is not neccessary for the average user. They perform a similar
// function to keyword arguments.
type SavGolOption savGolOption

// Deriv changes the smoothing kernel so that the slice returned by Smooth()
// is the derivative of the smoothed data. order is the order of the derivative
// and dx is the separation between points in the input slice.
//
// The order of the derivative must not be larger than the order of fitted
// polynomials. It is recommended to use a polynomial order of at least 4 if
// this option is used.
func Deriv(order int, dx float64) SavGolOption {
	panic("NYI")
}

// Order sets the order of the fitted polynomials in a Savitzky-Golay filter.
// It is also equal to the highest conserved moment of the underlying data.
//
// Order must be smaller than the window size, must be even, and must not be
// smaller than the derivative, if one is taken.
func Order(order int) SavGolOption {
	panic("NYI")
}

type smoothBoundaryCondition func(lower, uppper bool, p *smoothParams)

// SmoothBoundaryCondition specifies how the function Smooth() should handle the
// edges of the input slice.
type SmoothBoundaryCondition smoothBoundaryCondition

// Reflect sets the boundary conditions to be reflective.
var Reflect SmoothBoundaryCondition = mirror

// Extend sets the boundary conditions so that the points outside the data range
// are taken to be equal to the extremal value of the raw data.
var Extend SmoothBoundaryCondition = extend

// Periodic sets the boundary conditions to be periodic.
var Periodic SmoothBoundaryCondition = periodic

// Fit fits a polynomial to the points on the edge of the raw data and takes
// values outside of the data range to be equal to the value of the polynomial.
var Fit SmoothBoundaryCondition = fit

// Extend sets the boundary conditions so that the points outside the data range
// are taken to be equal to the given value.
func Const(c float64) SmoothBoundaryCondition {
	panic("NYI")
}

func fit(lower, upper bool, p *smoothParams){
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
