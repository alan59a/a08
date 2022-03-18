package main

import (
	"log"

	"github.com/google/uuid"
)

// Universal Data Container as a Tensor
type Data struct {
	id   uuid.UUID
	name string
	raw  []float64
	dim  []int
	dimc []int
}

// Outputs a Data tensor with the given data.
// The excess data will be ignored and the deficients will be set to'0'.
// any non-float64 data will be converted to float64.
func New(data []float64, dimensions ...int) *Data {
	size := 1
	dimc := make([]int, len(dimensions))

	for a, b := range dimensions {
		size *= b

		if a == 0 {
			dimc[0] = 1
		} else if a != len(dimc)-1 {
			dimc[a+1] = size
		}

	}

	d := make([]float64, size)

	copy(d, data)

	return &Data{
		id:   uuid.New(),
		name: "",
		raw:  d,
		dim:  dimensions,
		dimc: dimc,
	}
}

// Sets an individual element of the tensor.
func (d *Data) Set(data float64, location ...int) {

	for a := range location {

		if location[a] >= d.dim[a] {
			log.Fatalln("Wrong location")
		}

	}

	loc := 0

	for a, b := range location {
		loc += b * d.dimc[a]
	}

	d.raw[loc] = data
}

// Outputs an individual element of the tensor.
func (d *Data) Get(location ...int) float64 {

	for a := range location {

		if location[a] >= d.dim[a] {
			log.Fatalln("Wrong location")
		}

	}

	loc := 0

	for a, b := range location {
		loc += b * d.dimc[a]
	}

	return d.raw[loc]
}

// Outputs the addition of 2 compatible data tensors.
func (d1 *Data) Add(d2 *Data) *Data {

	for a := range d1.dim {

		if d1.dim[a] != d2.dim[a] {
			log.Fatalln("Incompatible data")
		}

	}

	d := New(nil, d1.dim...)

	for a := range d.raw {
		d.raw[a] = d1.raw[a] + d2.raw[a]
	}

	return d
}

// Outputs subtraction of 2 compatible data tensors.
func (d1 *Data) Sub(d2 *Data) *Data {

	for a := range d1.dim {

		if d1.dim[a] != d2.dim[a] {
			log.Fatalln("Incompatible data")
		}

	}

	d := New(nil, d1.dim...)

	for a := range d.raw {
		d.raw[a] = d1.raw[a] - d2.raw[a]
	}

	return d
}

// Outputs the element-wise multiplication of 2 compatible tensors.
func (d1 *Data) Mul(d2 *Data) *Data {

	for a := range d1.dim {

		if d1.dim[a] != d2.dim[a] {
			log.Fatalln("Incompatible data")
		}

	}

	d := New(nil, d1.dim...)

	for a := range d.raw {
		d.raw[a] = d1.raw[a] * d2.raw[a]
	}

	return d
}

// outputs the element-wise division of 2 compatible tensors.
func (d1 *Data) Div(d2 *Data) *Data {

	for a := range d1.dim {

		if d1.dim[a] != d2.dim[a] {
			log.Fatalln("Incompatible data")
		}

	}

	d := New(nil, d1.dim...)

	for a := range d.raw {
		d.raw[a] = d1.raw[a] / d2.raw[a]
	}

	return d
}

// Outputs the dot product of 2 compatible tensors.
func (d1 *Data) Dot(d2 *Data) *Data {

	for a := len(d1.dim) - 2; a > 0; a-- {

		if d1.dim[a] != d2.dim[a+1] {
			log.Fatalln("Incompatible data")
		}

	}

	dim := make([]int, len(d1.dim))

	for a := range dim {
		dim[a] = d1.dim[a]

		if a == len(d1.dim)-1 {
			dim[a] = d2.dim[len(d2.dim)-1]
		}

	}

	d := New(nil, dim...)

	// TO DO: indefinite loops ... how?

	return d
}

// Outputs the inverse of the data
func (d *Data) inv() {

	if d.dim[len(d.dim)-1] != d.dim[len(d.dim)-2] {
		log.Fatalln(" Incompatible Data")
	}
}

// Applies the structured function to all values within the Data.
// # No matter how many variables the first one will be treated as input value.
func (d *Data) Act(f func(vars ...float64) float64, vars ...float64) *Data {
	dd := New(nil, d.dim...)
	v := make([]float64, len(vars)+1)

	for a := range vars {
		v[a+1] = vars[a]
	}

	for a := range d.raw {
		v[0] = d.raw[a]
		dd.raw[a] = f(vars...)
	}

	return dd
}

// a special method to go down one dimension (largest one) and adding the corresponding values together.
// and the output is a lower-dimensional *Data.
// in case of a 1-dimensional data , a 1-dimenstional data with a single value will be returned.
func (d *Data) Sum() *Data {
	l := len(d.dim)

	if l == 1 {
		l++
	}

	dim := make([]int, l)

	for a := range dim {

		if l == 2 {
			dim[a] = 1
		} else {
			dim[a] = d.dim[a+1]
		}

	}

	dd := New(nil, dim...)

	for a := 0; a < d.dim[0]; a++ {

		for b := range dd.raw {
			dd.raw[b] += d.raw[a*d.dimc[0]+b]
		}

	}

	return dd
}

// a special method to go down one dimension (largest one) and cutting down the data.
// and the output is a lower-dimensional []*Data.
// in case of a 1-dimensional data , the programm will panic.
func (d *Data) Cut() []*Data {

	if len(d.dim) > 2 {
		log.Fatalln("Too small to cut")
	}

	dim := make([]int, len(d.dim)-1)

	for a := range dim {
		dim[a] = d.dim[a+1]
	}

	ddd := make([]*Data, d.dim[0])

	for a := 0; a < d.dim[0]; a++ {

		dd := New(nil, dim...)

		for b := range dd.raw {
			dd.raw[b] = d.raw[a*d.dimc[0]+b]
		}

		ddd[a] = dd
	}

	return ddd
}

// Sets the data name.
func (d *Data) SetName(name string) { d.name = name }

// Outputs the data name.
func (d *Data) GetName() string { return d.name }

// Outputs the data UUID as a string.
func (d *Data) GetID() string { return d.id.String() }

// Outputs the raw data as a slice of float64.
func (d *Data) GetRaw() []float64 { return d.raw }

// Outputs the data dimensions as a slice of int.
func (d *Data) GetDim() []int { return d.dim }
