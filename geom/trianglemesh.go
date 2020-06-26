package geom

import "github.com/go-gl/mathgl/mgl32"

type TriangleMesh interface {
	Geometry3D

	// Triangles returns the topology information
	Triangles() []Triangle

	// SetTriangles replaces all triangles
	SetTriangles(triangles []Triangle)
}

type Triangle [3]int

type triangleMesh struct {
	geometry3D

	triangles []Triangle
}

func NewTriangleMesh(name string, vertices []mgl32.Vec3, triangles []Triangle) TriangleMesh {
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
		triangles: triangles,
	}

	mesh.valid = true
	return mesh
}

func (g *triangleMesh) SetTriangles(triangles []Triangle) {
	g.triangles = triangles
}
func (g *triangleMesh) Triangles() []Triangle {
	return g.triangles
}
