package Vector

import (
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

// Returns a Vector with specified dimensions and data
// Note: The excess data will be ignored and if insuffcient, will be set to'0'
// and any dimension size < 1 is unacceptable
func New(data []float64, dimensions ...int) *Vector {
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

	copy(d, data)

	return &Vector{
		id:   uuid.New(),
		name: "",
		data: d,
		len:  size,
	}
}

// Returns a Vector with specified dimensions and value of '0'
func Zero(dimensions ...int) *Vector {
	size := 1
	dim := make([]int, len(dimensions))
	dimc := make([]int, len(dimensions))

	for a, b := range dimensions {

		if dimensions == nil {
			log.Fatalln("Provide dimensions")
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

	return &Vector{
		id:   uuid.New(),
		name: "",
		data: d,
		len:  size,
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
func (v *Vector) Zero_(dimensions ...int) {

	if dimensions == nil {

		for a := range v.data {
			v.data[a] = 0
		}

		return
	}

	size := 1
	dim := make([]int, len(dimensions))
	dimc := make([]int, len(dimensions))

	for a, b := range dimensions {

		if b < 1 {
			log.Fatalln("Bad dimensions")
		} else {
			dim[a] = b
		}

	}

	for a, b := range dim {
		dimc[a] = size
		size *= b
	}

	d := make([]float64, size)

	v.data = d
	v.len = size
}

// Returns a Vector with specified dimesnions and values of '1'
func Ones(dimensions ...int) *Vector {
	size := 1
	dim := make([]int, len(dimensions))
	dimc := make([]int, len(dimensions))

	for a, b := range dimensions {

		if dimensions == nil {
			log.Fatalln("Provide dimensions")
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

	for a := range d {
		d[a] = 1
	}

	return &Vector{
		id:   uuid.New(),
		name: "",
		data: d,
		len:  size,
	}
}

// Returns a Vector retaining the dimensions of the provided Vector and with value of '1'
func (v *Vector) Ones() *Vector {
	d := make([]float64, len(v.data))

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
func (v *Vector) Ones_(dimensions ...int) {

	if dimensions == nil {

		for a := range v.data {
			v.data[a] = 1
		}

		return
	}

	size := 1
	dim := make([]int, len(dimensions))
	dimc := make([]int, len(dimensions))

	for a, b := range dimensions {

		if b < 1 {
			log.Fatalln("Bad dimensions")
		} else {
			dim[a] = b
		}

	}

	for a, b := range dim {
		dimc[a] = size
		size *= b
	}

	d := make([]float64, size)

	for a := range d {
		v.data[a] = 1
	}

	v.data = d
	v.len = size
}

// Returns a Vector with specified dimensions and random values
// TODO: Normal randomization, Mean = 0, SD = 1 --> DONE
func Random(dimensions ...int) *Vector {
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

	rand.Seed(time.Now().UnixNano())

	for a := range d {
		d[a] = rand.NormFloat64()
	}

	return &Vector{
		id:   uuid.New(),
		name: "",
		data: d,
		len:  size,
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
func (v *Vector) Random_(dimensions ...int) {
	rand.Seed(time.Now().UnixNano())

	if dimensions == nil {

		for a := range v.data {
			v.data[a] = rand.NormFloat64()
		}

		return
	}

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

	for a := range d {
		v.data[a] = rand.NormFloat64()
	}

	v.data = d
	v.len = size
}
