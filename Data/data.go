package Data

import (
	"log"

	"github.com/alan59a/a08/Data/Matrix"
	"github.com/alan59a/a08/Data/Tensor"
	"github.com/alan59a/a08/Data/Vector"
)

type Datas interface {
	Set(float64, ...int)
	Get(...int)
	Apply(f func(...float64) (float64, error), vars ...float64) *Datas
	Apply_(f func(...float64) (float64, error), vars ...float64)
	Sum()
	SumR() float64
	Cut() []*Datas
	Min() float64
	MinLoc() (float64, *Tensor.Loc)
	ArgMin() int
	Max() float64
	MaxLoc() (float64, *Tensor.Loc)
	ArgMax() int
	SoftMax()
	Mean()
	SD()
	Reshape()
	SetName(string)
	Name() string
	ID() string
	Raw() []float64
	Size() []int
	Len() int
}

func Tensor2Matrix(t *Tensor.Tensor) *Matrix.Matrix {

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

	return Matrix.New(d, dim[1], dim[0])
}

func Tensor2Vector(t *Tensor.Tensor) *Vector.Vector {

	if t.Dims() != 1 {
		log.Fatalln("Unacceptable dimensions")
	}

	raw := make([]float64, t.Len())
	copy(raw, t.Raw())

	return Vector.New(raw, t.Len())
}

func Matrix2Tensor(m *Matrix.Matrix) *Tensor.Tensor {
	a, b := m.Size()
	raw := m.Raw()

	return Tensor.New(raw, a, b)
}

func Matrix2Vector(m *Matrix.Matrix) *Vector.Vector {
	a, b := m.Size()

	if a != 1 && b != 1 {
		log.Fatalln("Unacceptable dimensions")
	}

	raw := m.Raw()

	return Vector.New(raw, a*b)
}

func Vector2Tensro(v *Vector.Vector) *Tensor.Tensor { return Tensor.New(v.Raw(), v.Len()) }

func Vector2Matrix(v *Vector.Vector) *Matrix.Matrix {
	return Matrix.New([][]float64{v.Raw()}, 1, v.Len())
}

func Save(d *Datas) {

}
