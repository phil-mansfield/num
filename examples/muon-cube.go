package main

import (
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/phil-mansfield/num/rand"
	"github.com/phil-mansfield/num/objects/geom"
	"github.com/phil-mansfield/num/objects/vec"
)

const (
	width = 32.0
)

var (
	anchorBottom = []float64{0, 0, 0}
	anchorTop = []float64{width, width, width}
	bottom = geom.NewPlane(anchorBottom, []float64{0, 0, -1})
	left = geom.NewPlane(anchorBottom, []float64{0, -1, 0})
	back = geom.NewPlane(anchorBottom, []float64{-1, 0, 0})
	top = geom.NewPlane(anchorTop, []float64{0, 0, 1})
	right = geom.NewPlane(anchorTop, []float64{0, 1, 0})
	front = geom.NewPlane(anchorTop, []float64{1, 0, 0})
	sides = []*geom.Plane{top, left, bottom, right, front, back}
)

func cosSqr(x float64) float64 {
	cosX := math.Cos(x)
	return cosX * cosX
}

func muonDistance(angles []float64, line *geom.Line, gen *rand.Generator) float64 {
	line.Anchor[0] =  gen.Uniform(0, width)
	line.Anchor[1] =  gen.Uniform(0, width)
	line.Anchor[2] =  gen.Uniform(0, width)
	angles[0] = gen.MonteCarlo(cosSqr, 0, 2 * math.Pi, 0, 1)
	angles[1] = gen.Uniform(0, 2 * math.Pi)
	
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

func cmToMeV(dist float64) float64 {
	return 110. * math.Pow(dist / 32.0, 2.05)
}

func MeVToFlux(E float64) float64 {
	logE := math.Log10(E)
	var logF float64
	if logE < 2 {
		logF = 1e-6 * (logE - 2.0) / 2.0
	} else if logE > 3 {
		logF = 1e-6 * (logE - 3.0) * -2.0
	} else {
		logF = 1e-6
	}
	return math.Pow(10.0, logF)
}

func main() {
	if len(os.Args) != 3 {
		panic("")
	}

	trials, err := strconv.Atoi(os.Args[1])
	binNum, err := strconv.Atoi(os.Args[2])
	if err != nil { panic(err) }
	gen := rand.NewTimeSeed(rand.Default)

	angles := make([]float64, 2)
	line := geom.NewLine([]float64{0.5, 0.5, 0.5}, []float64{0.5, 0.5, 0.5})

	bins := make([]int, binNum)
	maxDist := math.Sqrt(3) * width

	for i := 0; i < trials; i++ {
		dist := muonDistance(angles, line, gen)
		binIdx := int(float64(binNum) * dist / maxDist)
		bins[binIdx]++
	}

	var percentBins vec.Vector = make([]float64, binNum)
	for i := 0; i < binNum; i++ {
		percentBins[i] = float64(bins[i] * binNum) / float64(trials)
	}

	binWidth := maxDist / float64(binNum)
	sum := 0.0
	measuredFluxSum := 0.0
	lowFluxSum := 0.0
	fmt.Printf("%15s %15s %15s %15s\n",
		"# dist [cm]", "E [MeV]", "prob density", "prob integral")
	for i := binNum - 1; i >= 0; i-- {
		sum += percentBins[i] / float64(binNum)
		x := binWidth * (float64(i) + 0.5)
		E := cmToMeV(x)
		flux := MeVToFlux(E)

		measuredFluxSum += sum * flux * binWidth
		if E < 110 {
			lowFluxSum += flux * binWidth
		}

		// fmt.Printf("%15.6g %15.6g %15.6g %15.6g\n",
		// 	x, E, percentBins[i], sum)
	}
	println("Full range measured flux:", measuredFluxSum)
	println("Actual in-range flux:", lowFluxSum)
	println("Flux ratio:", measuredFluxSum / lowFluxSum)
}
