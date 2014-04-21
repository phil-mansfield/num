package spatial

import (
	"testing"

	"github.com/phil-mansfield/num/rand"
)

var flattenIndexTests = []struct {
	inWidth, inX, inY int
	outIdx int
}{
	{10, 0, 0, 0},
	{10, 0, 1, 1},
	{10, 1, 0, 10},
	{7, 6, 6, 48},
}

func TestFlattenIndex(t *testing.T) {
	boxWidth := 13.0
	for i, test := range flattenIndexTests {
		grid := NewListGrid2D(test.inWidth, boxWidth)
		outIdx := grid.FlattenIndex(test.inX, test.inY)
		if outIdx != test.outIdx {
			t.Errorf("%d. FlattenIndex(%d, %d) => %d, not %d",
				i, test.inX, test.inY, outIdx, test.outIdx)
		}
	}
}

var cellTests = []struct {
	inGridWidth int
	inBoxWidth float64
	inX, inY float64
	outCell int
}{
	{5, 100, 0, 0, 0},
	{5, 100, 0, 30, 1},
	{5, 100, 30, 0, 5},
	{5, 100, 0, 40, 2},
	{5, 100, 40, 0, 10},
	{7, 49, 48, 48, 48},
}

func TestCell(t *testing.T) {
	for i, test := range cellTests {
		grid := NewListGrid2D(test.inGridWidth, test.inBoxWidth)
		pt := ListPoint2D{Point2D{test.inX, test.inY}, ListNil, ListNil, grid}
		outCell := pt.Cell()
		if outCell != test.outCell {
			t.Errorf("%d. ListPoint2D.Cell() => %d, not %d",
				i, outCell, test.outCell)
		}
	}
}

var modFlattenTests = []struct {
	inGridWidth int
	inX, inY int
	refX, refY int
}{
	{5, 0, 6, 0, 1},
	{5, 6, 0, 1, 0},
	{5, 0, -1, 0, 4},
	{5, -1, 0, 4, 0},
	{5, 37, -37, 2, 3},
	{5, -37, 37, 3, 2},
}

func TestModFlattenIndex(t *testing.T) {
	boxWidth := 13.0
	for i, test := range modFlattenTests {
		grid := NewListGrid2D(test.inGridWidth, boxWidth)
		outIdx := grid.ModFlattenIndex(test.inX, test.inY)
		refIdx := grid.FlattenIndex(test.refX, test.refY)
		if outIdx != refIdx {
			t.Errorf("%d. ListGrid2D.ModFlattenIndex(%d, %d) => %d, not %d",
				i, test.inX, test.inY, outIdx, refIdx)
		}
	}
}

func intSliceEq(xs, ys []int) bool {
	if len(xs) != len(ys) {
		return false
	}

	for i := range xs {
		if xs[i] != ys[i] {
			return false
		}
	}

	return true
}

var insertTests = []struct {
	inGridWidth int
	inBoxWidth float64
	pts []Point2D
	outSizes []int
}{
	{3, 9,
		[]Point2D{{0, 0}},
		[]int{1, 0, 0, 0, 0, 0, 0, 0, 0}},
	{3, 9,
		[]Point2D{{0, 0}, {1, 1}},
		[]int{2, 0, 0, 0, 0, 0, 0, 0, 0}},
	{3, 9,
		[]Point2D{{0, 0}, {3, 3}, {6, 6}},
		[]int{1, 0, 0, 0, 1, 0, 0, 0, 1}},
	{3, 9,
		[]Point2D{{0, 0}, {2, 2}, {3, 3}, {6, 6}, {1, 1}, {3, 0}},
		[]int{3, 0, 0, 1, 1, 0, 0, 0, 1}},
}

func TestInsert(t *testing.T) {
	for i, test := range insertTests {
		grid := NewListGrid2D(test.inGridWidth, test.inBoxWidth)
		grid.Insert(test.pts)
		if !intSliceEq(test.outSizes, grid.Sizes) {
			t.Errorf("%d. grid.Insert() gives Sizes = %v, not %v",
				i, grid.Sizes, test.outSizes)
		}
	}
}

var (
	defaultGridWidth = 3
	defaultBoxWidth = 9.0
	defaultPts = []Point2D{{0, 0}, {2, 2}, {3, 3}, {4, 0}, {1, 1}, {3, 0}}
)

var listHeadTests = []struct {
	inGridCell int
	outPtIdx int
}{
	{0, 4},
	{1, ListNil},
	{3, 5},
	{4, 2},
}

