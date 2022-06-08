package Data

import (
	"log"
)

// Returns the addition of 2 compatible Vectors
func (v1 *Vector) Add(v2 *Vector) *Vector {

	if v1.len != v2.len {
		log.Fatalln("Bad dimensions")
	}

	v := NewVector(nil, v1.len)

	for a := range v.data {
		v.data[a] = v1.data[a] + v2.data[a]
	}

	return v
}

// Adds corresponding values
func (v1 *Vector) Add_(v2 *Vector) {

	if v1.len != v2.len {
		log.Fatalln("Bad dimensions")
	}

	for a := range v1.data {
		v1.data[a] += v2.data[a]
	}
}

// Returns subtraction of 2 compatible Vectors
func (v1 *Vector) Subtract(v2 *Vector) *Vector {

	if v1.len != v2.len {
		log.Fatalln("Bad dimensions")
	}

	v := NewVector(nil, v1.len)

	for a := range v.data {
		v.data[a] = v1.data[a] - v2.data[a]
	}

	return v
}

// Subtracts corresponding values
func (v1 *Vector) Subtract_(v2 *Vector) {

	if v1.len != v2.len {
		log.Fatalln("Bad dimensions")
	}

	for a := range v1.data {
		v1.data[a] -= v2.data[a]
	}
}

// Returns the element-wise multiplication of 2 compatible Vectors
func (v1 *Vector) Multiply(v2 *Vector) *Vector {

	if v1.len != v2.len {
		log.Fatalln("Bad dimensions")
	}

	v := NewVector(nil, v1.len)

	for a := range v.data {
		v.data[a] = v1.data[a] * v2.data[a]
	}

	return v
}

// Multiplies corresponding values
func (v1 *Vector) Multiply_(v2 *Vector) {

	if v1.len != v2.len {
		log.Fatalln("Bad dimensions")
	}

	for a := range v1.data {
		v1.data[a] *= v2.data[a]
	}
}

// Returns the element-wise division of 2 compatible Vectors
func (v1 *Vector) Divide(v2 *Vector) *Vector {

	if v1.len != v2.len {
		log.Fatalln("Bad dimensions")
	}

	v := NewVector(nil, v1.len)

	for a := range v.data {

		if v2.data[a] != 0 {
			v.data[a] = v1.data[a] / v2.data[a]
		} else {
			log.Fatalln("Bad value")
		}
	}

	return v
}

// Divides corresponding values
func (v1 *Vector) Divide_(v2 *Vector) {

	if v1.len != v2.len {
		log.Fatalln("Bad dimensions")
	}

	for a := range v1.data {
		if v2.data[a] != 0 {
			v1.data[a] /= v2.data[a]
		} else {
			log.Fatalln("Bad value")
		}
	}
}

// Returns the dot product of 2 2-dimensional Vectors ... they are considered 1xn :D
func (v1 *Vector) Dot(v2 *Vector) *Matrix {
	m := NewMatrix(nil, v1.len, v2.len)

	for a := range v1.data {
		for b := range v2.data {
			m.data[a][b] += v1.data[a] * v2.data[b]
		}
	}

	return m
}
