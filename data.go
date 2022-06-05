// a pure Go data processing package
package main

import (
	"log"
	"math"

	"github.com/google/uuid"
)

// Universal Data Container as a Tensor
type Tensor struct {
	id   uuid.UUID
	name string
	raw  []float64
	dim  []int
	dimc []int
}

// Tensor methods with "_" at the end will output the results in-place

// Data Frame
type Frame struct {
	Tensor []*Tensor
}

// Sets an individual element of the tensor.
// For user convenience, similar to python, negative location is also acceptable
// Hold your horses. I know your proud Go-Vein is popping ... it's just for noobs ... you should NOT use it.
func (t *Tensor) Set(data float64, location ...int) {

	for a := range location {

		if location[a] < 0 {
			location[a] = t.dim[a] - location[a]
		}

		if location[a] >= t.dim[a] || location[a] < 0 {
			log.Fatalln("Wrong location")
		}

	}

	loc := 0

	for a, b := range location {
		loc += b * t.dimc[a]
	}

	t.raw[loc] = data
}

// Outputs an individual element of the tensor
// For user convenience, similar to python, negative location is also acceptable
// Hold your horses. I know your proud Go-Vein is popping ... it's just for noobs ... you should NOT use it.
func (t *Tensor) Get(location ...int) float64 {

	for a := range location {

		if location[a] < 0 {
			location[a] = t.dim[a] - location[a]
		}

		if location[a] >= t.dim[a] || location[a] < 0 {
			log.Fatalln("Wrong location")
		}

	}

	loc := 0

	for a, b := range location {
		loc += b * t.dimc[a]
	}

	return t.raw[loc]
}

// Applies the structured function to all values within the Tensor
// # No matter how many variables the first one will be treated as input value and the rest as function parameters.
func (t *Tensor) Act(f func(vars ...float64) float64, vars ...float64) *Tensor {
	dd := New(nil, t.dim...)
	v := make([]float64, len(vars)+1)

	for a := range vars {
		v[a+1] = vars[a]
	}

	for a := range t.raw {
		v[0] = t.raw[a]
		dd.raw[a] = f(vars...)
	}

	return dd
}

// a special method to go down one dimension (largest one) and adding the corresponding values together
// and the output is a lower-dimensional *Tensor.
// in case of a 1-dimensional data , a 1-dimenstional data with a single value will be returned.
func (t *Tensor) Sum() *Tensor {
	l := len(t.dim)

	if l == 1 {
		l++
	}

	dim := make([]int, l)

	for a := range dim {

		if l == 2 {
			dim[a] = 1
		} else {
			dim[a] = t.dim[a+1]
		}

	}

	dd := New(nil, dim...)

	for a := 0; a < t.dim[0]; a++ {

		for b := range dd.raw {
			dd.raw[b] += t.raw[a*t.dimc[0]+b]
		}

	}

	return dd
}

// a Recursive form of Sum() method
func (t *Tensor) SumR() float64 {
	s := 0.0

	for a := range t.raw {
		s += t.raw[a]
	}

	return s
}

// a special method to go down one dimension (largest one) and cutting down the data
// and the output is a lower-dimensional []*Tensor.
// in case of a 1-dimensional data , the programm will panic.
func (t *Tensor) Cut() []*Tensor {

	if len(t.dim) > 2 {
		log.Fatalln("Too small to cut")
	}

	dim := make([]int, len(t.dim)-1)

	for a := range dim {
		dim[a] = t.dim[a+1]
	}

	ddd := make([]*Tensor, t.dim[0])

	for a := 0; a < t.dim[0]; a++ {

		dd := New(nil, dim...)

		for b := range dd.raw {
			dd.raw[b] = t.raw[a*t.dimc[0]+b]
		}

		ddd[a] = dd
	}

	return ddd
}

// Outputs the Minimum value in the Tensor
func (t *Tensor) Min() float64 {
	min := t.raw[0]

	for a := range t.raw {

		if min > t.raw[a] {
			min = t.raw[a]
		}

	}

	return min
}

// Outputs the Maximum value in the Tensor
func (t *Tensor) Max() float64 {
	max := t.raw[0]

	for a := range t.raw {

		if max < t.raw[a] {
			max = t.raw[a]
		}

	}

	return max
}

// Outputs the Minimum value in the Tensor and its location
func (t *Tensor) MinL() (float64, []int) {
	min := t.raw[0]
	loc := 0
	dim := make([]int, len(t.dim))

	for a := range t.raw {

		if min > t.raw[a] {
			min = t.raw[a]
			loc = a
		}

	}

	for a := range t.dimc {
		dim[a] = loc / t.dimc[a]
		loc = loc % t.dimc[a]
	}

	return min, dim
}

// Outputs the Maximum value in the Tensor and its location
func (t *Tensor) MaxL() (float64, []int) {
	max := t.raw[0]
	loc := 0
	dim := make([]int, len(t.dim))

	for a := range t.raw {

		if max < t.raw[a] {
			max = t.raw[a]
			loc = a
		}

	}

	for a := range t.dimc {
		dim[a] = loc / t.dimc[a]
		loc = loc % t.dimc[a]
	}

	return max, dim
}

// Outputs the Mean of values stored in Tensor
func (t *Tensor) Mean() float64 {
	sum := 0.0

	for _, a := range t.raw {
		sum += a
	}

	return sum / float64(len(t.raw))
}

// Outputs the Standard Deviation of values stored in Tensor
func (t *Tensor) SD() float64 {
	sum := 0.0
	va := 0.0

	for _, a := range t.raw {
		sum += a
	}

	mean := sum / float64(len(t.raw))

	for _, a := range t.raw {
		va += math.Pow(mean-a, 2)
	}

	return math.Sqrt(va / float64(len(t.raw)))
}

// Sets the data name
func (t *Tensor) SetName(name string) { t.name = name }

// Outputs the data name
func (t *Tensor) GetName() string { return t.name }

// Outputs the data UUID as a string
func (t *Tensor) GetID() string { return t.id.String() }

// Outputs the raw data as a slice of float64
func (t *Tensor) GetRaw() []float64 { return t.raw }

// Outputs the data dimensions as a slice of int
func (t *Tensor) GetDim() []int { return t.dim }
