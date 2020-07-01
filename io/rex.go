package io

import (
	"bytes"
	"fmt"
	"io"

	"github.com/roboticeyes/gometry/geom"
	"github.com/roboticeyes/gorexfile/encoding/rexfile"
)

func ReadRexFile(r io.Reader) []geom.TriangleMesh {

	d := rexfile.NewDecoder(r)
	_, rex, err := d.Decode()

	if err != nil {
		return []geom.TriangleMesh{}
	}
	if rex == nil {
		return []geom.TriangleMesh{}
	}

	fmt.Printf("Found %d meshes\n", len(rex.Meshes))

	meshes := []geom.TriangleMesh{}
	for _, m := range rex.Meshes {
		mesh := geom.NewTriangleMesh(m.Name, m.Coords, rex2model(m.Triangles))
		meshes = append(meshes, mesh)
	}
	return meshes
}

func WriteTriangleMeshesToRex(meshes []geom.TriangleMesh, output io.Writer) error {

	rexFile := rexfile.File{}
	id := uint64(1)

	rexFile.Materials = append(rexFile.Materials, rexfile.NewMaterial(0))

	for _, m := range meshes {
		rexFile.Meshes = append(rexFile.Meshes, rexfile.Mesh{
			ID:         id,
			Name:       m.Name(),
			Coords:     m.Vertices(),
			Triangles:  model2rex(m.Triangles()),
			MaterialID: 0,
		})
		id++
	}

	var buf bytes.Buffer
	e := rexfile.NewEncoder(&buf)
	err := e.Encode(rexFile)
	if err != nil {
		fmt.Println(err)
	}

	_, err = output.Write(buf.Bytes())
	return err
}

func model2rex(input []geom.Triangle) []rexfile.Triangle {

	triangles := make([]rexfile.Triangle, len(input))

	for i, t := range input {
		triangles[i] = rexfile.Triangle{
			V0: uint32(t[0]),
			V1: uint32(t[1]),
			V2: uint32(t[2]),
		}
	}
	return triangles
}

func rex2model(input []rexfile.Triangle) []geom.Triangle {

	triangles := make([]geom.Triangle, len(input))

	for i, t := range input {
		triangles[i] = geom.Triangle{int(t.V0), int(t.V1), int(t.V2)}
	}
	return triangles
}
