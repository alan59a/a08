package main

import (
	"log"
)

// Outputs the addition of 2 compatible data Datas
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

// Outputs the addition of 2 compatible data Datas
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

// Outputs subtraction of 2 compatible data Datas
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

// Outputs subtraction of 2 compatible data Datas
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

// Outputs the element-wise multiplication of 2 compatible Datas
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

// Outputs the element-wise multiplication of 2 compatible Datas
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

// outputs the element-wise division of 2 compatible Datas
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

// outputs the element-wise division of 2 compatible Datas
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

// Outputs the dot product of 2 compatible Datas.
// if that is not your intension please use the Mul() method.
// as a general rule and having the user's convenience in mind; if you do NOT want any changes to the dimensions of your data please refraion from using this method.
// e.g.:using this method for 2 1-dimensional Tensor will result in a 2-dimensional Tensor.
func Dot(t1, t2 *Tensor) *Tensor {

	switch len(t1.dim) {
	case 1:
		switch len(t2.dim) {
		case 1:
			t := New(nil, t1.dim[0], t2.dim[0])

			for a := range t1.raw {

				for b := range t2.raw {
					t.raw[a*t.dim[0]+b] = t1.raw[a] * t2.raw[b]
				}

			}

			return t

		case 2:

			if t1.dim[0] != t2.dim[0] {
				log.Fatalln("Incompatible data")
			}

			t := New(nil, t2.dim[1])

			for a := 0; a < t2.dim[1]; a++ {

				for b := range t1.raw {
					t.raw[a] += t1.raw[b] * t2.raw[a*t2.dim[0]+b]
				}

			}

			return t

		default:

			if t2.dim[len(t1.dim)-2] != t1.dim[0] {
				log.Fatalln("Incompatible data")
			}

			dim := make([]int, len(t1.dim)-1)
			copy(dim, t1.dim)
			dim[len(dim)-1] = t1.dim[len(t1.dim)-1]

			t := New(nil, dim...)

			for a := range t.raw {

				for b := range t1.raw {
					t.raw[a] += t1.raw[b] * t2.raw[b*t1.dim[len(t1.dim)-1]+a]
				}

			}

			return t

		}
	case 2:
		switch len(t2.dim) {
		case 1:

			if t1.dim[1] != t2.dim[0] {
				log.Fatalln("Incompatible data")
			}

			t := New(nil, t1.dim[0])

			for a := range t2.raw {

				for b := 0; b < t1.dim[0]; b++ {
					t.raw[b] += t1.raw[a+b*t1.dim[1]] * t2.raw[a]
				}

			}

			return t

		case 2:

			if t1.dim[1] != t2.dim[0] {
				log.Fatalln("Incompatible data")
			}

			t := New(nil, t1.dim[0], t2.dim[1])

			for a := 0; a < t1.dim[0]; a++ {

				for b := 0; b < t1.dim[1]; b++ {

					for c := 0; c < t2.dim[0]; c++ {
						t.raw[a*t1.dim[1]+c] += t1.raw[a*t1.dim[1]+b] * t2.raw[b*t2.dim[0]+c]
					}

				}

			}

			return t

		default:

			if t1.dim[1] != t2.dim[len(t2.dim)-2] {
				log.Fatalln("Incompatible data")
			}

			dim := make([]int, len(t2.dim))
			copy(dim, t2.dim)
			dim[len(dim)-2] = t1.dim[0]

			t := New(nil, dim...)

			l := 1

			for a := 0; a < len(t2.dim)-3; a++ {
				l *= t2.dim[a]
			}

			l2 := t2.dim[len(t2.dim)-1] * t2.dim[len(t2.dim)-2]
			l3 := dim[len(dim)-1] * dim[len(dim)-2]

			for a := 0; a < l; a++ {

				for b := 0; b < t1.dim[len(t1.dim)-2]; b++ {

					for c := 0; c < t1.dim[len(t1.dim)-1]; c++ {

						for e := 0; e < t2.dim[0]; e++ {
							t.raw[a*l3+b*dim[len(dim)-1]+e] = t1.raw[b*t1.dim[len(t1.dim)-1]+c] * t2.raw[a*l2+c*t2.dim[1]+e]
						}

					}

				}

			}

			return t

		}
	default:
		switch len(t2.dim) {
		case 1:

			if t1.dim[len(t1.dim)-1] != t2.dim[0] {
				log.Fatalln("Incompatible data")
			}

			dim := make([]int, len(t1.dim)-1)
			copy(dim, t1.dim)

			t := New(nil, dim...)

			for a := range t.raw {

				for b := range t2.raw {
					t.raw[a] += t1.raw[a*t1.dim[len(t1.dim)-1]+b]
				}

			}

			return t

		case 2:

			if t1.dim[len(t1.dim)-1] != t2.dim[0] {
				log.Fatalln("Incompatible data")
			}

			dim := make([]int, len(t1.dim))
			copy(dim, t1.dim)
			dim[len(dim)-1] = t2.dim[1]

			t := New(nil, dim...)

			l := 1

			for a := 0; a < len(t1.dim)-3; a++ {
				l *= t1.dim[a]
			}

			l1 := t1.dim[len(t1.dim)-1] * t1.dim[len(t1.dim)-2]
			l3 := dim[len(dim)-1] * dim[len(dim)-2]

			for a := 0; a < l; a++ {

				for b := 0; b < t1.dim[len(t1.dim)-2]; b++ {

					for c := 0; c < t1.dim[len(t1.dim)-1]; c++ {

						for e := 0; e < t2.dim[0]; e++ {
							t.raw[a*l3+b*dim[len(dim)-1]+e] = t1.raw[a*l1+b*t1.dim[len(t1.dim)-1]+c] * t2.raw[c*t2.dim[1]+e]
						}

					}

				}

			}

			return t

		default:

			if len(t1.dim) != len(t2.dim) {
				log.Fatalln("Incompatible data")
			}

			for a := 0; a < len(t1.dim)-3; a++ {
				if t1.dim[a] != t2.dim[a] {
					log.Fatalln("Incompatible data")
				}
			}

			if t1.dim[len(t1.dim)-1] != t2.dim[len(t2.dim)-2] {
				log.Fatalln("Incompatible data")
			}

			dim := make([]int, len(t1.dim))
			copy(dim, t1.dim)
			dim[len(dim)-1] = t2.dim[len(t2.dim)-1]

			t := New(nil, dim...)

			l := 1

			for a := 0; a < len(t1.dim)-3; a++ {
				l *= t1.dim[a]
			}

			l1 := t1.dim[len(t1.dim)-1] * t1.dim[len(t1.dim)-2]
			l2 := t2.dim[len(t2.dim)-1] * t2.dim[len(t2.dim)-2]
			l3 := dim[len(dim)-1] * dim[len(dim)-2]

			for a := 0; a < l; a++ {

				for b := 0; b < t1.dim[len(t1.dim)-2]; b++ {

					for c := 0; c < t1.dim[len(t1.dim)-1]; c++ {

						for e := 0; e < t2.dim[len(t2.dim)-1]; e++ {
							t.raw[a*l3+b*t1.dim[len(dim)-1]+e] += t1.raw[a*l1+b*t1.dim[len(t1.dim)-1]+c] * t2.raw[a*l2+c*t2.dim[len(t2.dim)-1]+e]
						}

					}

				}

			}

			return t

		}
	}
}

