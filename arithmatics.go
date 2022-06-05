package main

import (
	"log"
)

// Returns the addition of 2 compatible data Datas
func Add(t1, t2 *Tensor) *Tensor {

	for a := range t1.dim {

		if t1.dim[a] != t2.dim[a] {
			log.Fatalln("Incompatible data")
		}

	}

	t := New(nil, t1.dim...)

	for a := range t.raw {
		t.raw[a] = t1.raw[a] + t2.raw[a]
	}

	return t
}

// Returns the addition of 2 compatible data Datas
func (t1 *Tensor) Add(t2 *Tensor) *Tensor {

	for a := range t1.dim {

		if t1.dim[a] != t2.dim[a] {
			log.Fatalln("Incompatible data")
		}

	}

	t := New(nil, t1.dim...)

	for a := range t.raw {
		t.raw[a] = t1.raw[a] + t2.raw[a]
	}

	return t
}

// Adds corresponding values
func (t1 *Tensor) Add_(t2 *Tensor) {

	for a := range t1.dim {

		if t1.dim[a] != t2.dim[a] {
			log.Fatalln("Incompatible data")
		}

	}

	for a := range t1.raw {
		t1.raw[a] += t2.raw[a]
	}

}

// Returns subtraction of 2 compatible data Datas
func Sub(t1, t2 *Tensor) *Tensor {

	for a := range t1.dim {

		if t1.dim[a] != t2.dim[a] {
			log.Fatalln("Incompatible data")
		}

	}

	t := New(nil, t1.dim...)

	for a := range t.raw {
		t.raw[a] = t1.raw[a] - t2.raw[a]
	}

	return t
}

// Returns subtraction of 2 compatible data Datas
func (t1 *Tensor) Sub(t2 *Tensor) *Tensor {

	for a := range t1.dim {

		if t1.dim[a] != t2.dim[a] {
			log.Fatalln("Incompatible data")
		}

	}

	t := New(nil, t1.dim...)

	for a := range t.raw {
		t.raw[a] = t1.raw[a] - t2.raw[a]
	}

	return t
}

// Subtracts corresponding values
func (t1 *Tensor) Sub_(t2 *Tensor) {

	for a := range t1.dim {

		if t1.dim[a] != t2.dim[a] {
			log.Fatalln("Incompatible data")
		}

	}

	for a := range t1.raw {
		t1.raw[a] -= t2.raw[a]
	}

}

// Returns the element-wise multiplication of 2 compatible Datas
func Mul(t1, t2 *Tensor) *Tensor {

	for a := range t1.dim {

		if t1.dim[a] != t2.dim[a] {
			log.Fatalln("Incompatible data")
		}

	}

	t := New(nil, t1.dim...)

	for a := range t.raw {
		t.raw[a] = t1.raw[a] * t2.raw[a]
	}

	return t
}

// Returns the element-wise multiplication of 2 compatible Datas
func (t1 *Tensor) Mul(t2 *Tensor) *Tensor {

	for a := range t1.dim {

		if t1.dim[a] != t2.dim[a] {
			log.Fatalln("Incompatible data")
		}

	}

	t := New(nil, t1.dim...)

	for a := range t.raw {
		t.raw[a] = t1.raw[a] * t2.raw[a]
	}

	return t
}

// Multiplies corresponding values
func (t1 *Tensor) Mul_(t2 *Tensor) {

	for a := range t1.dim {

		if t1.dim[a] != t2.dim[a] {
			log.Fatalln("Incompatible data")
		}

	}

	for a := range t1.raw {
		t1.raw[a] /= t2.raw[a]
	}

}

// Returns the element-wise division of 2 compatible Datas
func Div(t1, t2 *Tensor) *Tensor {

	for a := range t1.dim {

		if t1.dim[a] != t2.dim[a] {
			log.Fatalln("Incompatible data")
		}

	}

	t := New(nil, t1.dim...)

	for a := range t.raw {
		t.raw[a] = t1.raw[a] / t2.raw[a]
	}

	return t
}

