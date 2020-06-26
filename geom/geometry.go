package geom

type GeometryType int

const (
	LineSetType = iota
	TextType
	PointCloudType
	MeshType
	ImageType
)

// Geometry is the base interface for any geometry object
type Geometry interface {
	Name() string
	SetName(name string)
	Valid() bool
	Type() GeometryType
	Dimension() int
}

type geometry struct {
	geomType  GeometryType
	dimension int
	name      string
	valid     bool
}

func (g *geometry) SetName(name string) {
	g.name = name
}

func (g *geometry) Name() string {
	return g.name
}

func (g *geometry) Type() GeometryType {
	return g.geomType
}

func (g *geometry) Dimension() int {
	return g.dimension
}

func (g *geometry) Valid() bool {
	return g.valid
}
