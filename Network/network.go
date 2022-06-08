package Network

import "github.com/alan59a/a08/Data"

type Network struct {
	input  *Data.Tensor
	output *Data.Tensor
	target *Data.Tensor
}

type Graph struct {
	variables []*Data.Tensor
	constants []*Data.Tensor
	updates   []*Data.Tensor
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
