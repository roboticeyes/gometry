# gometry

**gometry** is a GO library which offers geometry processing algorithms.

## Requirements

* Lean and simple
* Little dependencies
* Simple architecture supporting various geometries (images, octree, pointclouds, ...)

## Features

### Meshing

* Mesh cleanup (duplicate vertices)
* Mesh merging
* Mesh simplification
* Mesh smoothing

### IO

* Import/export REXfile
* Export to OBJ/...

### Transformation

* Translate
* Scale
* Rotate

### Finding

* Find nearest vertex
* Find neighboring/connected vertices based on a given radius
* Calculate normal vector based on neighbors

### CAD

* Generate BIM model
* Find sharp corners

## Planned Dependencies

* Numeric library: [gonum](https://github.com/gonum/gonum)
* REXfile IO: [gorexfile](https://github.com/roboticeyes/gorexfile)
* Visualization [g3n](https://github.com/g3n/engine)

