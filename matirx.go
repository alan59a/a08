// a pure Go data processing package
package main

import (
	"log"
	"math"

	"github.com/google/uuid"
)

// 2D Data Container as a Matrix
type Matrix struct {
	id       uuid.UUID
	name     string
	raw      [][]float64
	len      int
	col, row int
}

// Matrix methods with "_" at the end will output the results in-place

// Sets an individual element of the Matrix.
// For user convenience, similar to python, negative location is also acceptable
// Hold your horses. I know your proud Go-Vein is popping ... it's just for noobs ... you should NOT use it.
func (m *Matrix) Set(data float64, col, row int) {

	if col < 0 {
		col = m.col + col
	}

	if row < 0 {
		row = m.row + row
	}

	if col >= m.col || col < 0 || row >= m.row || row < 0 {
		log.Fatalln("Wrong location")
	}

	m.raw[row][col] = data
}

// Returns an individual element of the Matrix
// For user convenience, similar to python, negative location is also acceptable
// Hold your horses. I know your proud Go-Vein is popping ... it's just for noobs ... you should NOT use it.
func (m *Matrix) Get(col, row int) float64 {

	if col < 0 {
		col = m.col + col
	}

	if row < 0 {
		row = m.row + row
	}

	if col >= m.col || col < 0 || row >= m.row || row < 0 {
		log.Fatalln("Wrong location")
	}

	return m.raw[row][col]
}

// Applies the structured function to all values within the Matrix
// # No matter how many variables the first one will be treated as input value and the rest as function parameters
// # I was even nice enough to accept an error output for your function .... where else have you seen that?
func (m *Matrix) Apply(f func(...float64) (float64, error), vars ...float64) *Matrix {
	dd := make([][]float64, m.row)

	for a := range dd {
		dd[a] = make([]float64, m.col)
	}

	v := make([]float64, len(vars)+1)

	for a := range vars {
		v[a+1] = vars[a]
	}

	for a := range m.raw {

		for b := range m.raw[a] {
			v[0] = m.raw[a][b]
			c, err := f(v...)

			if err != nil {
				log.Fatalln(err)
			}

			dd[a][b] = c
		}

	}

	return &Matrix{
		id:   uuid.New(),
		name: "",
		raw:  dd,
		len:  m.col * m.row,
		col:  m.col,
		row:  m.row,
	}
}

// Returns the Minimum value in the Matrix
func (m *Matrix) Min() float64 {
	min := m.raw[0][0]

	for a := range m.raw {

		for b := range m.raw[a] {

			if min > m.raw[a][b] {
				min = m.raw[a][b]
			}

		}

	}

	return min
}

// Returns the Minimum value in the Matrix and its location
func (m *Matrix) MinLoc() (float64, *Loc) {
	min := m.raw[0][0]
	col := 0
	row := 0

	for a := range m.raw {

		for b := range m.raw[0] {

			if min > m.raw[a][b] {
				min = m.raw[a][b]
				row = a
				col = b
			}

		}

	}

	return min, &Loc{
		Coord: []int{col, row},
	}
}

// Returns the index of Minimum value in the Matrix
func (m *Matrix) ArgMin() *Loc {
	min := m.raw[0][0]
	col, row := 0, 0

	for a := range m.raw {

		for b := range m.raw[0] {
			if min > m.raw[a][b] {
				min = m.raw[a][b]
				col = b
				row = a
			}
		}

	}

	return &Loc{
		Coord: []int{col, row},
	}
}

// Returns the Maximum value in the Matrix
func (m *Matrix) Max() float64 {
	max := m.raw[0][0]

	for a := range m.raw {

		for b := range m.raw[0] {

			if max < m.raw[a][b] {
				max = m.raw[a][b]
			}

		}

	}

	return max
}

// Returns the Maximum value in the Matrix and its location
func (m *Matrix) MaxLoc() (float64, *Loc) {
	max := m.raw[0][0]
	col := 0
	row := 0

	for a := range m.raw {

		for b := range m.raw[0] {

			if max < m.raw[a][b] {
				max = m.raw[a][b]
				col = b
				row = a
			}

		}

	}

	return max, &Loc{
		Coord: []int{col, row},
	}
}

// Returns the index of Maximum value in the Matrix
func (m *Matrix) ArgMax() *Loc {
	max := m.raw[0][0]
	col, row := 0, 0

	for a := range m.raw {

		for b := range m.raw[0] {
			if max < m.raw[a][b] {
				max = m.raw[a][b]
				col = b
				row = a
			}
		}

	}

	return &Loc{
		Coord: []int{col, row},
	}
}

// Returns the index of Soft Maximum value and its index in the Matrix
func (m *Matrix) SoftMax() (*Matrix, *Loc) {
	st := m.Zero()
	max := math.Exp(m.raw[0][0])
	st.raw[0][0] = max
	x, y := 0, 0
	sum := max

	for a := range m.raw {

		for b := 1; b < m.col; b++ {
			st.raw[a][b] = math.Exp(m.raw[a][b])
			sum += st.raw[a][b]

			if max < st.raw[a][b] {
				max = st.raw[a][b]
				x = b
				y = a
			}

		}

	}

	for a := range st.raw {

		for b := range st.raw[0] {
			st.raw[a][b] /= sum
		}

	}

	return st, &Loc{
		Coord: []int{x, y},
	}
}

// Returns the Mean of values stored in Matrix
func (m *Matrix) Mean() float64 {
	sum := 0.0

	for _, a := range m.raw {

		for _, b := range a {
			sum += b
		}

	}

	return sum / float64(m.len)
}

// Returns the Standard Deviation of values stored in Matrix
func (m *Matrix) SD() float64 {
	sum := 0.0
	va := 0.0

	for _, a := range m.raw {

		for _, b := range a {
			sum += b
		}

	}

	mean := sum / float64(m.len)

	for _, a := range m.raw {

		for _, b := range a {
			va += math.Pow(mean-b, 2)
		}

	}

	return math.Sqrt(va / float64(m.len))
}

// Re-structures the dimensions of the data
// Note: This is a forceful actions and excess data will be removed
func (m *Matrix) Reshape_(col, row int) {

	if col < 1 || row < 1 {
		log.Fatalln("Bad dimensions")
	}

	m.col = col
	m.row = row

	d := make([][]float64, row)

	for a := range d {
		d[a] = make([]float64, col)
	}

	for a := range d {

		for b := range d[a] {

			if a < m.row && b < m.col {
				d[a][b] = m.raw[a][b]
			} else {
				d[a][b] = 0.0
			}

		}

	}
	m.raw = d
	m.row = row
	m.col = col
}

// Sets the data name
func (m *Matrix) SetName(name string) { m.name = name }

// Returns the data name
func (m *Matrix) Name() string { return m.name }

// Returns the data UUID as a string
func (m *Matrix) ID() string { return m.id.String() }

// Returns the raw data as a slice of float64
func (m *Matrix) Raw() [][]float64 { return m.raw }

// Returns the data dimensions as a slice of int
func (m *Matrix) Size() (col, row int) { return m.col, m.row }

// Returns the data length
func (m *Matrix) Len() int { return m.len }