func TestListHead(t *testing.T) {
	grid := NewListGrid2D(defaultGridWidth, defaultBoxWidth)
	grid.Insert(defaultPts)
	var refPt *ListPoint2D
	for i, test := range listHeadTests {
		outPt := grid.ListHead(test.inGridCell)

		if test.outPtIdx == ListNil {
			refPt = nil
		} else {
			refPt = &grid.Points[test.outPtIdx]
		}

		if refPt != outPt {
			t.Errorf("%d. grid.TeastListHead(%d) => %v, not %v",
				i, test.inGridCell, outPt, refPt)
		}
	}
}

var isTailTests = []struct {
	inPtIdx int
	out bool
}{
	{0, true},
	{1, false}, 
	{2, true},
	{3, true},
	{4, false},
	{5, false},
}

func TestIsTail(t *testing.T) {
	grid := NewListGrid2D(defaultGridWidth, defaultBoxWidth)
	grid.Insert(defaultPts)
	for i, test := range isTailTests {
		pt := &grid.Points[test.inPtIdx]
		res := pt.IsTail()
		if res != test.out {
			t.Errorf("%d. ListPoint2D.IsTail() => %v, not %v.",
				i, res, test.out)
		}
	}
}

var isHeadTests = []struct {
	inPtIdx int
	out bool
}{
	{0, false},
	{1, false}, 
	{2, true},
	{3, false},
	{4, true},
	{5, true},
}

func TestIsHead(t *testing.T) {
	grid := NewListGrid2D(defaultGridWidth, defaultBoxWidth)
	grid.Insert(defaultPts)
	for i, test := range isHeadTests {
		pt := &grid.Points[test.inPtIdx]
		res := pt.IsHead()
		if res != test.out {
			t.Errorf("%d. ListPoint2D.IsHead() => %v, not %v.",
				i, res, test.out)
		}
	}
}

var prevTests = []struct {
	inIdx, outIdx int
}{
	{0, 1},
	{1, 4},
	{2, ListNil},
	{3, 5},
	{4, ListNil},
	{5, ListNil},
}

func TestPrev(t *testing.T) {
	grid := NewListGrid2D(defaultGridWidth, defaultBoxWidth)
	grid.Insert(defaultPts)
	var ref *ListPoint2D
	for i, test := range prevTests {
		pt := &grid.Points[test.inIdx]
		res := pt.Prev()

		if test.outIdx == ListNil {
			ref = nil
		} else {
			ref = &grid.Points[test.outIdx]
		}
		if res != ref {
			t.Errorf("%d. ListPoint2D.Prev() => %v, not %v.", i, res, ref)
		}
	}
}

var nextTests = []struct {
	inIdx, outIdx int
}{
	{0, ListNil},
	{1, 0},
	{2, ListNil},
	{3, ListNil},
	{4, 1},
	{5, 3},
}

func TestNext(t *testing.T) {
	grid := NewListGrid2D(defaultGridWidth, defaultBoxWidth)
	grid.Insert(defaultPts)
	var ref *ListPoint2D
	for i, test := range nextTests {
		pt := &grid.Points[test.inIdx]
		res := pt.Next()

		if test.outIdx == ListNil {
			ref = nil
		} else {
			ref = &grid.Points[test.outIdx]
		}

		if res != ref {
			t.Errorf("%d. ListPoint2D.Next() => %v, not %v.", i, res, ref)
		}
	}
}

var sliceTests = []struct {
	inGridIdx int
	outSlice []int
}{
	{0, []int{4, 1, 0}},
	{1, []int{}},
	{3, []int{5, 3}},
	{4, []int{2}},
}

func TestSlice(t *testing.T) {
	grid := NewListGrid2D(defaultGridWidth, defaultBoxWidth)
	grid.Insert(defaultPts)

	for i, test := range sliceTests {
		resSlice := grid.Slice(test.inGridIdx)
		if !intSliceEq(resSlice, test.outSlice) {
			t.Errorf("%d. ListGrid2D.Slice(%d) => %v, not %v.",
				i,test.inGridIdx, resSlice, test.outSlice)
		}
	}
}

var deleteTests = []struct {
	inPtIdx int
	targetCell int
	targetSlice []int
}{
	{0, 0, []int{4, 1}},
	{1, 0, []int{4, 0}},
	{2, 4, []int{}},
	{3, 3, []int{5}},
	{4, 0, []int{1, 0}},
	{5, 3, []int{3}},
}

