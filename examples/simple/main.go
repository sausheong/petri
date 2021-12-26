package main

import (
	"github.com/sausheong/petri"
)

// MySim is my simple simulation struct
type MySim struct {
	petri.Sim
}

func main() {
	s := &MySim{}
	petri.Run(s)
}
