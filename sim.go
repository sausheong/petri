package petri

import (
	"flag"
	"image/color"
	"math/rand"
	"time"
)

// Default implementation is Conway's Game of Life

var population *float64

func init() {
	population = flag.Float64("pop", 0.1, "percentage initial population")
}

// Sim is the default Simulator struct
type Sim struct {
	Units []Cellular
}

// Process runs every simulation day to process the cells
func (s *Sim) Process() {
	// test each cell against the rules to determine
	// if the cell dies, survives or is born anew
	for n := range s.Units {
		neighbours := FindNeighboursIndex(n)
		ncount := 0
		for _, neighbour := range neighbours {
			if s.Units[neighbour].RGB() == Deeppink {
				ncount++
			}
		}
		if s.Units[n].RGB() == Deeppink {
			if ncount < 2 || ncount > 3 {
				s.Units[n].Set(White)
			}
		} else {
			if ncount == 3 {
				s.Units[n].Set(Deeppink)
			}
		}
	}
	// change the color of the cells accordingly
	for n := range s.Units {
		s.Units[n].SetRGB(s.Units[n].State())
	}
}

// Cells returns the cells
func (s *Sim) Cells() []Cellular {
	return s.Units
}

// Exit executes when the simulation exits
// It catches the ctrl-c signal and runs this before exiting
func (s *Sim) Exit() {
	// execute this on exit
}

// Init creates the initial cell population
func (s *Sim) Init() {
	rand.Seed(time.Now().UTC().UnixNano())
	s.Units = make([]Cellular, *Width*(*Width))
	n := 0
	for i := 1; i <= *Width; i++ {
		for j := 1; j <= *Width; j++ {
			p := rand.Float64()
			if p < *population {
				s.Units[n] = s.CreateCell(i, j, Deeppink, Deeppink)
			} else {
				s.Units[n] = s.CreateCell(i, j, White, White)
			}
			n++
		}
	}
}

// CreateCell creates a cell given the x and y locations
func (s *Sim) CreateCell(x, y, clr, st int) Cellular {
	c := Cell{
		X:      x * *CellSize,
		Y:      y * *CellSize,
		Radius: *CellSize, // radius of cell
		Status: st,
		Color:  color.RGBA{getR(clr), getG(clr), getB(clr), uint8(255)},
	}
	return &c
}

// CreateCellWithIndex creates a cell from the index of an array
func (s *Sim) CreateCellWithIndex(n, clr, st int) Cellular {
	c := Cell{
		X:      (n % (*Width)) * *CellSize,
		Y:      (int(n / (*Width))) * *CellSize,
		Radius: *CellSize, // radius of cell
		Status: st,
		Color:  color.RGBA{getR(clr), getG(clr), getB(clr), uint8(255)},
	}
	return &c
}
