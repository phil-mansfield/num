package spatial

import (
	"fmt"
	"math"
)

const (
	ListNil = -1
)

// ListPoint2D is a combination of a Point2D and a linked list node. Given
// indices point into a ListPoint2D array within a parent grid.
type ListPoint2D struct {
	Point2D
	NextIdx, PrevIdx int
	grid    *ListGrid2D // This is wasteful, but makes for nice interfaces.
}

// ListGrid2D is a structure that grids a collection of Point2D's and places them
// into low-overhead lists. Cells are internally represented as a flat slice such
// that Heads[x * GridWidth + y] is logically equivelent to Heads[x][y]. Cells are
// inclusive on their lower spatial bound and exclusive on their upper spatial
// bound.
type ListGrid2D struct {
	Points       []ListPoint2D
	Sizes, Heads []int

	GridWidth           int
	CellWidth, BoxWidth float64
}

func (pt *ListPoint2D) String() string {
	return fmt.Sprintf("(%f, %f)[%d] (Next: %d, Prev: %d)",
		pt.X, pt.Y, pt.Cell(), pt.NextIdx, pt.PrevIdx)
}

// Cell returns the index of the grid cell which the given point is a member of.
func (pt *ListPoint2D) Cell() int {
	gridX := int(math.Floor(pt.X / pt.grid.CellWidth))
	gridY := int(math.Floor(pt.Y / pt.grid.CellWidth))
	return pt.grid.FlattenIndex(gridX, gridY)
}

// IsHead returns true if the given ListPoint2D is the first item in its list
// and false otherwise.

func (pt *ListPoint2D) IsHead() bool { return pt.PrevIdx == ListNil }

// IsTail returns true if the given ListPoint2D is the last itme in its list
// and false otherwise.
func (pt *ListPoint2D) IsTail() bool { return pt.NextIdx == ListNil }

// Next returns the ListPoint2D which follows the given one. If there is no
// such point, nil is returned.
func (pt *ListPoint2D) Next() *ListPoint2D {
	if pt.NextIdx == -1 {
		return nil
	}
	return &pt.grid.Points[pt.NextIdx]
}

// Prev returns the ListPoint2D which preceeds the given one. If there is
// no such point, nil is returned.
func (pt *ListPoint2D) Prev() *ListPoint2D {
	if pt.PrevIdx == -1 {
		return nil
	}
	return &pt.grid.Points[pt.PrevIdx]
}

// Delete removes the given ListPoint2D from whatever list is is a member of.
func (pt *ListPoint2D) Delete() {
	// TODO: check that it is in a list

	cell := pt.Cell()
	pt.grid.Sizes[cell]--

	next, prev := pt.Next(), pt.Prev()

	if next != nil {
		next.PrevIdx = pt.PrevIdx
	}

	if prev != nil {
		prev.NextIdx = pt.NextIdx
	} else {
		pt.grid.Heads[cell] = pt.NextIdx
	}

	pt.NextIdx = ListNil
	pt.PrevIdx = ListNil
}

// NewListGrid2D returns a new ListGrid2D instance with the given spatial
// width and the given number of cells on each side.
func NewListGrid2D(gridWidth int, boxWidth float64) *ListGrid2D {
	// TODO: sign checks

	grid := new(ListGrid2D)

	grid.Points = make([]ListPoint2D, 0)
	grid.Sizes = make([]int, gridWidth*gridWidth)
	grid.Heads = make([]int, gridWidth*gridWidth)

	for i := 0; i < len(grid.Heads); i++ {
		grid.Heads[i] = -1
	}

	grid.BoxWidth = boxWidth
	grid.GridWidth = gridWidth
	grid.CellWidth = boxWidth / float64(gridWidth)

	return grid
}

// FlattenIndex converts an X, Y coordinate pair into a single index into
// the flat grid cell slices.
//
// Note that it is chache-friendly to iterate over gridY before iterating
// over gridX so that the call signature imitates the cache behavior of
// a 2D array access, e.g. slice[idxSlow][idxFast].
func (grid *ListGrid2D) FlattenIndex(gridX, gridY int) int {
	// TODO: bounds checks
	return grid.GridWidth*gridX + gridY
}

// ModFlattenIndex converts an X, Y coordinate pair into a single index into
// the flat grid cell slices. If that index would be out of bounds, the modulo
// is taken so that it remains in bounds.
func (grid *ListGrid2D) ModFlattenIndex(gridX, gridY int) int {
	modX := gridX % grid.GridWidth
	modY := gridY % grid.GridWidth
	if modX < 0 {
		modX+= grid.GridWidth
	}
	if modY < 0 {
		modY+= grid.GridWidth
	}

	return grid.GridWidth*modX + modY
}

// ListHead returns the first ListPoint2D in the given cell. If there are no
// particles, nil is returned.
func (grid *ListGrid2D) ListHead(cell int) *ListPoint2D {
	// TODO: add bounds checks
	if grid.Heads[cell] == ListNil {
		return nil
	}

	return &grid.Points[grid.Heads[cell]]
}

// You only need to set pt.X and pt.Y before calling this.
func (grid *ListGrid2D) reinsert(pt *ListPoint2D, ptIdx int) {
	// TODO: add bounds checks
	pt.grid = grid
	cell := pt.Cell()
	pt.PrevIdx = ListNil

	next := grid.ListHead(cell)
	if next == nil {
		pt.NextIdx = ListNil
	} else {
		pt.NextIdx = grid.Heads[cell]
		next.PrevIdx = ptIdx
	}

	grid.Heads[cell] = ptIdx
	grid.Sizes[cell]++
}

// Insert inserts a slice of ListPoint2D's into the given grid.
func (grid *ListGrid2D) Insert(pts []Point2D) {
	// TODO: add nil check
	listPts := make([]ListPoint2D, len(pts))

	oldLen := len(grid.Points)
	if oldLen == 0 { 
		grid.Points = listPts
	} else {
		grid.Points = append(grid.Points, listPts...)
	}

	for i := 0; i < len(pts); i++ {
		ptIdx := i + oldLen
		listPts[i].X = pts[i].X
		listPts[i].Y = pts[i].Y
		
		grid.reinsert(&listPts[i], ptIdx)
	}
}

// Move moves the point at the given index to the target point.
func (grid *ListGrid2D) Move(ptIdx int, target *Point2D) {
	// TODO: bounds checks
	pt := &grid.Points[ptIdx]
	pt.Delete()
	pt.X, pt.Y = target.X, target.Y
	grid.reinsert(pt, ptIdx)
}

// Slices returns a slice containing all the 
func (grid *ListGrid2D) Slice(gridIdx int) []int {
	// TODO: bounds checks
	slice := make([]int, grid.Sizes[gridIdx])
	if len(slice) == 0 {
		return slice
	}

	slice[0] = grid.Heads[gridIdx]
	pt := grid.ListHead(gridIdx)
	for i := 0; i < len(slice) - 1; i++ {
		slice[i + 1] = pt.NextIdx
		pt = pt.Next()
	}

	return slice
}
