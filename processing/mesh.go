package processing

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/kyroy/kdtree"
	"github.com/kyroy/kdtree/points"
	"github.com/roboticeyes/gometry/geom"
)

// RemoveDuplicates removes vertex duplicates of a triangle mesh. If two vertices are closer than
// the given distanceThreshold, then the vertices are merged
func RemoveDuplicates(mesh geom.TriangleMesh, distanceThreshold float64) error {

	vertices := []mgl32.Vec3{}
	triangles := []geom.Triangle{}

	// create kdtree with first point
	tree := kdtree.New([]kdtree.Point{
		points.NewPoint([]float64{
			float64(mesh.Vertex(0).X()),
			float64(mesh.Vertex(0).Y()),
			float64(mesh.Vertex(0).Z()),
		}, 0 /* index in vertices slice */),
	})
	vertices = append(vertices, mesh.Vertex(0))

	// 1. Traverse all vertices and add only those vertices which seem to be duplicate because the
	// dictance to the found point in the KDtree is smaller than an epsilon
	for i := 1; i < len(mesh.Vertices()); i++ {
		vtx := mesh.Vertex(i)
		nearest := tree.KNN(&points.Point{
			Coordinates: []float64{
				float64(vtx.X()),
				float64(vtx.Y()),
				float64(vtx.Z()),
			},
		}, 1)
		if len(nearest) > 0 {
			// fmt.Println("Found nearest (search - treenode) ", vtx, nearest[0])
			dist := vtx.Sub(mgl32.Vec3{
				float32(nearest[0].Dimension(0)),
				float32(nearest[0].Dimension(1)),
				float32(nearest[0].Dimension(2)),
			}).Len()
			if math.Abs(float64(dist)) > distanceThreshold {
				vertices = append(vertices, vtx)
				tree.Insert(
					points.NewPoint([]float64{
						float64(vtx.X()),
						float64(vtx.Y()),
						float64(vtx.Z()),
					}, len(vertices)-1))
			} else {
				// ignore, since those points are already in the list
			}
		} else {
			// should not happen, just failsafe
			vertices = append(vertices, vtx)
			tree.Insert(
				points.NewPoint([]float64{
					float64(vtx.X()),
					float64(vtx.Y()),
					float64(vtx.Z()),
				}, len(vertices)-1))
		}
	}

	// 2. Walk through triangle list and lookup every vertex in the kd tree
	for i := 0; i < len(mesh.Triangles()); i++ {
		triangle := mesh.Triangles()[i]
		for j := 0; j < 3; j++ {
			vtx := mesh.Vertex(triangle[j])
			nearest := tree.KNN(&points.Point{
				Coordinates: []float64{
					float64(vtx.X()),
					float64(vtx.Y()),
					float64(vtx.Z()),
				},
			}, 1)
			if len(nearest) > 0 {
				pt := nearest[0].(*points.Point)
				triangle[j] = pt.Data.(int)
			}
		}
		triangles = append(triangles, triangle)
	}
	mesh.SetVertices(vertices)
	mesh.SetTriangles(triangles)

	return nil
}
