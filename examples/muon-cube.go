package main

import (
	"math"
	"os"
	"strconv"

	"github.com/phil-mansfield/num/rand"
	"github.com/phil-mansfield/num/objects/geom"
	"github.com/phil-mansfield/num/objects/vec"
)

const (
	width = 1.0
)

var (
	anchorBottom = []float64{0, 0, 0}
	anchorTop = []float64{width, width, width}
	bottom = geom.NewPlane(anchorBottom, []float64{0, 0, -1})
	left = geom.NewPlane(anchorBottom, []float64{0, -1, 0})
	back = geom.NewPlane(anchorBottom, []float64{-1, 0, 0})
	top = geom.NewPlane(anchorBottom, []float64{0, 0, 1})
	right = geom.NewPlane(anchorBottom, []float64{0, 1, 0})
	front = geom.NewPlane(anchorBottom, []float64{1, 0, 0})
	sides = []*geom.Plane{top, left, bottom, right, front, back}
)

func cosSqr(x float64) float64 {
	cosX := math.Cos(x)
	return cosX * cosX
}

func muonDistance(angles []float64, line *geom.Line, gen *rand.Generator) float64 {
	angles[0] = gen.MonteCarlo(cosSqr, 0, 2 * math.Pi, 0, 1)
	angles[1] = gen.Uniform(0, 2 * math.Pi)
	
	println(len(angles), len(line.Normal))
	vec.FromAnglesAt(1, angles, line.Normal)
	minPos := 2 * width
	maxNeg := -2 * width
	for j := 0; j < len(sides); j++ {
		dist := geom.PlaneIntersectionAt(sides[j], line, nil)
		if dist > 0 && dist < minPos {
			minPos = dist
		} else if dist < 0 && dist > maxNeg {
			maxNeg = dist
		}
	}
	return minPos - maxNeg
}
	
func main() {
	if len(os.Args) != 2 {
		panic("")
	}

	trials, err := strconv.Atoi(os.Args[1])
	if err != nil { panic(err) }
	gen := rand.NewTimeSeed(rand.Default)

	angles := make([]float64, 2)
	line := geom.NewLine([]float64{0.5, 0.5, 0.5}, []float64{0.5, 0.5, 0.5})

	for i := 0; i < trials; i++ {
		dist := muonDistance(angles, line, gen)
		println(dist)
	}
}
