package petri

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"html/template"
	"image"
	"image/color"
	"image/png"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"time"

	"github.com/llgcode/draw2d/draw2dimg"
)

// directory where the public directory is in
var dir *string

// CellSize is the radius of each cell
var CellSize *int

// Width is the number of cells on one side of the image
var Width *int

// Refresh is the number of milliseconds per refresh
var Refresh *int

// Port is the port where the simulation server starts
var Port *int

// Shape is the shape of the cell
var Shape *string

var frame string // display frame

func init() {

	d, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	dir = flag.String("dir", d, "directory where the public directory is in")
	Refresh = flag.Int("f", 100, "refresh rate to capture and display the images")
	CellSize = flag.Int("s", 10, "radius of each cell in pixels")
	Width = flag.Int("w", 36, "the number of cells on one side of the image")
	Port = flag.Int("p", 12345, "tthe port where the simulation server starts")
	Shape = flag.String("shape", "square", "shape of the cell")
}

// Run executes the simulation
func Run(sim Simulator) {
	flag.Parse()
	sim.Init()
	serve(sim)
}

// Simulator is a representation of a simulation
type Simulator interface {
	Init()                                      // initialize the simulation population
	Process()                                   // process all cells each simulation day
	Cells() []Cellular                          //  an array of all cells in the simulation
	CreateCell(int, int, int, int) Cellular     // create a cell given x, y and color
	CreateCellWithIndex(int, int, int) Cellular // create a cell given the index and color
}

// Cellular is a representation of a cell
type Cellular interface {
	XY() (int, int)    // the x and y positions of the cell
	GridIndex(int) int // get the index of the grid
	RGB() int          // the RGB color of the cell
	SetRGB(int)        // set the RGB color of the cell
	Clr() color.Color  // the go image/color object
	Size() int         // size of the cell
	State() int        // get the state the cell
	Set(int)           // set the state of the cell
}

// the color integer is 0x1A2B3CFF where
// 1A is the red, 2B is green and 3C is blue

// get the red (R) from the color integer i
func getR(i int) uint8 {
	return uint8((i >> 16) & 0x0000FF)
}

// get the green (G) from the color integer i
func getG(i int) uint8 {
	return uint8((i >> 8) & 0x0000FF)
}

// get the blue (B) from the color integer i
func getB(i int) uint8 {
	return uint8(i & 0x0000FF)
}

// FindEmpty find all empty cells in the simlator
func FindEmpty() (sim Simulator, empty []int) {
	for n, cell := range sim.Cells() {
		if cell.RGB() == 0 {
			empty = append(empty, n)
		}
	}
	return
}

// The web server for the simulator starts here

// set up the server, start generating frames, start the server and
// open up a browser to show the simulation
func serve(sim Simulator) {
	mux := http.NewServeMux()
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir(*dir+"/public"))))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/frame", getFrame)
	server := &http.Server{
		Addr:    "0.0.0.0:" + strconv.Itoa(*Port),
		Handler: mux,
	}
	fmt.Println("Started simulation server at", server.Addr)
	fmt.Println("ctrl-c to stop simulation.")
	// start generating frames in a new goroutine
	go generateFrames(sim)
	// open in default browser
	go open("http://0.0.0.0:" + strconv.Itoa(*Port))
	server.ListenAndServe()
}

// index for web server
func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(*dir + "/public/index.html")
	t.Execute(w, *Refresh)
}

// push the frame to the browser
func getFrame(w http.ResponseWriter, r *http.Request) {
	str := "data:image/png;base64," + frame
	w.Header().Set("Cache-Control", "no-cache")
	w.Write([]byte(str))
}

// continually generate frames at every period
func generateFrames(sim Simulator) {
	for {
		sim.Process()
		img := draw(*Width*(*CellSize)+(*CellSize), sim.Cells())
		createFrame(img) // create the frame from the sensor
		time.Sleep(time.Duration(*Refresh) * time.Millisecond)
	}
}

// create a frame from the image
func createFrame(img image.Image) {
	var buf bytes.Buffer
	png.Encode(&buf, img)
	frame = base64.StdEncoding.EncodeToString(buf.Bytes())
}

// draw the cells
func draw(w int, cells []Cellular) *image.RGBA {
	dest := image.NewRGBA(image.Rect(0, 0, w, w))
	gc := draw2dimg.NewGraphicContext(dest)
	for _, cell := range cells {
		x, y := cell.XY()

		gc.SetFillColor(cell.Clr())
		switch *Shape {
		case "square":
			drawSquare(gc, x, y)
		case "circle":
			drawCircle(gc, cell.Size(), x, y)
		default:
			drawSquare(gc, x, y)
		}
		gc.Close()
		gc.Fill()
	}
	return dest
}

// draw square cell
func drawSquare(gc *draw2dimg.GraphicContext, x, y int) {
	gc.BeginPath()
	gc.MoveTo(float64(x), float64(y))
	gc.LineTo(float64(x+(*CellSize)), float64(y))
	gc.LineTo(float64(x+(*CellSize)), float64(y+(*CellSize)))
	gc.LineTo(float64(x), float64(y+(*CellSize)))
	gc.LineTo(float64(x), float64(y))
}

// draw a circular cell
func drawCircle(gc *draw2dimg.GraphicContext, size, x, y int) {
	gc.MoveTo(float64(x), float64(y))
	gc.ArcTo(float64(x), float64(y),
		float64(size/2), float64(size/2), 0, 6.283185307179586)
}

// open in the default browser
func open(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}
