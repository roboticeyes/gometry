package geom

import "github.com/go-gl/mathgl/mgl32"

type Geometry3D interface {
	Geometry

	// MinBound returns min bounds for geometry coordinates
	MinBound() mgl32.Vec3

	// MaxBound returns max bounds for geometry coordinates
	MaxBound() mgl32.Vec3

	// Center returns the center of the geometry coordinates
	Center() mgl32.Vec3

	// Translate translates the geometry by the given 3D vector
	Translate(t mgl32.Vec3)

	// Scale scales the geometry by the given 3D sx, sy, sz values
	Scale(s mgl32.Vec3)

	// Rotate rotates the geometry by the given quaternion
	Rotate(q mgl32.Quat)

	// SetVertices replaces all vertices
	SetVertices(vertices []mgl32.Vec3)

	// Vertices returns the raw vertices of the geometry (coordinates)
	Vertices() []mgl32.Vec3

	// Returns a specific vertex of the vertex list
	Vertex(idx int) mgl32.Vec3
}

type geometry3D struct {
	geometry

	vertices []mgl32.Vec3
}

func (g *geometry3D) SetVertices(vertices []mgl32.Vec3) {
	g.vertices = vertices
}

func (g *geometry3D) Vertices() []mgl32.Vec3 {
	return g.vertices
}

func (g *geometry3D) Vertex(idx int) mgl32.Vec3 {
	return g.vertices[idx]
}

func (g *geometry3D) MinBound() mgl32.Vec3 {
	return mgl32.Vec3{0, 0, 0}
}

func (g *geometry3D) MaxBound() mgl32.Vec3 {
	return mgl32.Vec3{0, 0, 0}
}

func (g *geometry3D) Center() mgl32.Vec3 {
	return mgl32.Vec3{0, 0, 0}
}

func (g *geometry3D) Translate(t mgl32.Vec3) {
	panic("not implemented")
}

func (g *geometry3D) Scale(s mgl32.Vec3) {
	panic("not implemented")
}

func (g *geometry3D) Rotate(q mgl32.Quat) {
	panic("not implemented")
}
