package rand

import (
	"math/rand"
)

type goRandGenerator struct {
	r *rand.Rand
}

func (gen *goRandGenerator) Init(seed uint64) {
	src := rand.NewSource(int64(seed))
	gen.r = rand.New(src)
}

func (gen *goRandGenerator) Next() float64 {
	return gen.r.Float64()
}
