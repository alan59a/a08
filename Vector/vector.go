package Vector

import (
	"log"
	"math"

	"github.com/google/uuid"
)

// Universal Data Container as a Tensor
type Vector struct {
	id   uuid.UUID
	name string
	data []float64
	len  int
}

// a universal location identifier
// lower index, lower dimension
// e.g. : index 0 -> x (columns), 1 -> y (rows), 2 -> z, ...
type Loc struct {
	Coord []int
}

// Vector methods with "_" at the end will output the results in-place

// Sets an individual element of the Vector.
// For user convenience, similar to python, negative location is also acceptable
// Hold your horses. I know your proud Go-Vein is popping ... it's just for noobs ... you should NOT use it.
func (v *Vector) Set(data float64, i int) {

	if i < 0 {
		i = v.len + i
	}

	if i >= v.len || i < 0 {
		log.Fatalln("Wrong location")
	}

	v.data[i] = data
}

// Returns an individual element of the Vector
// For user convenience, similar to python, negative location is also acceptable
// Hold your horses. I know your proud Go-Vein is popping ... it's just for noobs ... you should NOT use it.
func (v *Vector) Get(i int) float64 {

	if i < 0 {
		i = v.len + i
	}

	if i >= v.len || i < 0 {
		log.Fatalln("Wrong location")
	}

	return v.data[i]
}

// Applies the structured function to all values within the Vector
// # No matter how many variables the first one will be treated as input value and the rest as function parameters
// # I was even nice enough to accept an error output for your function .... where else have you seen that?
func (v *Vector) Apply(f func(...float64) (float64, error), vars ...float64) *Vector {
	d := make([]float64, v.len)
	vv := make([]float64, len(vars)+1)

	for a := range vars {
		vv[a+1] = vars[a]
	}

	for a := range v.data {

		vv[0] = v.data[a]
		c, err := f(vv...)

		if err != nil {
			log.Fatalln(err)
		}

		d[a] = c

	}

	return &Vector{
		id:   uuid.New(),
		name: "",
		data: d,
		len:  v.len,
	}
}

// Returns the Minimum value in the Vector
func (v *Vector) Min() float64 {
	min := v.data[0]

	for a := range v.data {

		if min < v.data[a] {
			min = v.data[a]
		}

	}

	return min
}

// Returns the Minimum value in the Vector and its location
func (v *Vector) MinLoc() (float64, *Loc) {
	min := v.data[0]
	i := 0

	for a := range v.data {

		if min < v.data[a] {
			min = v.data[a]
			i = a
		}

	}

	return min, &Loc{
		Coord: []int{i},
	}
}

// Returns the index of Minimum value in the Vector
func (v *Vector) ArgMin() *Loc {
	min := v.data[0]
	i := 0

	for a := range v.data {

		if min > v.data[a] {
			min = v.data[a]
			i = a
		}

	}

	return &Loc{
		Coord: []int{i},
	}
}

// Returns the Maximum value in the Vector
func (v *Vector) Max() float64 {
	max := v.data[0]

	for a := range v.data {

		if max < v.data[a] {
			max = v.data[a]
		}

	}

	return max
}

// Returns the Maximum value in the Vector and its location
func (v *Vector) MaxLoc() (float64, *Loc) {
	max := v.data[0]
	i := 0

	for a := range v.data {

		if max < v.data[a] {
			max = v.data[a]
			i = a
		}

	}

	return max, &Loc{
		Coord: []int{i},
	}
}

// Returns the index of Maximum value in the Vector
func (v *Vector) ArgMax() *Loc {
	max := v.data[0]
	i := 0

	for a := range v.data {

		if max < v.data[a] {
			max = v.data[a]
			i = a
		}

	}

	return &Loc{
		Coord: []int{i},
	}
}

// Returns the index of Soft Maximum value and its index in the Vector
func (v *Vector) SoftMax() (*Vector, *Loc) {
	st := v.Zero()
	max := math.Exp(v.data[0])
	st.data[0] = max
	i := 0
	sum := max

	for a := range v.data {

		st.data[a] = math.Exp(v.data[a])
		sum += st.data[a]

		if max < st.data[a] {
			max = st.data[a]
			i = a
		}

	}

	for a := range st.data {
		st.data[a] /= sum
	}

	return st, &Loc{
		Coord: []int{i},
	}
}

// Returns the Mean of values stored in Vector
func (v *Vector) Mean() float64 {
	sum := 0.0

	for _, a := range v.data {
		sum += a
	}

	return sum / float64(v.len)
}

// Returns the Standard Deviation of values stored in Vector
func (v *Vector) SD() float64 {
	sum := 0.0
	va := 0.0

	for _, a := range v.data {
		sum += a
	}

	mean := sum / float64(v.len)

	for _, a := range v.data {
		va += math.Pow(mean-a, 2)
	}

	return math.Sqrt(va / float64(v.len))
}

// Re-structures the dimensions of the data
// Note: This is a forceful actions and excess data will be removed
func (v *Vector) Reshape_(len int) {

	if len < 1 {
		log.Fatalln("Bad dimensions")
	}

	v.len = len
	d := make([]float64, len)

	for a := range d {

		if a < v.len {
			d[a] = v.data[a]
		}

	}
	v.data = d
	v.len = len
}

// Sets the data name
func (v *Vector) SetName(name string) { v.name = name }

// Returns the data name
func (v *Vector) Name() string { return v.name }

// Returns the data UUID as a string
func (v *Vector) ID() string { return v.id.String() }

// Returns the data data as a slice of float64
func (v *Vector) Raw() []float64 { return v.data }

// Returns the data dimensions as a slice of int
func (v *Vector) Size() (len int) { return v.len }

// Returns the data length
func (v *Vector) Len() int { return v.len }
