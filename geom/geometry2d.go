package geom

import "github.com/go-gl/mathgl/mgl32"

type Geometry2D interface {
	Geometry
	Width() int
	Height() int
	MinBound() mgl32.Vec2
	MaxBound() mgl32.Vec2
}
