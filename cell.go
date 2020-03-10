package petri

import "image/color"

// Cell is a representation of a cell within the grid
type Cell struct {
	X      int
	Y      int
	Radius int
	Status int
	Color  color.Color
}

// XY is the x and y positions of the cell
func (c *Cell) XY() (int, int) {
	return c.X, c.Y
}

// GridIndex is the index of the grid arrays
func (c *Cell) GridIndex(width int) int {
	return (c.Y * width) + c.X
}

// RGB color integer in the form 0x1A2B3C
func (c *Cell) RGB() int {
	r, g, b, _ := c.Color.RGBA()
	return int((r & 0x00FF << 16) + (g & 0x00FF << 8) + b&0x00FF)
}

// SetRGB sets the color using the color interger in the form 0x1A2B3C
func (c *Cell) SetRGB(i int) {
	c.Color = color.RGBA{getR(i), getG(i), getB(i), uint8(255)}
}

// Clr is the Go Color interface from image/color
func (c *Cell) Clr() color.Color {
	return c.Color
}

// Size of the cell
func (c *Cell) Size() int {
	return c.Radius
}

// State is the state of the cell
func (c *Cell) State() int {
	return c.Status
}

// Set sets the state of the cell
func (c *Cell) Set(s int) {
	c.Status = s
}
