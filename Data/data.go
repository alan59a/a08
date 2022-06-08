package Data

import (
	"log"
)

func (t *Tensor) Matrix() *Matrix {

	if t.Dims() != 2 {
		log.Fatalln("Unacceptable dimensions")
	}

	dim := t.Size()

	d := make([][]float64, dim[1])
	raw := make([]float64, dim[0]*dim[1])
	copy(raw, t.Raw())

	for a := range d {
		d[a] = make([]float64, dim[0])

		for b := range d[0] {
			d[a][b] = raw[a*dim[0]+b]
		}

	}

	return NewMatrix(d, dim[1], dim[0])
}

func (t *Tensor) Vector() *Vector {

	if t.Dims() != 1 {
		log.Fatalln("Unacceptable dimensions")
	}

	raw := make([]float64, t.Len())
	copy(raw, t.Raw())

	return NewVector(raw, t.Len())
}

func (m *Matrix) Tensor() *Tensor {
	a, b := m.Size()
	raw := m.Raw()

	return NewTensor(raw, a, b)
}

func (m *Matrix) Vector() *Vector {
	a, b := m.Size()

	if a != 1 && b != 1 {
		log.Fatalln("Unacceptable dimensions")
	}

	raw := m.Raw()

	return NewVector(raw, a*b)
}

func (v *Vector) Tensro() *Tensor { return NewTensor(v.Raw(), v.Len()) }

func (v *Vector) Matrix() *Matrix {
	return NewMatrix([][]float64{v.Raw()}, 1, v.Len())
}
