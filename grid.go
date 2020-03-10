package petri

// FindNeighboursIndex finds the indices of the neighbouring cells
func FindNeighboursIndex(n int) (nb []int) {
	switch {
	// corner cases
	case topLeft(n):
		nb = append(nb, c5(n))
		nb = append(nb, c7(n))
		nb = append(nb, c8(n))
		return
	case topRight(n):
		nb = append(nb, c4(n))
		nb = append(nb, c6(n))
		nb = append(nb, c7(n))
		return
	case bottomLeft(n):
		nb = append(nb, c2(n))
		nb = append(nb, c3(n))
		nb = append(nb, c5(n))
		return
	case bottomRight(n):
		nb = append(nb, c1(n))
		nb = append(nb, c2(n))
		nb = append(nb, c4(n))
		return
		// side cases
	case top(n):
		nb = append(nb, c4(n))
		nb = append(nb, c5(n))
		nb = append(nb, c6(n))
		nb = append(nb, c7(n))
		nb = append(nb, c8(n))
		return
	case left(n):
		nb = append(nb, c2(n))
		nb = append(nb, c3(n))
		nb = append(nb, c5(n))
		nb = append(nb, c7(n))
		nb = append(nb, c8(n))
		return
	case right(n):
		nb = append(nb, c1(n))
		nb = append(nb, c2(n))
		nb = append(nb, c4(n))
		nb = append(nb, c6(n))
		nb = append(nb, c7(n))
		return
	case bottom(n):
		nb = append(nb, c1(n))
		nb = append(nb, c2(n))
		nb = append(nb, c3(n))
		nb = append(nb, c4(n))
		nb = append(nb, c5(n))
		return
		// everything else
	default:
		nb = append(nb, c1(n))
		nb = append(nb, c2(n))
		nb = append(nb, c3(n))
		nb = append(nb, c4(n))
		nb = append(nb, c5(n))
		nb = append(nb, c6(n))
		nb = append(nb, c7(n))
		nb = append(nb, c8(n))
	}
	return
}

// functions to check for corners and sides
func topLeft(n int) bool     { return n == 0 }
func topRight(n int) bool    { return n == *Width-1 }
func bottomLeft(n int) bool  { return n == *Width*(*Width-1) }
func bottomRight(n int) bool { return n == (*Width*(*Width))-1 }

func top(n int) bool    { return n < *Width }
func left(n int) bool   { return n%(*Width) == 0 }
func right(n int) bool  { return n%(*Width) == *Width-1 }
func bottom(n int) bool { return n >= *Width*(*Width-1) }

// functions to get the index of the neighbours
func c1(n int) int { return n - *Width - 1 }
func c2(n int) int { return n - *Width }
func c3(n int) int { return n - *Width + 1 }
func c4(n int) int { return n - 1 }
func c5(n int) int { return n + 1 }
func c6(n int) int { return n + (*Width) - 1 }
func c7(n int) int { return n + (*Width) }
func c8(n int) int { return n + (*Width) + 1 }
