package rand

import (
	"math"
	"time"
	
	"github.com/phil-mansfield/num/objects/vec"
	"github.com/phil-mansfield/num/objects/geom"
)

// TODO: Checker whether or not we get faster runtimes with inheritance.

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
func (gen *Generator) UniformInt(low, high int) int {
	unif := gen.Uniform(0, float64(high - low + 1))
	return low + int(unif)
}

func (gen *Generator) FinitePlaneAt(plane *geom.FinitePlane, target vec.Vector) {
	if len(plane.Normal) != len(target) {
		panic("")
	}

	x := gen.Uniform(0, plane.Width)
	y := gen.Uniform(0, plane.Height)

	xVec := plane.CoplanarX.Scale(x)
	plane.CoplanarY.ScaleAt(y, target)
	vec.AddAt(xVec, target, target)
	vec.AddAt(plane.Anchor, target, target)
}

func (gen *Generator) FiniteLineAt(line *geom.FiniteLine, target vec.Vector) {
	if len(line.Normal) != len(target) {
		panic("")
	}

	s := gen.Uniform(0, line.Length)
	line.Normal.ScaleAt(s, target)
	vec.AddAt(line.Anchor, target, target)
}

func (gen *Generator) FinitePlane(plane *geom.FinitePlane) vec.Vector {
	target := make([]float64, len(plane.Normal))
	gen.FinitePlaneAt(plane, target)
	return target
}

func (gen *Generator) FiniteLine(line *geom.FiniteLine) vec.Vector {
	target := make([]float64, len(line.Normal))
	gen.FiniteLineAt(line, target)
	return target
}

func (gen *Generator) SphereAt(radius float64, target vec.Vector) {
	if len(target) != 3 {
		panic("")
	}

	sqrSum := 2.0
	var x, y float64
	for sqrSum >= 1 {
		x = gen.Uniform(-1, 1)
		y = gen.Uniform(-1, 1)
		sqrSum = x * x + y * y
	}
	s := math.Sqrt(1 - sqrSum)

	target[0] = radius * 2 * x * s
	target[1] = radius * 2 * y * s
	target[2] = radius * 1 - 2 * sqrSum
}

func (gen *Generator) Sphere(radius float64) vec.Vector {
	target := make([]float64, 3)
	gen.SphereAt(radius, target)
	return target
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
