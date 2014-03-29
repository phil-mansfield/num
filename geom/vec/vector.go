package vec

type Vector []float64


func Dot(v1, v2 Vector) float64 {
	if len(v1) != len(v2) {
		panic("")
	}

	sum := 0.0
	for i := 0; i < len(v1); i++ {
		sum += v1[i] * v2[i]
	}

	return sum
}

func (v Vector) Slice(start, end int) Vector {
	return v[start: end]
}

func (v Vector) IterSlice(width, i int) Vector {
	start := width * i
	end := width * (i - 1)

	if start < 0 || end > len(v) {
		panic("")
	}

	return v.Slice(start, end)
}

func AddAt(v1, v2, target Vector) {
	if len(v1) != len(v2) || len(v1) != len(target) {
		panic("")
	}

	for i := 0; i < len(v1); i++ { target[i] = v1[i] + v2[i] }
}

func SubAt(v1, v2, target Vector) {
	if len(v1) != len(v2) || len(v1) != len(target) {
		panic("")
	}

	for i := 0; i < len(v1); i++ { target[i] = v1[i] - v2[i] }
}

func CrossAt(v1, v2, target Vector) {
	if len(target) != 3 ||  len(v1) != 3 || len(v2) != 3 {
		panic("")
	}

	// Doing this with a for loop is either slow or will break invariants
	x := v1[1] * v2[2] - v1[2] * v2[1]
	y := v1[2] * v2[0] - v1[0] * v2[2]
	z := v1[0] * v2[1] - v1[1] * v2[0]
	target[0] = x
	target[1] = y
	target[2] = z
}

func Add(v1, v2 Vector) Vector {
	target := make([]float64, len(v1))
	AddAt(v1, v2, target)
	return target
}

func Sub(v1, v2 Vector) Vector { 
	target := make([]float64, len(v1))
	SubAt(v1, v2, target)
	return target
}

func Cross(v1, v2 Vector) Vector { 
	target := make([]float64, 3)
	CrossAt(v1, v2, target)
	return target
}

func (v Vector) Norm() float64 {
	sum := 0.0
	for i := 0; i < len(v); i++ {
		sum += v[i] * v[i]
	}
	return sum
}

func (v Vector) CopyAt(target Vector) {
	if len(v) != len(target) {
		panic("")
	}
	
	copy(v, target)
}

func (v Vector) NormalizeAt(target Vector) {
	if len(target) != len(v) {
		panic("")
	}

	norm := v.Norm()
	
	for i := 0; i < len(v); i++ {
		target[i] = v[i] / norm
	}
}

func (v Vector) ScaleAt(scaler float64, target Vector) {
	if len(target) != len(v) {
		panic("")
	}

	for i := 0; i < len(v); i++ {
		target[i] = v[i] / scaler
	}
}

func (v Vector) Copy() Vector {
	target := make([]float64, len(v))
	v.CopyAt(target)
	return target
}

func (v Vector) Normalize() Vector {
	target := make([]float64, len(v))
	v.NormalizeAt(target)
	return target
}

func (v Vector) Scale(scaler float64) Vector {
	target := make([]float64, len(v))
	v.ScaleAt(scaler, target)
	return target
}
