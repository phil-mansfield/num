/*
package mat implements basic matrix operations on real-valued matrices, based
around the type Matrix. package mat is written in pure Go and aims to provide
a cleaner and more user-friendly interface than would be found in a direct
transliteration of the BLAS, LAPACK, or GSL libraries while still providing
acceptable performance for most scientific applications.

This package comment contains information on provided Matrix operations and
interfaces, allocation and manipulation of Matrices, the utility fuctions
provided by this package, error reporting and propogation, and the three
provided subpackages: package cmat, package optmat, and package optcmat.
Example usage of functions can be found in the *_example.go files contained
in this pacakge.

Matrix Operations and Interfaces:

Any operation which results in a new Matirx has two different interfaces. The
first is a function call which returns a new Matrix of the correct size and the
second is a method call on a target Matrix which the results of the operation
will be stored within. For programmer convenience, the target Matrix is also
returned.

For example, consider two methods for computing the exponenet of a sum of two
matrices, A and B: first, one which creates new matrices with each operation:

	import "github.com/phil-mansfield/num/mat"
	var A, B *Matrix
	
	// Initialize A and B.
	
	res := mat.Exp(mat.Add(A, B))

	if res.IsError() {
		// More sophisticated error-handling should be added if approriate
		panic(res.Error())
	}

and second, one which computes the result of the sum in place, saving on an
allocation call:

	import "github.com/phil-mansfield/num/mat"
	var A, B *Matrix
	
	// Initialize A and B.
	
	res := mat.New(2, 2)
	res.Exp(A.Add(A, B))
	
	if res.IsError() {
		// More sophisticated error-handling should be added if approriate
		panic(res.Error())
	}

As long as the target Matrix is the same size as the result of its method,
it is always safe to call that method, even if that Matrix is a method
argument. For certain operations this may result in a temporary matrix being
allocated underneath the hood. It is garuanteed that regardless of
implementation Add(m1, m2), Sub(m1, m2), and Scale(m1, m2) will not allocate
any temporary matrices.

Routines which return a value that is not a Matrix and only depend on a
single Matrix, such as m.Eigenvectors() or m.Determinant(), will always be
implemented only as methods. Operations which depend on two Matrices, such as
Equal(m1, m2) will always be implemented as non-method function calls.

Supported binary Matrix operations are Mult, Add, Sub, and Equal. Supported
unary Marix operation are Eigenvalues, Eigenvectors, Trace, Determinant,
Transpose, Inverse, and Scale. Supported special functions are Exp, Sin, Cos,
Sinh, Cosh, Log, Sqrt, and the general purpose Func.

Usage examples can be found in num/mat/operation_examples.go

Allocation and Manipulation:

There are four functions which allow for the creation of new Matrices.
New(width, height), Identity(width), FromSlice(width, height, values), and
FromGrid(values). The first two allow for the creation of the two most useful
"blank" Matrices, while the latter two allow for the creation of arbitrary
Matrices.

Matrix dimensions can be accesses via the m.Width() and m.Height() methods and
copies of data within Matrices can be generated with the m.Slice() and
m.Grid() methods.

Individual Matrix elements can be accessed and changed through the m.Get(x, y)
and m.Set(x, y) methods. Get and Set are unique among all the operations in
package mat in that they, in analogy to slice accessing, will panic if given
element coordinates which are out of bounds. Bounds checking can be done via
the m.InBounds(x, y) method.

Usage examples can be found in manipulation_examples.go

Utilities:

Usage examples can be found in num/mat/util_examples.go

Error Handling:

Usage examples can be found in error_examples.go

Subpackages:

The subpackage cmat/ implements the same Matrix operations as package mat for
complex-valued Matrices. In the interest of providing clean interfaces several
non-trivial optimizations based on parameter-spamming and algorithm selection
have been ignored. These clunkier, optimized interfaces can be found in the
optmat/ and optcmat/ subpackages. These pacakges are not currently implemented
and will remain that way until package mat is completed.
*/
package mat