package geom

import (
	"github.com/phil-mansfield/num/objects/vec"
)

type Line struct {
	Normal, Anchor vec.Vector
}

type Plane struct {
	Anchor, Normal vec.Vector
	anchorDotNormal float64
}

func NewLine(anchor vec.Vector, direction vec.Vector) *Line {
	if len(anchor) != len(direction) {
		panic("")
	}

	line := new(Line)
	copy(line.Anchor, anchor)
	direction.NormalizeAt(line.Normal)
	return line
}

func NewPlane(anchor vec.Vector, direction vec.Vector) *Plane {
	if len(anchor) != len(direction) {
		panic("")
	}

	plane := new(Plane)
	copy(plane.Anchor, anchor)
	direction.NormalizeAt(plane.Anchor)
	plane.anchorDotNormal = vec.Dot(plane.Anchor, plane.Normal)

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
