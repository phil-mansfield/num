package rand

import (
	"math"
	"time"
	
	"github.com/phil-mansfield/num"
)

type generatorBackend interface {
	Init(seed uint64)
	Next() float64 // In the range [0, 1)
}

type Generator struct {
	backend generatorBackend
	savedGaussian bool
	nextGaussianDx float64
}

type GeneratorType uint8
const (
	Xorshift GeneratorType = iota
	Golang
	MultiplyWithCarry
	Tausworthe
	Gsl
	Sys

	Default = Tausworthe
)

func NewTimeSeed(gt GeneratorType) *Generator {
	return New(gt, uint64(time.Now().UnixNano()))
}

func New(gt GeneratorType, seed uint64) *Generator {
	var backend generatorBackend

	switch(gt) {
	case Xorshift:
		backend = new(xorshiftGenerator)
	case Golang:
		backend = new(goRandGenerator)
	case MultiplyWithCarry:
		backend = new(multiplyWithCarryGenerator)
	case Tausworthe:
		backend = new(tauswortheGenerator)
	case Gsl:
		backend = new(gslRandGenerator)
	case Sys:
		backend = new(sysRandGenerator)
	default:
		panic("Unrecognized GeneratorType")
	}

	backend.Init(seed)
	gen := &Generator{ backend, false, -1 }
	return gen
}

// These interfaces are intended to be user friendly, not performance
// friendly. There are a few extraneous floating point operations,
// default variable size is 64 bits, and function calls are made 
// in places where they are not maybe neccesary. Deal with it.

func (gen *Generator) Uniform(low, high float64) float64 {
	return (gen.backend.Next() * (high - low)) + low
}

// Inclusive on the upper bound.
func (gen *Generator) UniformInt(low, high int) int {
	unif := gen.Uniform(0, float64(high - low + 1))
	return low + int(unif)
}

func (gen *Generator) MonteCarlo(f num.Func1D, lowX, highX, lowY, highY float64) float64 {
	for {
		x := gen.Uniform(lowX, highX)
		y := gen.Uniform(lowY, highY)

		if y < f(x) { return x }
	}
}

// It would be possible to get rid of about half our calls to gen.Next()
// here, by storing math.Sin(gen.Uniform(0, 2 * math.Pi)) * radius
// (same Uniform as in dx calculation) somewhere. We may also want to
// check whether or not it's faster to just monte-carlo our way into
// a circle instead of using cos and sin.
func (gen *Generator) Gaussian(mu, sigma float64) float64 {
	if gen.savedGaussian {
		gen.savedGaussian = false
		return mu + sigma * gen.nextGaussianDx
	}

	// We subtract from 1 because our range is [0, 1)
	radius := math.Sqrt(-2.0 * math.Log(1.0 - gen.backend.Next()))
	theta := gen.backend.Next() * 2 * math.Pi
	cosTheta := math.Cos(theta)
	sinTheta := math.Sin(theta)
	dx := cosTheta * radius
	dy := sinTheta * radius    

	gen.savedGaussian = true
	gen.nextGaussianDx = dy
	return mu + sigma * dx
}

// In the range [0, n). Lovingly adapted from rand.Perm().
func (gen *Generator) PermutationAt(n int, target []int) []int {
	for i, _ := range target {
		target[i] = i
	}

	for i, _ := range target {
		j := gen.UniformInt(0, i)
		target[i], target[j] = target[j], target[i]
	}

	return target
}

func (gen *Generator) Permutation(n int) []int {
	target := make([]int, n)
	gen.PermutationAt(n, target)
	return target
}
