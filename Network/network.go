package Network

import "github.com/alan59a/a08/Tensor"

type Network struct {
	input  *Tensor.Tensor
	output *Tensor.Tensor
	target *Tensor.Tensor
}

type Graph struct {
	variables []*Tensor.Tensor
	constants []*Tensor.Tensor
	updates   []*Tensor.Tensor
	operator  []func()
}

/*
func (n *Network) Sequence(funcs ...func(t1 *Tensor) (t2 *Tensor)) *Tensor {
	net := new(Tensor)

	for a := range funcs {
		tt = funcs[a](tt)
	}

	return tt
}
*/
