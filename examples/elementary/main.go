package main

import (
	"flag"

	"github.com/sausheong/petri"
)

var w int
var on int = petri.Deeppink
var off int = petri.White

var ruleNum *int
var start *int

func init() {
	ruleNum = flag.Int("rule", 90, "Wolfram's rules number")
	start = flag.Int("start", 0, "Start at cell")
}

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

	m.Units[*start] = m.CreateCellWithIndex(*start, on, on)
}

// returns the resultant cell
func rule(num, n uint8) uint8 {
	return (num >> n) & 1
}

// solving edge cases, if cell has no left neighbour
func before(n int) int {
	if (n % w) == 0 {
		return w - 1
	} else {
		return n - 1
	}
}

// solving edge cases, if cell has no right neighbour
func after(n int) int {
	if (n % w) == (w - 1) {
		return 0
	} else {
		return n + 1
	}
}

// Process the cells
func (m *MySim) Process() {
	var c uint8
	for n := range m.Units {
		// figure which pattern to use
		c = 0
		if m.Units[before(n)].RGB() == on {
			c = c + 4
		}
		if m.Units[n].RGB() == on {
			c = c + 2
		}
		if m.Units[after(n)].RGB() == on {
			c = c + 1
		}
		// doesn't work for the last row, so just stop at the row before
		if n < w*(w-1) {
			// set the cell below the current cell accordingly
			set := n + w
			if rule(uint8(*ruleNum), c) == 1 {
				m.Units[set].Set(on)
			} else {
				m.Units[set].Set(off)
			}
		}
	}

	// change the color of the cells accordingly
	for n := range m.Units {
		m.Units[n].SetRGB(m.Units[n].State())
	}
}

func main() {
	s := &MySim{
		petri.Sim{},
	}
	petri.Run(s)
}
