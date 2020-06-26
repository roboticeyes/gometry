package io

import (
	"bufio"
	"fmt"
	"io"

	"github.com/roboticeyes/gometry/geom"
)

func WriteTriangleMeshToObj(t geom.TriangleMesh, output io.Writer) error {

	w := bufio.NewWriter(output)
	w.WriteString("# gometry export, (c) Robotic Eyes\n")

	for _, v := range t.Vertices() {
		w.WriteString(fmt.Sprintf("v %f %f %f\n", v.X(), v.Y(), v.Z()))
	}

	for _, f := range t.Triangles() {
		w.WriteString(fmt.Sprintf("f %d %d %d\n", f[0]+1, f[1]+1, f[2]+1))
	}
	w.Flush()
	return nil
}
