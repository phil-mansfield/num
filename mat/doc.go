/*
package mat implements basic matrix operations on real-valued matrices and is
based around the single struct type Matrix. package mat is written in pure Go
and aims to provide a cleaner and more user-friendly interface than would be
found in a direct transliteration of the BLAS, LAPACK, or GSL libraries while
still providing acceptable performance for most algorithms and scientific
applications.

This package comment contains information on provided Matrix operations and
interfaces, allocation and manipulation of Matrices, the utility fuctions
provided by this package, error reporting and propogation, and the three
provided subpackages: package cmat, package optmat, and package optcmat.
Example usage of functions can be found in the mat/*_example.go files contained
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
	var a, b *Matrix
	
	// Initialize a and b.
	
	res := mat.Exp(mat.Add(a, b))

and second, one which computes the result of the sum in place, saving on an
allocation call:

	import "github.com/phil-mansfield/num/mat"
	var a, b *Matrix
	
	// Initialize a and b.
	
	res := mat.New(2, 2)
	res.Exp(A.Add(a, b))

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
Transpose, Inverse, Scale, and Copy. Supported special functions are Exp, Sin,
Cos, Sinh, Cosh, Log, Sqrt, and the general purpose Func.

Usage examples can be found in mat/operation_examples.go

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
element coordinates which are out of bounds of if called on a nil Matrix.
Bounds checking can be done via the m.InBounds(x, y) method. This is done
because inlcuding it would drastically reduce the utility of these operations.

Usage examples can be found in mat/manipulation_examples.go

Utilities:

package mat provides two types of utility funcitons: those for compatibility
checks and those for printing.

Compatible(m1, m2) checks that two matices are the same shape,
MultCompatible(m1, m2) checks that two Matrices can be multiplied together,
TransposeCompatible(m1, m2) checks that m1 is the same shape as the transpose
of m2, and m.InBounds(x, y) checks that the (x, y) coordinate pair is within
the bounds of m.

m.Print() and m.Printf(fmt) print a Matrix to stdout. The output will have the
following format:

	import "github.com/phil-mansfield/num/mat"

	values = []float64{4, 8, 15, 16, 23, 42}
	m := mat.FromSlice(3, 2, values)
	
	m.Print()
	
	// Output is
	// [[4, 8, 15],
	//  [16, 23, 42]]

Usage examples can be found in mat/util_examples.go

Error Handling:

With the exception of bounds errors in m.Get(x, y) and m.Set(x, y, value),
functions in package mat do not panic by default and instead return error
structs. This default behavior can be switched by calling TogglePanic(). The
use of this function in non-debugging contexts is strongly discouraged.

All explicit errors which are returned by funcitons in package mat are
pointers to the MatrixError struct. Most operations do not return explicit 
errors and instead return Matrices. If an error occurs in such a function,
the result Matrix becomes a 0 by 0 Matrix and its method m.IsError() will
return true. Such Matrices are referred to as "error Matrices" in function
documentation and implement Go's error interface. A MatrixError struct can be
extracted from an error Matrix by the m.MatrixError() method.

If an error Matrix is used as a non-target parameter to a package mat function,
the error is propogated to the new result. This allows for multiple Matrix
operations to be done in one line without losing the location of the first 
error:

	import "github.com/phil-mansfield/num/mat"

	a := mat.FromSlice(2, 2, []float64{0, 1, 0, 1})
	b := mat.FromSlice(3, 3, []float64{2, 0, 0, 0, 2, 0, 0, 0, 2})

	if a.IsError() || b.IsError() {
		// Error handling or panic
	}

	res := mat.Exp(mat.Add(a, b))

	if res.IsError() {
		err := res.MatrixError()
		println(err.Operation)
		// Output is "Add", not "Exp".
	}

Error strings will not contain stacktraces, but these can be obtained from the
Stack field of MatrixError.

Usage examples can be found in mat/error_examples.go

Subpackages:

The subpackage cmat/ implements the same Matrix operations as package mat for
complex-valued Matrices. In the interest of providing clean interfaces several
non-trivial optimizations based on parameter-spamming and algorithm selection
have been ignored. These clunkier, optimized interfaces can be found in the
optmat/ and optcmat/ subpackages. These pacakges are not currently implemented
and will remain that way until package mat is completed.
*/
package mat