// Outputs the dot product of 2 compatible Datas.
// if that is not your intension please use the Mul() method.
// as a general rule and having the user's convenience in mind; if you do NOT want any changes to the dimensions of your data please refraion from using this method.
// e.g.:using this method for 2 1-dimensional Tensor will result in a 2-dimensional Tensor.
func (t1 *Tensor) Dot(t2 *Tensor) *Tensor {

	switch len(t1.dim) {
	case 1:
		switch len(t2.dim) {
		case 1:
			t := New(nil, t1.dim[0], t2.dim[0])

			for a := range t1.raw {

				for b := range t2.raw {
					t.raw[a*t.dim[0]+b] = t1.raw[a] * t2.raw[b]
				}

			}

			return t

		case 2:

			if t1.dim[0] != t2.dim[0] {
				log.Fatalln("Incompatible data")
			}

			t := New(nil, t2.dim[1])

			for a := 0; a < t2.dim[1]; a++ {

				for b := range t1.raw {
					t.raw[a] += t1.raw[b] * t2.raw[a*t2.dim[0]+b]
				}

			}

			return t

		default:

			if t2.dim[len(t1.dim)-2] != t1.dim[0] {
				log.Fatalln("Incompatible data")
			}

			dim := make([]int, len(t1.dim)-1)
			copy(dim, t1.dim)
			dim[len(dim)-1] = t1.dim[len(t1.dim)-1]

			t := New(nil, dim...)

			for a := range t.raw {

				for b := range t1.raw {
					t.raw[a] += t1.raw[b] * t2.raw[b*t1.dim[len(t1.dim)-1]+a]
				}

			}

			return t

		}
	case 2:
		switch len(t2.dim) {
		case 1:

			if t1.dim[1] != t2.dim[0] {
				log.Fatalln("Incompatible data")
			}

			t := New(nil, t1.dim[0])

			for a := range t2.raw {

				for b := 0; b < t1.dim[0]; b++ {
					t.raw[b] += t1.raw[a+b*t1.dim[1]] * t2.raw[a]
				}

			}

			return t

		case 2:

			if t1.dim[1] != t2.dim[0] {
				log.Fatalln("Incompatible data")
			}

			t := New(nil, t1.dim[0], t2.dim[1])

			for a := 0; a < t1.dim[0]; a++ {

				for b := 0; b < t1.dim[1]; b++ {

					for c := 0; c < t2.dim[0]; c++ {
						t.raw[a*t1.dim[1]+c] += t1.raw[a*t1.dim[1]+b] * t2.raw[b*t2.dim[0]+c]
					}

				}

			}

			return t

		default:

			if t1.dim[1] != t2.dim[len(t2.dim)-2] {
				log.Fatalln("Incompatible data")
			}

			dim := make([]int, len(t2.dim))
			copy(dim, t2.dim)
			dim[len(dim)-2] = t1.dim[0]

			t := New(nil, dim...)

			l := 1

			for a := 0; a < len(t2.dim)-3; a++ {
				l *= t2.dim[a]
			}

			l2 := t2.dim[len(t2.dim)-1] * t2.dim[len(t2.dim)-2]
			l3 := dim[len(dim)-1] * dim[len(dim)-2]

			for a := 0; a < l; a++ {

				for b := 0; b < t1.dim[len(t1.dim)-2]; b++ {

					for c := 0; c < t1.dim[len(t1.dim)-1]; c++ {

						for e := 0; e < t2.dim[0]; e++ {
							t.raw[a*l3+b*dim[len(dim)-1]+e] = t1.raw[b*t1.dim[len(t1.dim)-1]+c] * t2.raw[a*l2+c*t2.dim[1]+e]
						}

					}

				}

			}

			return t

		}
	default:
		switch len(t2.dim) {
		case 1:

			if t1.dim[len(t1.dim)-1] != t2.dim[0] {
				log.Fatalln("Incompatible data")
			}

			dim := make([]int, len(t1.dim)-1)
			copy(dim, t1.dim)

			t := New(nil, dim...)

			for a := range t.raw {

				for b := range t2.raw {
					t.raw[a] += t1.raw[a*t1.dim[len(t1.dim)-1]+b]
				}

			}

			return t

		case 2:

			if t1.dim[len(t1.dim)-1] != t2.dim[0] {
				log.Fatalln("Incompatible data")
			}

			dim := make([]int, len(t1.dim))
			copy(dim, t1.dim)
			dim[len(dim)-1] = t2.dim[1]

			t := New(nil, dim...)

			l := 1

			for a := 0; a < len(t1.dim)-3; a++ {
				l *= t1.dim[a]
			}

			l1 := t1.dim[len(t1.dim)-1] * t1.dim[len(t1.dim)-2]
			l3 := dim[len(dim)-1] * dim[len(dim)-2]

			for a := 0; a < l; a++ {

				for b := 0; b < t1.dim[len(t1.dim)-2]; b++ {

					for c := 0; c < t1.dim[len(t1.dim)-1]; c++ {

						for e := 0; e < t2.dim[0]; e++ {
							t.raw[a*l3+b*dim[len(dim)-1]+e] = t1.raw[a*l1+b*t1.dim[len(t1.dim)-1]+c] * t2.raw[c*t2.dim[1]+e]
						}

					}

				}

			}

			return t

		default:

			if len(t1.dim) != len(t2.dim) {
				log.Fatalln("Incompatible data")
			}

			for a := 0; a < len(t1.dim)-3; a++ {
				if t1.dim[a] != t2.dim[a] {
					log.Fatalln("Incompatible data")
				}
			}

			if t1.dim[len(t1.dim)-1] != t2.dim[len(t2.dim)-2] {
				log.Fatalln("Incompatible data")
			}

			dim := make([]int, len(t1.dim))
			copy(dim, t1.dim)
			dim[len(dim)-1] = t2.dim[len(t2.dim)-1]

			t := New(nil, dim...)

			l := 1

			for a := 0; a < len(t1.dim)-3; a++ {
				l *= t1.dim[a]
			}

			l1 := t1.dim[len(t1.dim)-1] * t1.dim[len(t1.dim)-2]
			l2 := t2.dim[len(t2.dim)-1] * t2.dim[len(t2.dim)-2]
			l3 := dim[len(dim)-1] * dim[len(dim)-2]

			for a := 0; a < l; a++ {

				for b := 0; b < t1.dim[len(t1.dim)-2]; b++ {

					for c := 0; c < t1.dim[len(t1.dim)-1]; c++ {

						for e := 0; e < t2.dim[len(t2.dim)-1]; e++ {
							t.raw[a*l3+b*t1.dim[len(dim)-1]+e] += t1.raw[a*l1+b*t1.dim[len(t1.dim)-1]+c] * t2.raw[a*l2+c*t2.dim[len(t2.dim)-1]+e]
						}

					}

				}

			}

			return t

		}
	}
}