func TestDelete(t *testing.T) {
	for i, test := range deleteTests {
		grid := NewListGrid2D(defaultGridWidth, defaultBoxWidth)
		grid.Insert(defaultPts)

		pt := &grid.Points[test.inPtIdx]
		pt.Delete()
		resSlice := grid.Slice(test.targetCell)

		if !intSliceEq(resSlice, test.targetSlice) {
			t.Errorf("%d. ListPoint2D.Delete() gave the list %v at cell %d, not %v.",
				i, test.targetCell, resSlice, test.targetSlice)
		}
	}
}

var moveTests = []struct {
	inPtIdx int
	targetX, targetY float64
	sourceCell int
	sourceSlice []int
	targetCell int
	targetSlice []int
}{
	{0, 1.0, 1.0, 0, []int{0, 4, 1}, 0, []int{0, 4, 1}},
	{1, 4.0, 4.0, 0, []int{4, 0}, 4, []int{1, 2}},
	{1, 8.0, 8.0, 0, []int{4, 0}, 8, []int{1}},
	{2, 4.0, 4.0, 4, []int{2}, 4, []int{2}},
	{2, 8.0, 8.0, 4, []int{}, 8, []int{2}},
}

func TestMove(t *testing.T) {
	for i, test := range moveTests {
		grid := NewListGrid2D(defaultGridWidth, defaultBoxWidth)
		grid.Insert(defaultPts)

		grid.Move(test.inPtIdx, &Point2D{test.targetX, test.targetY})

		sourceSlice := grid.Slice(test.sourceCell)
		targetSlice := grid.Slice(test.targetCell)

		if !intSliceEq(sourceSlice, test.sourceSlice) {
			t.Errorf("%d. ListPoint2D.Move(%f, %f) left cell %d as %v, not %v.",
				i, test.targetX, test.targetY, test.sourceCell,
				sourceSlice, test.sourceSlice)
		}
		if !intSliceEq(targetSlice, test.targetSlice) {
			t.Errorf("%d. ListPoint2D.Move(%f, %f) left cell %d as %v, not %v.",
				i, test.targetX, test.targetY, test.targetCell,
				targetSlice, test.targetSlice)
		}
	}
}

func BenchmarkListGrid2DInsert(b *testing.B) {
	gen := rand.NewTimeSeed(rand.Xorshift)
	grid := NewListGrid2D(50, 1.0)

	coords := make([]Point2D, b.N)
	rawFloats := make([]float64, b.N * 2)
	gen.UniformAt(0.0, 1.0, rawFloats)
	for i := 0; i < b.N; i++ {
		coords[i].X = rawFloats[2 * i]
		coords[i].Y = rawFloats[2 * i + 1]
	}

	b.ResetTimer()

	grid.Insert(coords)
}

func BenchmarkListGrid2DInsertCacheMiss(b *testing.B) {
	gen := rand.NewTimeSeed(rand.Xorshift)
	grid := NewListGrid2D(1000, 1.0)

	coords := make([]Point2D, b.N)
	rawFloats := make([]float64, b.N * 2)
	gen.UniformAt(0.0, 1.0, rawFloats)
	for i := 0; i < b.N; i++ {
		coords[i].X = rawFloats[2 * i]
		coords[i].Y = rawFloats[2 * i + 1]
	}

	b.ResetTimer()

	grid.Insert(coords)
}

func BenchmarkListGrid2DInsertSingleton(b *testing.B) {
	gen := rand.NewTimeSeed(rand.Xorshift)
	grid := NewListGrid2D(50, 1.0)

	rawFloats := make([]float64, b.N * 2)
	inSlice := make([]Point2D, 1)
	gen.UniformAt(0.0, 1.0, rawFloats)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		inSlice[0].X = rawFloats[i * 2]
		inSlice[0].Y = rawFloats[i * 2 + 1]
		grid.Insert(inSlice)
	}
}

func BenchmarkListGrid2DMoveIter(b *testing.B) {
	gen := rand.NewTimeSeed(rand.Xorshift)

	coords := make([]Point2D, 100 * 1000)
	rawFloats := make([]float64, 2 * len(coords))
	gen.UniformAt(0.0, 1.0, rawFloats)
	for i := 0; i < len(coords); i++ {
		coords[i].X = rawFloats[2 * i]
		coords[i].Y = rawFloats[2 * i + 1]
	}

	grid := NewListGrid2D(50, 1.0)
	grid.Insert(coords)

	rawFloats = make([]float64, b.N * 2)
	coords = make([]Point2D, b.N)
	for i := 0; i < b.N; i++ {
		coords[i].X = rawFloats[2 * i]
		coords[i].Y = rawFloats[2 * i + 1]
	}

	gen.UniformAt(0.0, 1.0, rawFloats)
	

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		grid.Move(i % len(grid.Points), &coords[i])
	}
}
