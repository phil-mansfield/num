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
	line.Anchor = make([]float64, len(anchor))
	copy(line.Anchor, anchor)
	line.Normal = normal.Normalize()
	return line
}

func NewFiniteLine(anchor, normal vec.Vector, length float64) *FiniteLine {
	if len(anchor) != len(normal) {
		panic("")
	}

	line := new(FiniteLine)
	line.Anchor = make([]float64, len(anchor))
	copy(line.Anchor, anchor)
	line.Normal = normal.Normalize()

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
	plane.Anchor = make([]float64, len(anchor))
	copy(plane.Anchor, anchor)
	plane.Normal = normal.Normalize()
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
	plane.CoplanarX = coplanarX.Normalize()
	plane.CoplanarY = coplanarY.Normalize()
	plane.Normal = vec.Cross(plane.CoplanarX, plane.CoplanarY)

	// If coplanarX and coplanarY are colinear, we'll have issues.
	if num.AlmostEqual(0.0, plane.Normal.Norm()) {
		panic("")
	}

	return plane
}

// This should maybe be a plane method.
func PlaneIntersectionAt(plane *Plane, line *Line, target vec.Vector) float64 {
	dist := (plane.anchorDotNormal - vec.Dot(line.Anchor, plane.Normal)) /
		vec.Dot(line.Normal, plane.Normal)

	if target == nil { return dist }

	line.Normal.ScaleAt(dist, target)
	vec.AddAt(line.Anchor, target, target)
	return dist
}

func PlaneIntersection(plane *Plane, line *Line) (vec.Vector, float64) {
	if len(line.Anchor) != len(plane.Anchor) {
		panic("")
	}

	target := make([]float64, len(line.Anchor))
	dist := PlaneIntersectionAt(plane, line, target)
	return target, dist
}