// Outputs the inverse of the Tensor
// TO DO
func (t *Tensor) Inv() *Tensor {

	if t.dim[len(t.dim)-1] != t.dim[len(t.dim)-2] {
		log.Fatalln(" Incompatible Tensor")
	}

	dd := New(nil, t.dim...)

	return dd
}

// Outputs the dot product of 2 compatible Datas.
// if that is not your intension please use the Mul() method.
// as a general rule and having the user's convenience in mind; if you do NOT want any changes to the dimensions of your data please refraion from using this method.
// e.g.:using this method for 2 1-dimensional Tensor will result in a 2-dimensional Tensor.
func (t1 *Tensor) Dot_(t2 *Tensor) *Tensor {

	switch len(t1.dim) {
	case 1:
		switch len(t2.dim) {
		case 1:
			t := New(nil, t1.dim[0], t2.dim[0])

			for a := range t1.raw {

				for b := range t2.raw {
					t.raw[a*t.dim[0]+b] = t1.raw[a] * t2.raw[b]
				}

			}

			return t

		case 2:

			if t1.dim[0] != t2.dim[0] {
				log.Fatalln("Incompatible data")
			}

			t := New(nil, t2.dim[1])

			for a := 0; a < t2.dim[1]; a++ {

				for b := range t1.raw {
					t.raw[a] += t1.raw[b] * t2.raw[a*t2.dim[0]+b]
				}

			}

			return t

		default:

			if t2.dim[len(t1.dim)-2] != t1.dim[0] {
				log.Fatalln("Incompatible data")
			}

			dim := make([]int, len(t1.dim)-1)
			copy(dim, t1.dim)
			dim[len(dim)-1] = t1.dim[len(t1.dim)-1]

			t := New(nil, dim...)

			for a := range t.raw {

				for b := range t1.raw {
					t.raw[a] += t1.raw[b] * t2.raw[b*t1.dim[len(t1.dim)-1]+a]
				}

			}

			return t

		}
	case 2:
		switch len(t2.dim) {
		case 1:

			if t1.dim[1] != t2.dim[0] {
				log.Fatalln("Incompatible data")
			}

			t := New(nil, t1.dim[0])

			for a := range t2.raw {

				for b := 0; b < t1.dim[0]; b++ {
					t.raw[b] += t1.raw[a+b*t1.dim[1]] * t2.raw[a]
				}

			}

			return t

		case 2:

			if t1.dim[1] != t2.dim[0] {
				log.Fatalln("Incompatible data")
			}

			t := New(nil, t1.dim[0], t2.dim[1])

			for a := 0; a < t1.dim[0]; a++ {

				for b := 0; b < t1.dim[1]; b++ {

					for c := 0; c < t2.dim[0]; c++ {
						t.raw[a*t1.dim[1]+c] += t1.raw[a*t1.dim[1]+b] * t2.raw[b*t2.dim[0]+c]
					}

				}

			}

			return t

		default:

			if t1.dim[1] != t2.dim[len(t2.dim)-2] {
				log.Fatalln("Incompatible data")
			}

			dim := make([]int, len(t2.dim))
			copy(dim, t2.dim)
			dim[len(dim)-2] = t1.dim[0]

			t := New(nil, dim...)

			l := 1

			for a := 0; a < len(t2.dim)-3; a++ {
				l *= t2.dim[a]
			}

			l2 := t2.dim[len(t2.dim)-1] * t2.dim[len(t2.dim)-2]
			l3 := dim[len(dim)-1] * dim[len(dim)-2]

			for a := 0; a < l; a++ {

				for b := 0; b < t1.dim[len(t1.dim)-2]; b++ {

					for c := 0; c < t1.dim[len(t1.dim)-1]; c++ {

						for e := 0; e < t2.dim[0]; e++ {
							t.raw[a*l3+b*dim[len(dim)-1]+e] = t1.raw[b*t1.dim[len(t1.dim)-1]+c] * t2.raw[a*l2+c*t2.dim[1]+e]
						}

					}

				}

			}

			return t

		}
	default:
		switch len(t2.dim) {
		case 1:

			if t1.dim[len(t1.dim)-1] != t2.dim[0] {
				log.Fatalln("Incompatible data")
			}

			dim := make([]int, len(t1.dim)-1)
			copy(dim, t1.dim)

			t := New(nil, dim...)

			for a := range t.raw {

				for b := range t2.raw {
					t.raw[a] += t1.raw[a*t1.dim[len(t1.dim)-1]+b]
				}

			}

			return t

		case 2:

			if t1.dim[len(t1.dim)-1] != t2.dim[0] {
				log.Fatalln("Incompatible data")
			}

			dim := make([]int, len(t1.dim))
			copy(dim, t1.dim)
			dim[len(dim)-1] = t2.dim[1]

			t := New(nil, dim...)

			l := 1

			for a := 0; a < len(t1.dim)-3; a++ {
				l *= t1.dim[a]
			}

			l1 := t1.dim[len(t1.dim)-1] * t1.dim[len(t1.dim)-2]
			l3 := dim[len(dim)-1] * dim[len(dim)-2]

			for a := 0; a < l; a++ {

				for b := 0; b < t1.dim[len(t1.dim)-2]; b++ {

					for c := 0; c < t1.dim[len(t1.dim)-1]; c++ {

						for e := 0; e < t2.dim[0]; e++ {
							t.raw[a*l3+b*dim[len(dim)-1]+e] = t1.raw[a*l1+b*t1.dim[len(t1.dim)-1]+c] * t2.raw[c*t2.dim[1]+e]
						}

					}

				}

			}

			return t

		default:

			if len(t1.dim) != len(t2.dim) {
				log.Fatalln("Incompatible data")
			}

			for a := 0; a < len(t1.dim)-3; a++ {
				if t1.dim[a] != t2.dim[a] {
					log.Fatalln("Incompatible data")
				}
			}

			if t1.dim[len(t1.dim)-1] != t2.dim[len(t2.dim)-2] {
				log.Fatalln("Incompatible data")
			}

			dim := make([]int, len(t1.dim))
			copy(dim, t1.dim)
			dim[len(dim)-1] = t2.dim[len(t2.dim)-1]

			t := New(nil, dim...)

			l := 1

			for a := 0; a < len(t1.dim)-3; a++ {
				l *= t1.dim[a]
			}

			l1 := t1.dim[len(t1.dim)-1] * t1.dim[len(t1.dim)-2]
			l2 := t2.dim[len(t2.dim)-1] * t2.dim[len(t2.dim)-2]
			l3 := dim[len(dim)-1] * dim[len(dim)-2]

			for a := 0; a < l; a++ {

				for b := 0; b < t1.dim[len(t1.dim)-2]; b++ {

					for c := 0; c < t1.dim[len(t1.dim)-1]; c++ {

						for e := 0; e < t2.dim[len(t2.dim)-1]; e++ {
							t.raw[a*l3+b*t1.dim[len(dim)-1]+e] += t1.raw[a*l1+b*t1.dim[len(t1.dim)-1]+c] * t2.raw[a*l2+c*t2.dim[len(t2.dim)-1]+e]
						}

					}

				}

			}

			return t

		}
	}
}

// Outputs the inverse of the Tensor
// TO DO
func (t *Tensor) Inv_() {

	if t.dim[len(t.dim)-1] != t.dim[len(t.dim)-2] {
		log.Fatalln(" Incompatible Tensor")
	}

}
