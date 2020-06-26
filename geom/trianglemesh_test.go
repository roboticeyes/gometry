package geom

import (
	"testing"

	"github.com/go-gl/mathgl/mgl32"
)

func TestCreateMesh(t *testing.T) {

	v := []mgl32.Vec3{
		mgl32.Vec3{0.0, 0.0, 0.0},
		mgl32.Vec3{1.0, 0.0, 0.0},
		mgl32.Vec3{0.0, 1.0, 0.0},
		mgl32.Vec3{0.0, 0.0, 1.0},
	}

	f := []Triangle{
		Triangle{0, 1, 2},
	}
	m := NewTriangleMesh("mesh-1", v, f)
	if !m.Valid() {
		t.Fatal("Mesh should be valid")
	}
	if m.Name() != "mesh-1" {
		t.Fatal("Mesh name does not match")
	}
	if m.Type() != MeshType {
		t.Fatal("Mesh type does not match")
	}
	if len(m.Vertices()) != 4 {
		t.Fatal("Vertices number does not match")
	}
}
