package processing

import (
	"os"
	"testing"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/roboticeyes/gometry/geom"
	"github.com/roboticeyes/gometry/io"
)

var (
	v0 = mgl32.Vec3{0, 0, 0}
	v1 = mgl32.Vec3{1, 0, 0}
	v2 = mgl32.Vec3{2, 0, 0}
	v3 = mgl32.Vec3{0, 1, 0}
	v4 = mgl32.Vec3{1, 1, 0}
	v5 = mgl32.Vec3{2, 1, 0}
	v6 = mgl32.Vec3{0, 2, 0}
	v7 = mgl32.Vec3{1, 2, 0}
	v8 = mgl32.Vec3{2, 2, 0}
)

func TestRemoveDuplicates(t *testing.T) {

	v := []mgl32.Vec3{
		v0, v1, v4,
		v0, v4, v3,
		v1, v2, v5,
		v1, v5, v4,
		v3, v4, v7,
		v3, v7, v6,
		v4, v5, v8,
		v4, v8, v7,
	}

	var f []geom.Triangle

	for i := 0; i < 24; i += 3 {
		f = append(f, geom.Triangle{i, i + 1, i + 2})
	}
	m := geom.NewTriangleMesh("mesh", v, f)
	if len(m.Vertices()) != 24 {
		t.Fatal("Vertices number does not match")
	}
	if len(m.Triangles()) != 8 {
		t.Fatal("Vertices number does not match")
	}
	w, _ := os.Create("/tmp/mesh.obj")
	io.WriteTriangleMeshToObj(m, w)

	RemoveDuplicates(m, 0.001)

	if len(m.Vertices()) != 9 {
		t.Fatal("Vertices number after removal does not match")
	}
	w2, _ := os.Create("/tmp/mesh-cleaned.obj")
	io.WriteTriangleMeshToObj(m, w2)

}