// Returns the element-wise division of 2 compatible Datas
func (t1 *Tensor) Div(t2 *Tensor) *Tensor {

	for a := range t1.dim {

		if t1.dim[a] != t2.dim[a] {
			log.Fatalln("Incompatible data")
		}

	}

	t := New(nil, t1.dim...)

	for a := range t.raw {
		t.raw[a] = t1.raw[a] / t2.raw[a]
	}

	return t
}

// Divides corresponding values
func (t1 *Tensor) Div_(t2 *Tensor) {

	for a := range t1.dim {

		if t1.dim[a] != t2.dim[a] {
			log.Fatalln("Incompatible data")
		}

	}

	for a := range t1.raw {
		t1.raw[a] += t2.raw[a]
	}

}

// Returns the dot product of 2 compatible Tensors
// NOTE: only > 2 dimensional Tensors are accepted
func Dot(t1, t2 *Tensor) *Tensor {

	if t1.dims != t2.dims || t1.dims < 2 {
		log.Fatalln("wrong tensor dimensions")
	}

	for a := range t1.dim {

		if a > 1 && t1.dim[a] != t2.dim[a] {
			log.Fatalln("incompatible tensors")
		}

	}

	dim := make([]int, t1.dims)
	copy(dim, t1.dim)
	dim[0] = t2.dim[0]

	t := New(nil, dim...)

	for x := 0; x < t.dim[t.dims-1]/(t.dim[0]*t.dim[1]); x++ {

		for a := 0; a < t1.dim[1]; a++ {

			for b := 0; b < t1.dim[0]; b++ {

				for c := 0; c < t2.dim[0]; c++ {
					t.raw[x*t.dimc[2]+a*t.dimc[1]+c] += t1.raw[x*t1.dimc[2]+a*t1.dimc[1]+b] * t2.raw[x*t2.dimc[2]+b*t2.dimc[1]+c]
				}

			}

		}

	}

	return t

}

// Returns the dot product of 2 2-dimensional Tensors
func (t1 *Tensor) Dot(t2 *Tensor) *Tensor {

	if t1.dims != t2.dims || t1.dims < 2 {
		log.Fatalln("wrong tensor dimensions")
	}

	for a := range t1.dim {

		if a > 1 && t1.dim[a] != t2.dim[a] {
			log.Fatalln("incompatible tensors")
		}

	}

	dim := make([]int, t1.dims)
	copy(dim, t1.dim)
	dim[0] = t2.dim[0]

	t := New(nil, dim...)

	for x := 0; x < t.dim[t.dims-1]/(t.dim[0]*t.dim[1]); x++ {

		for a := 0; a < t1.dim[1]; a++ {

			for b := 0; b < t1.dim[0]; b++ {

				for c := 0; c < t2.dim[0]; c++ {
					t.raw[x*t.dimc[2]+a*t.dimc[1]+c] += t1.raw[x*t1.dimc[2]+a*t1.dimc[1]+b] * t2.raw[x*t2.dimc[2]+b*t2.dimc[1]+c]
				}

			}

		}

	}

	return t

}

// Returns the dot product of 2 compatible Tensors using the 1st tensor as the reciever
func (t1 *Tensor) Dot_(t2 *Tensor) {

	if t1.dims != t2.dims || t1.dims < 2 {
		log.Fatalln("wrong tensor dimensions")
	}

	for a := range t1.dim {

		if a > 1 && t1.dim[a] != t2.dim[a] {
			log.Fatalln("incompatible tensors")
		}

	}

	dim := make([]int, t1.dims)
	copy(dim, t1.dim)
	dim[0] = t2.dim[0]

	t := New(nil, dim...)

	for x := 0; x < t.dim[t.dims-1]/(t.dim[0]*t.dim[1]); x++ {

		for a := 0; a < t1.dim[1]; a++ {

			for b := 0; b < t1.dim[0]; b++ {

				for c := 0; c < t2.dim[0]; c++ {
					t.raw[x*t.dimc[2]+a*t.dimc[1]+c] += t1.raw[x*t1.dimc[2]+a*t1.dimc[1]+b] * t2.raw[x*t2.dimc[2]+b*t2.dimc[1]+c]
				}

			}

		}

	}

	t1.dim = t.dim
	t1.dimc = t.dimc
	t1.raw = t.raw
	t1.len = t.len
}

