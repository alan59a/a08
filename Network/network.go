package main

type Network struct {
	input  *Tensor
	output *Tensor
	target *Tensor
}

type Graph struct {
	variables []*Tensor
	constants []*Tensor
	updates   []*Tensor
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
