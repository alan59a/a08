package main

/*
import (
	"log"

	"github.com/alan59a/a08/Matrix"
	"github.com/alan59a/a08/Tensor"
	"github.com/google/uuid"
)

// Returns a Matrix of acceptable (2D) Tensor
func (t *Tensor.Tensor) Matrix() *Matrix.Matrix {

	if t.dims != 2 {
		log.Fatalln("Unacceptable dimensions")
	}

	d := make([][]float64, t.dim[1])

	for a := range d {
		d[a] = make([]float64, t.dim[0])

		for b := range d[0] {
			d[a][b] = t.raw[a*t.dimc[1]+b]
		}

	}

	return &Matrix.Matrix{
		id:   uuid.New(),
		name: t.name,
		data: d,
		len:  t.len,
		col:  t.dim[0],
		row:  t.dim[1],
	}
}
*/
