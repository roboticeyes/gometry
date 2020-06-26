package geom

import (
	"github.com/go-gl/mathgl/mgl32"
)

type GeometryType int

const (
	LineSetType = iota
	TextType
	PointCloudType
	MeshType
	ImageType
)

type Geometry interface {
	SetName(name string)
	Valid() bool
}

type Geometry2D interface {
	Geometry
	Width() int
	Height() int
	MinBound() mgl32.Vec2
	MaxBound() mgl32.Vec2
}

type geometry struct {
	geomType  GeometryType
	dimension int
	name      string
}

func (g *geometry) SetName(name string) {
	g.name = name
}

type geometry3D struct {
	geometry
}

type TriangleMesh interface {
}

type triangleMesh struct {
	geometry3D
}

type PointCloud interface {
}

type pointCloud struct {
	geometry3D
}
