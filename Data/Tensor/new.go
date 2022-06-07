package Tensor

import (
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

// Returns a Tensor with specified dimensions and data
// Note: The excess data will be ignored and if insuffcient, will be set to'0'
// and any dimension size < 1 is unacceptable
func New(data []float64, dimensions ...int) *Tensor {
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

	return &Tensor{
		id:   uuid.New(),
		name: "",
		raw:  d,
		len:  size,
		dim:  dim,
		dimc: dimc,
	}
}

// Returns a Tensor with specified dimensions and value of '0'
func Zero(dimensions ...int) *Tensor {
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

	return &Tensor{
		id:   uuid.New(),
		name: "",
		raw:  d,
		len:  size,
		dim:  dim,
		dimc: dimc,
	}
}

// Returns a Tensor retaining the dimensions of the provided Tensor and with value of '0'
func (t *Tensor) Zero() *Tensor {
	dim := make([]int, 0)
	dimc := make([]int, len(dim))
	copy(dim, t.dim)
	copy(dimc, t.dimc)
	d := make([]float64, t.len)

	return &Tensor{
		id:   uuid.New(),
		name: "",
		raw:  d,
		len:  t.len,
		dim:  dim,
		dimc: dimc,
	}
}

// Repopulate a Tensor and with value of '0' , retaining the dimensions or using the provided dimensions
// NOTE: NOT porviding dimensions will result in using the same dimensions of the Tensor provided
func (t *Tensor) Zero_(dimensions ...int) {

	if dimensions == nil {

		for a := range t.raw {
			t.raw[a] = 0
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

	t.raw = d
	t.dim = dim
	t.dimc = dimc
	t.len = size
}

// Returns a Tensor with specified dimesnions and values of '1'
func Ones(dimensions ...int) *Tensor {
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

	return &Tensor{
		id:   uuid.New(),
		name: "",
		raw:  d,
		len:  size,
		dim:  dim,
		dimc: dimc,
	}
}

// Returns a Tensor retaining the dimensions of the provided Tensor and with value of '1'
func (t *Tensor) Ones() *Tensor {
	dim := make([]int, 0)
	dimc := make([]int, len(dim))
	copy(dim, t.dim)
	copy(dimc, t.dimc)
	d := make([]float64, len(t.raw))

	for a := range d {
		d[a] = 1
	}

	return &Tensor{
		id:   uuid.New(),
		name: "",
		raw:  d,
		len:  t.len,
		dim:  dim,
		dimc: dimc,
	}
}

// Repopulate a Tensor and with value of '1' , retaining the dimensions or using the provided dimensions
// NOTE: NOT porviding dimensions will result in using the same dimensions of the Tensor provided
func (t *Tensor) Ones_(dimensions ...int) {

	if dimensions == nil {

		for a := range t.raw {
			t.raw[a] = 1
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
		t.raw[a] = 1
	}

	t.raw = d
	t.dim = dim
	t.dimc = dimc
	t.len = size
}

// Returns a Tensor with specified dimensions and random values
// TODO: Normal randomization, Mean = 0, SD = 1 --> DONE
func Random(dimensions ...int) *Tensor {
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

	return &Tensor{
		id:   uuid.New(),
		name: "",
		raw:  d,
		len:  size,
		dim:  dim,
		dimc: dimc,
	}
}

// Returns a Tensor retaining the dimensions of the provided Tensor and with random values
// TODO: Normal randomization, Mean = 0, SD = 1 --> DONE
func (t *Tensor) Random() *Tensor {
	dim := make([]int, 0)
	dimc := make([]int, len(dim))
	copy(dim, t.dim)
	copy(dimc, t.dimc)
	d := make([]float64, t.len)

	rand.Seed(time.Now().UnixNano())

	for a := range d {
		d[a] = rand.NormFloat64()
	}

	return &Tensor{
		id:   uuid.New(),
		name: "",
		raw:  d,
		len:  t.len,
		dim:  dim,
		dimc: dimc,
	}
}

// Repopulate a Tensor and with random values , retaining the dimensions or using the provided dimensions
// NOTE: NOT porviding dimensions will result in using the same dimensions of the Tensor provided
// TODO: Normal randomization, Mean = 0, SD = 1 --> DONE
func (t *Tensor) Random_(dimensions ...int) {
	rand.Seed(time.Now().UnixNano())

	if dimensions == nil {

		for a := range t.raw {
			t.raw[a] = rand.NormFloat64()
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
		t.raw[a] = rand.NormFloat64()
	}

	t.raw = d
	t.dim = dim
	t.dimc = dimc
	t.len = size
}

// Return a replicate of the provided Tensor
func Clone(t *Tensor) *Tensor {
	d := make([]float64, t.len)
	dim := make([]int, t.dims)
	dimc := make([]int, t.dims)
	copy(dim, t.dim)
	copy(dimc, t.dimc)

	for a := range d {
		d[a] = t.raw[a]
	}

	return &Tensor{
		id:   uuid.New(),
		name: t.name,
		raw:  d,
		dim:  dim,
		dimc: dimc,
		dims: t.dims,
		len:  t.len,
	}
}

// Return a replicate of the provided Tensor
func (t *Tensor) Clone() *Tensor {
	d := make([]float64, t.len)
	dim := make([]int, t.dims)
	dimc := make([]int, t.dims)
	copy(dim, t.dim)
	copy(dimc, t.dimc)

	for a := range d {
		d[a] = t.raw[a]
	}

	return &Tensor{
		id:   uuid.New(),
		name: t.name,
		raw:  d,
		dim:  dim,
		dimc: dimc,
		dims: t.dims,
		len:  t.len,
	}
}