// Returns the dot product of 2 compatible Tensors using the 1st tensor as the reciever
// NOTE: Any 1-dimensional data is considered a n*1 or 1*n (based on whitch product needed) dimensional Tensor
// NOTE: The results may be difficult to predict
func DotForce(t1, t2 *Tensor) *Tensor {
	switch t1.dims {
	case 1:
		switch t2.dims {

		// Tensor 1: n, Tensor 3: m 	=> Tensor: n*m
		case 1:
			t := New(nil, t1.dim[0], t2.dim[0])

			for a := range t1.raw {

				for b := range t2.raw {
					t.raw[a*t.dimc[1]+b] = t1.raw[a] * t2.raw[b]
				}

			}

			return t

		// Tensor 1: n, Tensor 3: m*n*... 	=> Tensor: m*...
		default:

			if t2.dim[1] != t1.dim[0] {
				log.Fatalln("Incompatible data")
			}

			dim := make([]int, t1.dims-1)

			for a := range dim {
				dim[a] = t2.dim[a+1]
			}

			t := New(nil, dim...)

			for a := 0; a < t2.dim[1]; a++ {

				for b := range t1.raw {
					t.raw[a*t.dimc[1]+b] += t1.raw[b] * t2.raw[b*t2.dimc[1]+a]
				}

			}

			return t

		}

	default:
		switch t2.dims {

		// Tensor 1:n*m*..., Tensor 3: n 	=> Tensor: m*...
		case 1:

			if t1.dim[t1.dims-1] != t2.dim[0] {
				log.Fatalln("Incompatible data")
			}

			dim := make([]int, t1.dims-1)
			copy(dim, t1.dim)

			t := New(nil, dim...)

			for a := range t.raw {

				for b := range t2.raw {
					t.raw[a] += t1.raw[a*t1.dim[t1.dims-1]+b]
				}

			}

			return t

		default:

			if t1.dims != t2.dims {
				log.Fatalln("Incompatible data")
			}

			for a := 0; a < t1.dims-3; a++ {
				if t1.dim[a] != t2.dim[a] {
					log.Fatalln("Incompatible data")
				}
			}

			if t1.dim[t1.dims-1] != t2.dim[t2.dims-2] {
				log.Fatalln("Incompatible data")
			}

			dim := make([]int, t1.dims)
			copy(dim, t1.dim)
			dim[len(dim)-1] = t2.dim[t2.dims-1]

			t := New(nil, dim...)

			l := 1

			for a := 0; a < t1.dims-3; a++ {
				l *= t1.dim[a]
			}

			l1 := t1.dim[t1.dims-1] * t1.dim[t1.dims-2]
			l2 := t2.dim[t2.dims-1] * t2.dim[t2.dims-2]
			l3 := dim[len(dim)-1] * dim[len(dim)-2]

			for a := 0; a < l; a++ {

				for b := 0; b < t1.dim[t1.dims-2]; b++ {

					for c := 0; c < t1.dim[t1.dims-1]; c++ {

						for e := 0; e < t2.dim[t2.dims-1]; e++ {
							t.raw[a*l3+b*t1.dim[len(dim)-1]+e] += t1.raw[a*l1+b*t1.dim[t1.dims-1]+c] * t2.raw[a*l2+c*t2.dim[t2.dims-1]+e]
						}

					}

				}

			}

			return t

		}
	}
}
