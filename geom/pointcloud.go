package geom

type PointCloud interface {
	Geometry3D
}

type pointCloud struct {
	geometry
}
