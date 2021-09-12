package adapter

import (
	"fmt"
	"strings"
)

// this api is the axample case for the adapter

type Line struct {
	X1, Y1, X2, Y2 int
}
type VectorImage struct {
	Lines []Line
}

// this is the interface that we are given
func NewRectangle(width, height int) *VectorImage {
	width -= 1
	height -= 1
	return &VectorImage{[]Line{
		Line{0, 0, width, 0},
		Line{0, 0, 0, height},
		Line{width, 0, width, height},
		Line{0, height, width, height}}}
}

// issue is we can't work with this
// we have tis interface to deal with
// it deals with points instead of lines

type Point struct {
	X, Y int
}

type RasterImage interface {
	GetPoints() []Point
}

func DrawPoints(owner RasterImage) string {
	maxX, maxY := 0, 0
	points := owner.GetPoints()
	for _, pixel := range points {
		if pixel.X > maxX {
			maxX = pixel.X
		}
		if pixel.Y > maxY {
			maxY = pixel.Y
		}
	}
	maxX += 1
	maxY += 1

	// preallocate

	data := make([][]rune, maxY)
	for i := 0; i < maxY; i++ {
		data[i] = make([]rune, maxX)
		for j := range data[i] {
			data[i][j] = ' '
		}
	}

	for _, point := range points {
		data[point.Y][point.X] = '*'
	}

	b := strings.Builder{}
	for _, line := range data {
		b.WriteString(string(line))
		b.WriteRune('\n')
	}

	return b.String()

}

func api() {
	// only way to create a rectangle is by creating VectorImage but
	// the only way to print something is to provide a raster image
	// hence we need an adapter to convert the VectorImage to a Raster image

	rc := NewRectangle(6, 4)
	a := VectorToRaster(rc)
	fmt.Println(DrawPoints(a))

}
