// a pure Go data processing package
package Data

import (
	"log"
	"math"

	"github.com/google/uuid"
)

// Matrix methods with "_" at the end will output the results in-place

// Sets an individual element of the Matrix.
// For user convenience, similar to python, negative location is also acceptable
// Hold your horses. I know your proud Go-Vein is popping ... it's just for noobs ... you should NOT use it.
func (m *Matrix) Set(data float64, dimension ...int) {

	if len(dimension) != 2 {
		log.Fatalln("Bad dimensions")
	}

	row, col := dimension[1], dimension[0]

	if col < 0 {
		col = m.col + col
	}

	if row < 0 {
		row = m.row + row
	}

	if col >= m.col || col < 0 || row >= m.row || row < 0 {
		log.Fatalln("Wrong location")
	}

	m.data[row][col] = data
}

// Returns an individual element of the Matrix
// For user convenience, similar to python, negative location is also acceptable
// Hold your horses. I know your proud Go-Vein is popping ... it's just for noobs ... you should NOT use it.
func (m *Matrix) Get(dimension ...int) float64 {

	if len(dimension) != 2 {
		log.Fatalln("Bad dimensions")
	}

	row, col := dimension[1], dimension[0]

	if col < 0 {
		col = m.col + col
	}

	if row < 0 {
		row = m.row + row
	}

	if col >= m.col || col < 0 || row >= m.row || row < 0 {
		log.Fatalln("Wrong location")
	}

	return m.data[row][col]
}

// Applies the structured function to all values within the Matrix
// # No matter how many variables the first one will be treated as input value and the rest as function parameters
// # I was even nice enough to accept an error output for your function .... where else have you seen that?
func (m *Matrix) Apply(f func(...float64) (float64, error), vars ...float64) *Matrix {
	d := make([][]float64, m.row)

	for a := range d {
		d[a] = make([]float64, m.col)
	}

	v := make([]float64, len(vars)+1)

	for a := range vars {
		v[a+1] = vars[a]
	}

	for a := range m.data {

		for b := range m.data[a] {
			v[0] = m.data[a][b]
			c, err := f(v...)

			if err != nil {
				log.Fatalln(err)
			}

			d[a][b] = c
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

// Returns the specified row
func (m *Matrix) Row(row int) []float64 {
	d := make([]float64, m.col)

	for a := range m.data[0] {
		d[a] = m.data[row][a]
	}

	return d
}

// Returns the specified column
func (m *Matrix) Column(col int) []float64 {
	d := make([]float64, m.row)

	for a := range m.data {
		d[a] = m.data[a][col]
	}

	return d
}

// Returns the Minimum value in the Matrix
func (m *Matrix) Min() float64 {
	min := m.data[0][0]

	for a := range m.data {

		for b := range m.data[a] {

			if min > m.data[a][b] {
				min = m.data[a][b]
			}

		}

	}

	return min
}

// Returns the Minimum value in the Matrix and its location
func (m *Matrix) MinLoc() (float64, *Loc) {
	min := m.data[0][0]
	col := 0
	row := 0

	for a := range m.data {

		for b := range m.data[0] {

			if min > m.data[a][b] {
				min = m.data[a][b]
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
	min := m.data[0][0]
	col, row := 0, 0

	for a := range m.data {

		for b := range m.data[0] {
			if min > m.data[a][b] {
				min = m.data[a][b]
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
	max := m.data[0][0]

	for a := range m.data {

		for b := range m.data[0] {

			if max < m.data[a][b] {
				max = m.data[a][b]
			}

		}

	}

	return max
}

// Returns the Maximum value in the Matrix and its location
func (m *Matrix) MaxLoc() (float64, *Loc) {
	max := m.data[0][0]
	col := 0
	row := 0

	for a := range m.data {

		for b := range m.data[0] {

			if max < m.data[a][b] {
				max = m.data[a][b]
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
	max := m.data[0][0]
	col, row := 0, 0

	for a := range m.data {

		for b := range m.data[0] {
			if max < m.data[a][b] {
				max = m.data[a][b]
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
	max := math.Exp(m.data[0][0])
	st.data[0][0] = max
	x, y := 0, 0
	sum := max

	for a := range m.data {

		for b := 1; b < m.col; b++ {
			st.data[a][b] = math.Exp(m.data[a][b])
			sum += st.data[a][b]

			if max < st.data[a][b] {
				max = st.data[a][b]
				x = b
				y = a
			}

		}

	}

	for a := range st.data {

		for b := range st.data[0] {
			st.data[a][b] /= sum
		}

	}

	return st, &Loc{
		Coord: []int{x, y},
	}
}

// Returns the Mean of values stored in Matrix
func (m *Matrix) Mean() float64 {
	sum := 0.0

	for _, a := range m.data {

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

	for _, a := range m.data {

		for _, b := range a {
			sum += b
		}

	}

	mean := sum / float64(m.len)

	for _, a := range m.data {

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
				d[a][b] = m.data[a][b]
			} else {
				d[a][b] = 0.0
			}

		}

	}
	m.data = d
	m.row = row
	m.col = col
}

// Returns the transpose of the Matrix
func (m *Matrix) Transpose() *Matrix {
	mm := NewMatrix(nil, m.col, m.row)

	for a := range m.data {

		for b := range m.data[0] {
			mm.data[b][a] = m.data[a][b]
		}

	}

	return mm
}

// Transposes the Matrix
func (m *Matrix) Transpose_() {
	d := make([][]float64, m.col)

	for a := range d {
		d[a] = make([]float64, m.row)

		for b := range d[0] {
			d[a][b] = m.data[b][a]
		}

	}

	m.data = d
	temp := m.col
	m.col = m.row
	m.row = temp
}

// Sets the data name
func (m *Matrix) SetName(name string) { m.name = name }

// Returns the data name
func (m *Matrix) Name() string { return m.name }

// Returns the data UUID as a string
func (m *Matrix) ID() string { return m.id.String() }

// Returns the main data
func (m *Matrix) Data() [][]float64 {
	d := make([][]float64, m.len)

	for a := range m.data {
		copy(d[a], m.data[a])
	}

	return d
}

// Returns the raw data as a slice of float64
func (m *Matrix) Raw() []float64 {
	d := make([]float64, m.len)

	for a := range m.data {

		for b := range m.data[0] {
			d[a*m.col+b] = m.data[a][b]
		}

	}

	return d
}

// Returns the data dimensions as a slice of int
func (m *Matrix) Size() (col, row int) { return m.col, m.row }

// Returns the data length
func (m *Matrix) Len() int { return m.len }
