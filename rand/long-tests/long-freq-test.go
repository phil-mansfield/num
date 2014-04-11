package main

import (
	"fmt"
	"github.com/phil-mansfield/num/rand"
)

const (
	genType = rand.Sys
	subseqs = 1000
	reseeds = 10

	bins = 12
)

var (
	totalSeqLen = int(1e9)
	subseqLen = totalSeqLen / subseqs
	seedTrials = subseqs / reseeds
)

func main() {
	for i := 0; i < reseeds; i++ {
		println("Reseed:", i)
		gen := rand.NewTimeSeed(genType)
		chis, ps := rand.FrequencyTest(gen, bins, subseqLen * seedTrials, seedTrials)
		for i, _ := range ps {
			fmt.Printf("%g %g\n", ps[i], chis[i])
		}
	}
}
