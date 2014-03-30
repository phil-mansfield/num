package geom

import (
	"github.com/phil-mansfield/num/objects/vec"
	"github.com/phil-mansfield/num"
)

type Line struct {
	Normal, Anchor vec.Vector
}

type FiniteLine struct {
	Line
	Length float64
}

type Plane struct {
	Anchor, Normal vec.Vector
	anchorDotNormal float64
}

type FinitePlane struct {
	Plane
	CoplanarX, CoplanarY vec.Vector
	Width, Height float64
}

// TODO: NewFooAt functions.

func NewLine(anchor, normal vec.Vector) *Line {
	if len(anchor) != len(normal) {
		panic("")
	}

	line := new(Line)
	copy(line.Anchor, anchor)
	normal.NormalizeAt(line.Normal)
	return line
}

func NewFiniteLine(anchor, normal vec.Vector, length float64) *FiniteLine {
	if len(anchor) != len(normal) {
		panic("")
	}

	line := new(FiniteLine)
	copy(line.Anchor, anchor)
	normal.NormalizeAt(line.Normal)

	if length < 0 {
		line.Length = length * -1
		line.Normal.ScaleAt(-1, line.Normal)
	} else {
		line.Length = length
	}

	return line
}

func NewPlane(anchor, normal vec.Vector) *Plane {
	if len(anchor) != len(normal) {
		panic("")
	}

	plane := new(Plane)
	copy(plane.Anchor, anchor)
	normal.NormalizeAt(plane.Anchor)
	plane.anchorDotNormal = vec.Dot(plane.Anchor, plane.Normal)

	return plane
}

func NewFinitePlane(anchor, coplanarX, coplanarY vec.Vector, width, height float64) *FinitePlane {
	if len(anchor) != len(coplanarX) || len(anchor) != len(coplanarY) {
		panic("")
	}

	plane := new(FinitePlane)
	plane.Width = width
	plane.Height = height
	coplanarX.NormalizeAt(plane.CoplanarX)
	coplanarY.NormalizeAt(plane.CoplanarY)
	vec.CrossAt(plane.CoplanarX, plane.CoplanarY, plane.Normal)

	// If coplanarX and coplanarY are colinear, we'll have issues.
	if num.AlmostEqual(0.0, plane.Normal.Norm()) {
		panic("")
	}

	return plane
}

func PlaneIntersectionAt(plane *Plane, line *Line, target vec.Vector) {
	dist := (plane.anchorDotNormal - vec.Dot(line.Anchor, plane.Normal)) /
		vec.Dot(line.Normal, plane.Normal)

	line.Normal.ScaleAt(dist, target)
	vec.Add(line.Anchor, target)
}

func PlaneIntersection(plane *Plane, line *Line) vec.Vector {
	if len(line.Anchor) != len(plane.Anchor) {
		panic("")
	}

	target := make([]float64, len(line.Anchor))
	PlaneIntersectionAt(plane, line, target)
	return target
}
