package Data

import "github.com/google/uuid"

// a universal location identifier
// lower index, lower dimension
// e.g. : index 0 -> x (columns), 1 -> y (rows), 2 -> z, ...
type Loc struct {
	Coord []int
}

// 1D Data Container as Vector
type Vector struct {
	id   uuid.UUID
	name string
	data []float64
	len  int
}

// 2D Data Container as Matrix
type Matrix struct {
	id   uuid.UUID
	name string
	data [][]float64
	len  int
	col  int
	row  int
}

// Universal Data Container as Tensor
type Tensor struct {
	id   uuid.UUID
	name string
	raw  []float64
	dim  []int
	dimc []int
	dims int
	len  int
}

type Dataset struct {
	Data     []Datas
	Lable    []int
	size     int
	root     string
	train    bool
	download bool
	link     string
}

type Datas interface {
	Set(float64, ...int)
	Get(...int)
	Apply(f func(...float64) (float64, error), vars ...float64) *Datas
	Apply_(f func(...float64) (float64, error), vars ...float64)
	Sum()
	SumR() float64
	Cut() []*Datas
	Min() float64
	MinLoc() (float64, *Loc)
	ArgMin() int
	Max() float64
	MaxLoc() (float64, *Loc)
	ArgMax() int
	SoftMax()
	Mean()
	SD()
	Reshape()
	SetName(string)
	Name() string
	ID() string
	Raw() []float64
	Size() []int
	Len() int

	Add() Datas
	Add_()
	Subtract() Datas
	Subtract_()
	Multiply() Datas
	Multiply_()
	Divide() Datas
	Divide_()
	Dot() Datas
}
