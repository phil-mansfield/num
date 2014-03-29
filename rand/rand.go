package rand

import (
	"math"
	"time"
)

type generatorBackend interface {
	Init(seed uint64)
	Next() float64 // In the range [0, 1)
}

type Generator struct {
	backend generatorBackend
}

type GeneratorType uint8
const (
	Xorshift GeneratorType = iota
	GoRand
	MultiplyWithCarry
	Tausworthe
	GslRand
)

func NewTimeSeed(gt GeneratorType) *Generator {
	return New(gt, uint64(time.Now().UnixNano()))
}

func New(gt GeneratorType, seed uint64) *Generator {
	var backend generatorBackend

	switch(gt) {
	case Xorshift:
		backend = new(xorshiftGenerator)
	case GoRand:
		backend = new(goRandGenerator)
	case MultiplyWithCarry:
		backend = new(multiplyWithCarryGenerator)
	case Tausworthe:
		backend = new(tauswortheGenerator)
	case GslRand:
		backend = new(gslRandGenerator)
	default:
		panic("Unrecognized GeneratorType")
	}

	backend.Init(seed)
	gen := &Generator{ backend }
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
func (gen *Generator) UniformInt(low, high int64) int64 {
	unif := gen.Uniform(0, float64(high - low + 1))
	return low + int64(unif)
}

// It would be possible to get rid of about half our calls to gen.Next()
// here, by storing math.Sin(gen.Uniform(0, 2 * math.Pi)) * radius
// (same Uniform as in dx calculation) somewhere. We may also want to
// check whether or not it's faster to just monte-carlo our way into
// a circle instead of using cos and sin.
func (gen *Generator) Gaussian(mu, sigma float64) float64 {
	// We subtract from 1 because our range is [0, 1)
	radius := math.Sqrt(-2.0 * math.Log(1.0 - gen.backend.Next()))
	dx := math.Cos(gen.Uniform(0, 2 * math.Pi)) * radius
	return mu + sigma * dx
}

// In the range [0, n). Lovingly adapted from rand.Perm().
func (gen *Generator) Permutation(n int64) []int64 {
	seq := make([]int64, n)

	for i, _ := range seq {
		seq[i] = int64(i)
	}

	for i, _ := range seq {
		j := gen.UniformInt(0, int64(i))
		seq[i], seq[j] = seq[j], seq[i]
	}

	return seq
}
