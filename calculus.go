package num

import (
	"math"
)

type Func1D func(float64) float64
type Func1DArray func(float64) (xs, ys []float64)

type ScaleType int

const (
	Log ScaleType = iota
	Linear
)

type DomainType int

const (
	Spherical DomainType = iota
	Flat
)

// Derivative returns a function that computes the derivative of f. scale
// is the distance at which "interesting" features of the function can be
// seen.
func Derivative(f Func1D, scale float64) Func1D {
	dx := scale / 1e4
	return func(x float64) float64 {
		return (f(x+dx) - f(x-dx)) / (2.0 * dx)
	}
}

func integrateBlock(f Func1D, center, width float64, st ScaleType) float64 {
	var stepStart, stepEnd, stepMiddle float64

	switch st {
	case Log:
		stepStart = math.Pow(10.0, center-width/2.0)
		stepEnd = math.Pow(10.0, center+width/2.0)
		stepMiddle = math.Pow(10.0, center)
	case Linear:
		stepStart = center - width/2.0
		stepEnd = center + width/2.0
		stepMiddle = center
	}

	return f(stepMiddle) * (stepEnd - stepStart)
}

// Integral returns a function that computest the integral of f starting
// at xStart. scale determines the distance at whcih "interesting"
// features of the function can be seen.
//
// st determines whether the integration steps are distributed in
// log-space and dt determines whether the integral is linear or spherical.
func Integral(f Func1D, xStart, scale float64, st ScaleType, dt DomainType) Func1D {

	var g Func1D
	switch dt {
	case Spherical:
		g = func(x float64) float64 { return f(x) * x * x * 4.0 * math.Pi }
	case Flat:
		g = func(x float64) float64 { return f(x) }
	}

	if st == Log {
		xStart = math.Log10(xStart)
	}
	dx := scale / 100.0

	return func(xEnd float64) float64 {
		var width, signedDx float64
		if st == Log {
			xEnd = math.Log10(xEnd)
		}

		if xStart == xEnd {
			return 0
		} else if xStart < xEnd {
			width = xEnd - xStart
			signedDx = dx
		} else {
			width = xStart - xEnd
			signedDx = -dx
		}

		fullSteps := int(math.Floor(width / dx))

		x := xStart + (signedDx / 2.0)
		sum := 0.0
		for i := 0; i < fullSteps; i++ {
			sum += integrateBlock(g, x, signedDx, st)
			x += signedDx
		}

		x -= signedDx / 2.0
		sum += integrateBlock(g, x+(xEnd-x)/2.0, xEnd-x, st)

		return sum
	}
}

// IntegralArray is equivelent to Integral, but returns an array of the
// the x values of the intermediate steps and the value of the integral
// at those points.
func IntegralArray(f Func1D, xStart, scale float64, st ScaleType, dt DomainType) Func1DArray {
	var g Func1D
	switch dt {
	case Spherical:
		g = func(x float64) float64 { return f(x) * x * x * 4.0 * math.Pi }
	case Flat:
		g = func(x float64) float64 { return f(x) }
	}

	if st == Log {
		xStart = math.Log10(xStart)
	}
	dx := scale / 100.0

	return func(xEnd float64) (xs, ys []float64) {
		var width, signedDx float64
		if st == Log {
			xEnd = math.Log10(xEnd)
		}

		if xStart == xEnd {
			return []float64{0.0}, []float64{0.0}
		} else if xStart < xEnd {
			width = xEnd - xStart
			signedDx = dx
		} else {
			width = xStart - xEnd
			signedDx = -dx
		}

		fullSteps := int(math.Floor(width / dx))

		xs = make([]float64, fullSteps+2)
		ys = make([]float64, fullSteps+2)

		xs[0] = xStart
		ys[0] = 0.0

		xs[1] = xStart + (signedDx / 2.0)
		ys[1] = integrateBlock(g, xs[1], signedDx, st)

		for i := 1; i < fullSteps; i++ {
			xs[i+1] = xs[i] + signedDx
			ys[i+1] += ys[i] + integrateBlock(g, xs[i+1], signedDx, st)
		}

		x := xs[fullSteps]
		xs[fullSteps+1] = xEnd
		ys[fullSteps+1] = (ys[fullSteps] +
			integrateBlock(g, x+(xEnd-x)/2.0, xEnd-x, st))

		return xs, ys
	}
}
