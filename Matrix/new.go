package Matrix

import (
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

// Returns a Matrix with specified dimensions
// any dimension < 1 is unacceptable
func New(data [][]float64, row, col int) *Matrix {

	if col < 1 || row < 1 {
		log.Fatalln("Unacceptable dimensions")
	}

	d := make([][]float64, row)

	for a := range d {
		d[a] = make([]float64, col)
	}

	copy(d, data)

	return &Matrix{
		id:   uuid.New(),
		name: "",
		data: d,
		len:  col * row,
		col:  col,
		row:  row,
	}
}

// Returns a Matrix with specified dimensions and value of '0'
func Zero(row, col int) *Matrix {
	d := make([][]float64, row)

	for a := range d {
		d[a] = make([]float64, col)
	}

	return &Matrix{
		id:   uuid.New(),
		name: "",
		data: d,
		len:  col * row,
		col:  col,
		row:  row,
	}
}

// Returns a Matrix retaining the dimensions of the provided Matrix and with value of '0'
func (m *Matrix) Zero() *Matrix {
	d := make([][]float64, m.row)

	for a := range d {
		d[a] = make([]float64, m.col)
	}

	return &Matrix{
		id:   uuid.New(),
		name: "",
		data: d,
		len:  m.col * m.row,
		col:  m.col,
		row:  m.row,
	}
}

// Repopulate a Matrix and with value of '0' , retaining the dimensions or using the provided dimensions
// NOTE: NOT porviding dimensions will result in using the same dimensions of the Matrix provided
func (m *Matrix) Zero_(dimensions ...int) {

	if dimensions == nil {

		for a := range m.data {

			for b := range m.data[0] {
				m.data[a][b] = 0
			}

		}

		return
	} else if len(dimensions) != 2 {
		log.Fatalln("Bad dimensions")
	} else {

		for a := range dimensions {

			if dimensions[a] < 1 {
				log.Fatalln("Bad dimensions")
			}

		}

	}

	d := make([][]float64, dimensions[1])

	for a := range d {
		d[a] = make([]float64, dimensions[0])
	}

	m.data = d
	m.col = dimensions[0]
	m.row = dimensions[1]
	m.len = dimensions[0] * dimensions[1]
}

// Returns a Matrix with specified dimesnions and values of '1'
func Ones(row, col int) *Matrix {

	if col < 1 || row < 1 {
		log.Fatalln("Bad dimensions")
	}

	d := make([][]float64, row)

	for a := range d {
		d[a] = make([]float64, col)

		for b := range d[0] {
			d[a][b] = 1
		}

	}

	return &Matrix{
		id:   uuid.New(),
		name: "",
		data: d,
		len:  col * row,
		col:  col,
		row:  row,
	}
}

// Returns a Matrix retaining the dimensions of the provided Matrix and with value of '1'
func (m *Matrix) Ones() *Matrix {
	d := make([][]float64, m.row)

	for a := range d {
		d[a] = make([]float64, m.col)

		for b := range d[0] {
			d[a][b] = 1
		}

	}

	return &Matrix{
		id:   uuid.New(),
		name: "",
		data: d,
		len:  m.col * m.row,
		col:  m.col,
		row:  m.row,
	}
}

// Repopulate a Matrix and with value of '1' , retaining the dimensions or using the provided dimensions
// NOTE: NOT porviding dimensions will result in using the same dimensions of the Matrix provided
func (m *Matrix) Ones_(dimensions ...int) {

	if dimensions == nil {

		for a := range m.data {

			for b := range m.data[0] {
				m.data[a][b] = 1
			}

		}

		return
	}

	for a := range dimensions {

		if dimensions[a] < 1 {
			log.Fatalln("Bad dimensions")
		}

	}

	d := make([][]float64, dimensions[1])

	for a := range d {
		d[a] = make([]float64, dimensions[0])

		for b := range d[0] {
			d[a][b] = 1
		}

	}

	m.data = d
	m.col = dimensions[0]
	m.row = dimensions[1]
	m.len = dimensions[0] * dimensions[1]
}

// Returns a Matrix with specified dimensions and random values
// TODO: Normal randomization, Mean = 0, SD = 1 --> DONE
func Random(row, col int) *Matrix {

	if col < 1 || row < 1 {
		log.Fatalln("Bad dimensions")
	}

	d := make([][]float64, row)
	rand.Seed(time.Now().UnixNano())

	for a := range d {
		d[a] = make([]float64, col)

		for b := range d[a] {
			d[a][b] = rand.NormFloat64()
		}

	}

	return &Matrix{
		id:   uuid.New(),
		name: "",
		data: d,
		len:  col * row,
		col:  col,
		row:  row,
	}
}

// Returns a Matrix retaining the dimensions of the provided Matrix and with random values
// TODO: Normal randomization, Mean = 0, SD = 1 --> DONE
func (m *Matrix) Random() *Matrix {
	d := make([][]float64, m.row)
	rand.Seed(time.Now().UnixNano())

	for a := range d {
		d[a] = make([]float64, m.col)

		for b := range d[0] {
			d[a][b] = rand.NormFloat64()
		}

	}

	return &Matrix{
		id:   uuid.New(),
		name: "",
		data: d,
		len:  m.col * m.row,
		col:  m.col,
		row:  m.row,
	}
}

// Repopulate a Matrix and with random values , retaining the dimensions or using the provided dimensions
// NOTE: NOT porviding dimensions will result in using the same dimensions of the Matrix provided
// TODO: Normal randomization, Mean = 0, SD = 1 --> DONE
func (m *Matrix) Random_(dimensions ...int) {
	rand.Seed(time.Now().UnixNano())

	if dimensions == nil {

		for a := range m.data {

			for b := range m.data[0] {
				m.data[a][b] = rand.NormFloat64()
			}

		}

		return
	} else if len(dimensions) != 2 {
		log.Fatalln("Bad dimensions")
	}

	if dimensions[0] < 1 || dimensions[1] < 1 {
		log.Fatalln("Bad dimensions")
	}

	d := make([][]float64, dimensions[1])

	for a := range d {
		d[a] = make([]float64, dimensions[0])

		for b := range d[0] {
			d[a][b] = rand.NormFloat64()
		}

	}

	m.col = dimensions[0]
	m.row = dimensions[1]
	m.data = d
}
