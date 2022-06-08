package Data

import (
	"log"
)

// Returns the addition of 2 compatible Matrices
func (m1 *Matrix) Add(m2 *Matrix) *Matrix {

	if m1.col != m2.col || m1.row != m2.row {
		log.Fatalln("Incompatible data")

	}

	m := NewMatrix(nil, m2.row, m1.col)

	for a := range m.data {

		for b := range m.data[0] {
			m.data[a][b] = m1.data[a][b] + m2.data[a][b]
		}

	}

	return m
}

// Adds corresponding values
func (m1 *Matrix) Add_(m2 *Matrix) {

	if m1.col != m2.col || m1.row != m2.row {
		log.Fatalln("Incompatible data")
	}

	for a := range m1.data {

		for b := range m1.data[0] {
			m1.data[a][b] += m2.data[a][b]
		}

	}
}

// Returns subtraction of 2 compatible Matrices
func (m1 *Matrix) Subtract(m2 *Matrix) *Matrix {

	if m1.col != m2.col || m1.row != m2.row {
		log.Fatalln("Incompatible data")
	}

	m := NewMatrix(nil, m2.row, m1.col)

	for a := range m.data {

		for b := range m.data[0] {
			m.data[a][b] = m1.data[a][b] - m2.data[a][b]
		}

	}

	return m
}

// Subtracts corresponding values
func (m1 *Matrix) Subtract_(m2 *Matrix) {

	if m1.col != m2.col || m1.row != m2.row {
		log.Fatalln("Incompatible data")
	}

	for a := range m1.data {

		for b := range m1.data[0] {
			m1.data[a][b] -= m2.data[a][b]
		}

	}
}

// Returns the element-wise multiplication of 2 compatible Datas
func (m1 *Matrix) Multiply(m2 *Matrix) *Matrix {

	if m1.col != m2.col || m1.row != m2.row {
		log.Fatalln("Incompatible data")
	}

	m := NewMatrix(nil, m2.row, m1.col)

	for a := range m.data {

		for b := range m.data[0] {
			m.data[a][b] = m1.data[a][b] * m2.data[a][b]
		}

	}

	return m
}

// Multiplies corresponding values
func (m1 *Matrix) Multiply_(m2 *Matrix) {

	if m1.col != m2.col || m1.row != m2.row {
		log.Fatalln("Incompatible data")
	}

	for a := range m1.data {

		for b := range m1.data[0] {
			m1.data[a][b] *= m2.data[a][b]
		}

	}
}

// Returns the element-wise division of 2 compatible Datas
func (m1 *Matrix) Divide(m2 *Matrix) *Matrix {

	if m1.col != m2.col || m1.row != m2.row {
		log.Fatalln("Incompatible data")
	}

	m := NewMatrix(nil, m2.row, m1.col)

	for a := range m.data {

		for b := range m.data[0] {
			m.data[a][b] = m1.data[a][b] / m2.data[a][b]
		}

	}

	return m
}

// Divides corresponding values
func (m1 *Matrix) Divide_(m2 *Matrix) {

	if m1.col != m2.col || m1.row != m2.row {
		log.Fatalln("Incompatible data")
	}

	for a := range m1.data {

		for b := range m1.data[0] {

			if m2.data[a][b] == 0 {
				log.Fatalln("Bad data value")
			}

			m1.data[a][b] /= m2.data[a][b]
		}

	}
}

// Returns the dot product of 2 2-dimensional Matrices
func (m1 *Matrix) Dot(m2 *Matrix) *Matrix {

	if m1.col != m2.row {
		log.Fatalln("wrong matrix dimensions")
	}

	m := NewMatrix(nil, m1.row, m2.col)

	for a := range m1.data {

		for b := range m1.data[0] {

			for c := range m2.data[0] {
				m.data[a][c] += m1.data[a][b] * m2.data[b][c]
			}

		}

	}

	return m
}

// Returns the dot product of 2 compatible Matrices using the 1st tensor as the reciever
func (m1 *Matrix) Dot_(m2 *Matrix) {

	if m1.col != m2.row {
		log.Fatalln("wrong matrix dimensions")
	}

	d := make([][]float64, m1.row)

	for a := range m1.data {
		d[a] = make([]float64, m2.col)

		for b := range m1.data[0] {

			for c := range m2.data[0] {
				d[a][c] += m1.data[a][b] * m2.data[b][c]
			}

		}

	}

	m1.data = d
	m1.len = m1.row * m2.col
	m1.col = m2.col
}
