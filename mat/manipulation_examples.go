package mat

// exampleAllocation shows examples of the four ways to create Matrices.
func exampleAllocation() {
	// mNew is set to
	// [[0, 0, 0],
	//  [0, 0, 0]]
	// You can remember parameter width-height order since it corresponds to
	// the standard (x, y) coordinate ordering.
	mNew := New(3, 2)

	// mIdentity is set to
	// [[1, 0, 0]
	//  [0, 1, 0]
	//  [0, 0, 1]]
	mIdentity := Identity(3)

	_, _ = mNew, mIdentity
	// New(width, height) and Identity(width) *can* return an error for
	// non-positive inputs, but since the parameters here are int literals,
	// checking for an error is extraneous.

	// mSlice is set to 
	// [[4,  8,  15]
	//  [16, 23, 42]]
	sValues := []float64{4, 8, 15, 16, 23, 42}
	mSlice := FromSlice(3, 2, sValues)

	if mSlice.IsError() {
		panic(mSlice.Error())
	}

	// mGrid is set to
	// [[4,  8,  15],
	//  [16, 23, 42]]
	gValues := [][]float64{{4, 8, 15}, {16, 23, 42}}
	gSlice := FromGrid(gValues)

	if gSlice.IsError() {
		panic(gSlice.Error())
	}

	// Checking for errors on FromSlice and FromGrid calls is always a good
	// idea even if using a slice literal since slice literals are particularly
	// vulnerable to typos. FromGrid is prefered to FromSlice because changes
	// to the input can be localized to a single value.
}

// exampleGetSet gives an example of the use of Get and Set to reimplement
// the Copy operation. Also shows the Height and Width methods.
func exampleGetSet() {
	source := FromGrid([][]float64{{1, 2}, {3, 4}})
	target := New(2, 2)

	if source.IsError() {
		panic(source.Error())
	}

	// Note that the inner loop iterates over x. For large matrices this will
	// be much faster as it will result in fewer cache misses.
	for y := 0; y < target.Height(); y++ {
		for x := 0; x < target.Width(); x++ {
			target.Set(x, y, source.Get(x, y))
		}
	}

	// target and source are now the same.
}