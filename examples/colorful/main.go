package main

import (
	"flag"
	"math/rand"
	"time"

	"github.com/sausheong/petri"
)

var w int
var off int = petri.White
var ruleNum *int
var start *int

// var onColors []int = []int{0xFE2410, 0xFD3D12, 0xFD4D0D, 0xFD5D08,
// 	0xFD7208, 0xFB8804, 0xFC9803, 0xFBAA09, 0xFDBA12, 0xFCCB1D,
// 	0xFDDC21, 0xFDEB2B, 0xFEFE34}

// var onColors []int
var onColors []int = []int{0xB0171F, 0xDC143C, 0xFFB6C1, 0xFFAEB9, 0xEEA2AD, 0xCD8C95, 0x8B5F65, 0xFFC0CB, 0xFFB5C5, 0xEEA9B8, 0xCD919E, 0x8B636C, 0xDB7093, 0xFF82AB, 0xEE799F, 0xCD6889, 0x8B475D, 0xFFF0F5, 0xEEE0E5, 0xCDC1C5, 0x8B8386, 0xFF3E96, 0xEE3A8C, 0xCD3278, 0x8B2252, 0xFF69B4, 0xFF6EB4, 0xEE6AA7, 0xCD6090, 0x8B3A62, 0x872657, 0xFF1493, 0xEE1289, 0xCD1076, 0x8B0A50, 0xFF34B3, 0xEE30A7, 0xCD2990, 0x8B1C62, 0xC71585, 0xD02090, 0xDA70D6, 0xFF83FA, 0xEE7AE9, 0xCD69C9, 0x8B4789, 0xD8BFD8, 0xFFE1FF, 0xEED2EE, 0xCDB5CD, 0x8B7B8B, 0xFFBBFF, 0xEEAEEE, 0xCD96CD, 0x8B668B, 0xDDA0DD, 0xEE82EE, 0xFF00FF, 0xEE00EE, 0xCD00CD, 0x8B008B, 0x800080, 0xBA55D3, 0xE066FF, 0xD15FEE, 0xB452CD, 0x7A378B, 0x9400D3, 0x9932CC, 0xBF3EFF, 0xB23AEE, 0x9A32CD, 0x68228B, 0x4B0082, 0x8A2BE2, 0x9B30FF, 0x912CEE, 0x7D26CD, 0x551A8B, 0x9370DB, 0xAB82FF, 0x9F79EE, 0x8968CD, 0x5D478B, 0x483D8B, 0x8470FF, 0x7B68EE, 0x6A5ACD, 0x836FFF, 0x7A67EE, 0x6959CD, 0x473C8B}

func init() {
	ruleNum = flag.Int("rule", 90, "Wolfram's rules number")
	start = flag.Int("start", 0, "Start at cell")
	// 	onColors = []int{0xff0000}

	// 	for j := 0; j < 16256; j++ {
	// 		onColors = append(onColors, onColors[j]+1024)
	// 	}
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

	m.Units[*start] = m.CreateCellWithIndex(*start, onColors[0], onColors[0])
}

// returns the resultant cell
func rule(num, n uint8) uint8 {
	return (num >> n) & 1
}

// solving edge cases, if cell has no left neighbour
func before(n int) int {
	if (n % w) == 0 {
		return 0
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
	rand.Seed(time.Now().UTC().UnixNano())
	var c uint8
	for n := range m.Units {
		// figure which pattern to use
		c = 0
		if m.Units[before(n)].RGB() != off {
			c = c + 4
		}
		if m.Units[n].RGB() != off {
			c = c + 2
		}
		if m.Units[after(n)].RGB() != off {
			c = c + 1
		}
		// doesn't work for the last row, so just stop at the row before
		if n < w*(w-1) {
			// set the cell below the current cell accordingly
			set := n + w
			if rule(uint8(*ruleNum), reverse(c)) == 1 {
				color := onColors[n%len(onColors)]
				m.Units[set].Set(color)
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
