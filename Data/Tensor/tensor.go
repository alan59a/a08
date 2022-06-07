// a pure Go data processing package
package Tensor

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
	dims int
	len  int
}

// a universal location identifier
// lower index, lower dimension
// e.g. : index 0 -> x (columns), 1 -> y (rows), 2 -> z, ...
type Loc struct {
	Coord []int
}

// Tensor methods with "_" at the end will output the results in-place

// Sets an individual element of the tensor.
// For user convenience, similar to python, negative location is also acceptable
// Hold your horses. I know your proud Go-Vein is popping ... it's just for noobs ... you should NOT use it.
func (t *Tensor) Set(data float64, location ...int) {

	if len(location) != t.dims {
		log.Fatalln("Specify the full location")
	}

	for a := range location {

		if location[a] < 0 {
			location[a] = t.dim[a] + location[a]
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

// Returns an individual element of the tensor
// For user convenience, similar to python, negative location is also acceptable
// Hold your horses. I know your proud Go-Vein is popping ... it's just for noobs ... you should NOT use it.
func (t *Tensor) Get(location ...int) float64 {

	if len(location) != len(t.dim) {
		log.Fatalln("Worng dimensions")
	}

	for a := range location {

		if location[a] < 0 {
			location[a] = t.dim[a] + location[a]
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

// Returns a Tensor Applied with the structured to all values within the Tensor
// # No matter how many variables the first one will be treated as input value and the rest as function parameters
// # I was even nice enough to accept an error output for your function .... where else have you seen that?
func Apply(t *Tensor, f func(...float64) (float64, error), vars ...float64) *Tensor {
	dd := New(nil, t.dim...)
	v := make([]float64, len(vars)+1)

	for a := range vars {
		v[a+1] = vars[a]
	}

	for a := range t.raw {
		v[0] = t.raw[a]
		b, err := f(v...)
		if err != nil {
			log.Fatalln(err)
		}
		dd.raw[a] = b
	}

	return dd
}

// Returns a Tensor Applied with the structured function to all values within the Tensor
// # No matter how many variables the first one will be treated as input value and the rest as function parameters
// # I was even nice enough to accept an error output for your function .... where else have you seen that?
func (t *Tensor) Apply(f func(...float64) (float64, error), vars ...float64) *Tensor {
	dd := New(nil, t.dim...)
	v := make([]float64, len(vars)+1)

	for a := range vars {
		v[a+1] = vars[a]
	}

	for a := range t.raw {
		v[0] = t.raw[a]
		b, err := f(v...)
		if err != nil {
			log.Fatalln(err)
		}
		dd.raw[a] = b
	}

	return dd
}

// Applies the structured function to all values within the Tensor
// # No matter how many variables the first one will be treated as input value and the rest as function parameters
// # I was even nice enough to accept an error output for your function .... where else have you seen that?
func (t *Tensor) Apply_(f func(...float64) (float64, error), vars ...float64) {
	v := make([]float64, len(vars)+1)

	for a := range vars {
		v[a+1] = vars[a]
	}

	for a := range t.raw {
		v[0] = t.raw[a]
		b, err := f(v...)
		if err != nil {
			log.Fatalln(err)
		}
		t.raw[a] = b
	}

}

// a special method to go down one dimension (largest one) and adding the corresponding values together
// and the output is a lower-dimensional *Tensor
// in case of a 1-dimensional data , you will end up with another 1-dimenstional data with a single value
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

// a Recursive form of .Sum() method down to the last element
func (t *Tensor) SumR() float64 {
	s := 0.0

	for a := range t.raw {
		s += t.raw[a]
	}

	return s
}

// Slices the highest dimension into seperate Tensors
// so it returns is a lower-dimension []*Tensors
// in case of a 1-dimensional data , the programm will panic
func (t *Tensor) Cut() []*Tensor {

	if t.dims < 2 {
		log.Fatalln("Too small to cut")
	}

	dim := make([]int, t.dims-1)
	copy(dim, t.dim)
	d := make([]*Tensor, t.dim[t.dims-1])

	for a := 0; a < t.dim[t.dims-1]; a++ {

		dd := New(nil, dim...)

		for b := range dd.raw {
			dd.raw[b] = t.raw[a*t.dimc[t.dims-1]+b]
		}

		d[a] = dd
	}

	return d
}

// Returns the Minimum value in the Tensor
func (t *Tensor) Min() float64 {
	min := t.raw[0]

	for a := range t.raw {

		if min > t.raw[a] {
			min = t.raw[a]
		}

	}

	return min
}

// Returns the Minimum value in the Tensor and its location
func (t *Tensor) MinLoc() (float64, *Loc) {
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

	return min, &Loc{
		Coord: dim,
	}
}

// Returns the index of Minimum value in the Tensor
func (t *Tensor) ArgMin() int {
	min := t.raw[0]
	arg := 0

	for a := range t.raw {

		if min > t.raw[a] {
			min = t.raw[a]
			arg = a
		}

	}

	return arg
}

// Returns the Maximum value in the Tensor
func (t *Tensor) Max() float64 {
	max := t.raw[0]

	for a := range t.raw {

		if max < t.raw[a] {
			max = t.raw[a]
		}

	}

	return max
}

// Returns the Maximum value in the Tensor and its location
func (t *Tensor) MaxLoc() (float64, *Loc) {
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

	return max, &Loc{
		Coord: dim,
	}
}

// Returns the index of Maximum value in the Tensor
func (t *Tensor) ArgMax() int {
	max := t.raw[0]
	arg := 0

	for a := range t.raw {

		if max < t.raw[a] {
			max = t.raw[a]
			arg = a
		}

	}

	return arg
}

// Returns the index of Soft Maximum value and its index in the Tensor
func (t *Tensor) SoftMax() (*Tensor, int) {
	st := t.Zero()
	max := math.Exp(t.raw[0])
	st.raw[0] = max
	arg := 0
	sum := max

	for a := 1; a < t.len; a++ {
		st.raw[a] = math.Exp(t.raw[a])
		sum += st.raw[a]

		if max < st.raw[a] {
			max = st.raw[a]
			arg = a
		}

	}

	for a := range st.raw {
		st.raw[a] /= sum
	}

	return st, arg
}

// Returns the Mean of values stored in Tensor
func (t *Tensor) Mean() float64 {
	sum := 0.0

	for _, a := range t.raw {
		sum += a
	}

	return sum / float64(t.len)
}

// Returns the Standard Deviation of values stored in Tensor
func (t *Tensor) SD() float64 {
	sum := 0.0
	va := 0.0

	for _, a := range t.raw {
		sum += a
	}

	mean := sum / float64(t.len)

	for _, a := range t.raw {
		va += math.Pow(mean-a, 2)
	}

	return math.Sqrt(va / float64(t.len))
}

// Re-structures the dimensions of the data
// Note: This is a forceful actions and excess data will be removed
// as well as smaller data size will result in
func (t *Tensor) Reshape(dimensions ...int) {
	size := 1
	dim := make([]int, len(dimensions))
	dimc := make([]int, len(dimensions))

	for a, b := range dimensions {

		if len(dimensions) == 0 {
			dim = append(dim, 1)
		} else {

			if b < 1 {
				log.Fatalln("Bad dimensions")
			} else {
				dim[a] = b
			}

		}

	}

	for a, b := range dim {
		dimc[a] = size
		size *= b
	}

	d := make([]float64, size)

	copy(d, t.raw)
	t.dim = dim
	t.dims = len(dim)
	t.dimc = dimc
}

// Sets the data name
func (t *Tensor) SetName(name string) { t.name = name }

// Returns the data name
func (t *Tensor) Name() string { return t.name }

// Returns the data UUID as a string
func (t *Tensor) ID() string { return t.id.String() }

// Returns the raw data as a slice of float64
func (t *Tensor) Raw() []float64 { return t.raw }

// Returns the data dimensions as a slice of int
func (t *Tensor) Size() []int { return t.dim }

// Returns the data length
func (t *Tensor) Len() int { return t.len }

// Returns the Tensor dimensions
func (t *Tensor) Dims() int { return t.dims }
