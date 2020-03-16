package main

import (
	"github.com/sausheong/petri"
)

var w int
var on int = petri.Deeppink
var off int = petri.White

// MySim is my simple simulation struct
type MySim struct {
	petri.Sim
}

// Init overrides the DefaultSim's Init method
func (m *MySim) Init() {
	w = *petri.Width
	m.Units = make([]petri.Cellular, w*w)
	for n := range m.Units {
		m.Units[n] = m.CreateCellWithIndex(n, off, off)
	}

	// populate the gun
	for _, xy := range gun {
		n := (xy[1] * w) + xy[0]
		m.Units[n] = m.CreateCellWithIndex(n, on, on)
	}
}

func main() {
	s := &MySim{
		petri.Sim{},
	}
	petri.Run(s)
}

// the gun pattern
var gun = [][]int{
	{1, 5}, {1, 6}, {2, 5}, {2, 6},
	{11, 5}, {11, 6}, {11, 7}, {12, 4}, {12, 8}, {13, 3},
	{13, 9}, {14, 3}, {14, 9}, {15, 6}, {16, 4}, {16, 8},
	{17, 5}, {17, 6}, {17, 7}, {18, 6},
	{21, 3}, {21, 4}, {21, 5}, {22, 3}, {22, 4}, {22, 5},
	{23, 2}, {23, 6}, {25, 1}, {25, 2}, {25, 6}, {25, 7},
	{35, 3}, {35, 4}, {36, 3}, {36, 4},
}
