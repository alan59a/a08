package Data

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

// Returns a Vector with specified dimensions and data
// Note: The excess data will be ignored and if insuffcient, will be set to'0'
// and any dimension size < 1 is unacceptable
func NewVector(data []float64, length int) *Vector {
	d := make([]float64, length)
	copy(d, data)

	return &Vector{
		id:   uuid.New(),
		name: "",
		data: d,
		len:  length,
	}
}

// Returns a Vector retaining the dimensions of the provided Vector and with value of '0'
func (v *Vector) Zero() *Vector {
	d := make([]float64, v.len)

	return &Vector{
		id:   uuid.New(),
		name: "",
		data: d,
		len:  v.len,
	}
}

// Repopulate a Vector and with value of '0' , retaining the dimensions or using the provided dimensions
// NOTE: NOT porviding dimensions will result in using the same dimensions of the Vector provided
func (v *Vector) Zero_(length int) {
	d := make([]float64, length)

	v.data = d
	v.len = length
}

// Returns a Vector with specified dimesnions and values of '1'
func OnesVector(length int) *Vector {
	d := make([]float64, length)

	for a := range d {
		d[a] = 1
	}

	return &Vector{
		id:   uuid.New(),
		name: "",
		data: d,
		len:  length,
	}
}

// Returns a Vector retaining the dimensions of the provided Vector and with value of '1'
func (v *Vector) Ones() *Vector {
	d := make([]float64, v.len)

	for a := range d {
		d[a] = 1
	}

	return &Vector{
		id:   uuid.New(),
		name: "",
		data: d,
		len:  v.len,
	}
}

// Repopulate a Vector and with value of '1' , retaining the dimensions or using the provided dimensions
// NOTE: NOT porviding dimensions will result in using the same dimensions of the Vector provided
func (v *Vector) Ones_(length int) {
	d := make([]float64, length)

	for a := range d {
		v.data[a] = 1
	}

	v.data = d
	v.len = length
}

// Returns a Vector with specified dimensions and random values
// TODO: Normal randomization, Mean = 0, SD = 1 --> DONE
func RandomVector(length int) *Vector {
	d := make([]float64, length)

	rand.Seed(time.Now().UnixNano())

	for a := range d {
		d[a] = rand.NormFloat64()
	}

	return &Vector{
		id:   uuid.New(),
		name: "",
		data: d,
		len:  length,
	}
}

// Returns a Vector retaining the dimensions of the provided Vector and with random values
// TODO: Normal randomization, Mean = 0, SD = 1 --> DONE
func (v *Vector) Random() *Vector {
	d := make([]float64, v.len)

	rand.Seed(time.Now().UnixNano())

	for a := range d {
		d[a] = rand.NormFloat64()
	}

	return &Vector{
		id:   uuid.New(),
		name: "",
		data: d,
		len:  v.len,
	}
}

// Repopulate a Vector and with random values , retaining the dimensions or using the provided dimensions
// NOTE: NOT porviding dimensions will result in using the same dimensions of the Vector provided
// TODO: Normal randomization, Mean = 0, SD = 1 --> DONE
func (v *Vector) Random_(length int) {
	rand.Seed(time.Now().UnixNano())
	d := make([]float64, length)

	for a := range d {
		v.data[a] = rand.NormFloat64()
	}

	v.data = d
	v.len = length
}

// Return a replicate of the provided Vector
func (v *Vector) Clone() *Vector {
	d := make([]float64, v.len)

	for a := range d {
		d[a] = v.data[a]
	}

	return &Vector{
		id:   uuid.New(),
		name: v.name,
		data: d,
		len:  v.len,
	}
}
