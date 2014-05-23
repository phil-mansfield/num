/*
package mat implements basic matrix operations on real-valued matrices. The
subpackage cmat/ implements the operations for complex-valued matrices. In the
interest of providing clean interfaces several non-trivial optimizations based
on parameter-spamming and algorithm selection have been ignored. These
clunkier, optimized interfaces can be found in the optmat/ and optcmat/
subpackages.

General purpose:

 Equal()
 Compatible()
 MultCompatible()

Error handling:

 m.IsError()
 m.Error()

Matrix properties:

 m.Get()
 m.Set()
 m.Height()
 m.Width()
 m.Slice()
 m.Grid()

Initializers:

 New()
 Identity()
 FromSlice()
 FromGrid()

Arithmetic operations:

 Add()
 Sub()
 Mult()
 Scale()

Special functions:

 Exp()
 Sin()
 Cos()
 Sinh()
 Cosh()
 Log()
 Sqrt()

Linear algebra:

 m.Eigenvalues()
 m.Eigenvectors()
 m.Determinant()

 Transpose()
 Invert()

Printing:

 m.Print()
 m.Printf()

*/