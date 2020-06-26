package geom

import "github.com/go-gl/mathgl/mgl32"

type TriangleMesh interface {
	Geometry3D
}

type triangleMesh struct {
	geometry3D
}

func NewTriangleMesh(name string, vertices []mgl32.Vec3) TriangleMesh {
	mesh := &triangleMesh{
		geometry3D: geometry3D{
			geometry: geometry{
				name:      name,
				geomType:  MeshType,
				dimension: 3,
				valid:     false,
			},
			vertices: vertices,
		},
	}

	mesh.valid = true
	return mesh
}
