package main

type opts struct {
	routines int
}

var opt opts

func SetGoroutines(n int) {
	opt.routines = n
}
