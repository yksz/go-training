// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"image/color"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

var (
	red    = color.RGBA{0xff, 0x00, 0x00, 0xff}
	orange = color.RGBA{0xff, 0xa5, 0x00, 0xff}
	yellow = color.RGBA{0xff, 0xff, 0x00, 0xff}
	green  = color.RGBA{0x00, 0xff, 0x00, 0xff}
	cyan   = color.RGBA{0x00, 0xff, 0xff, 0xff}
	blue   = color.RGBA{0x00, 0x00, 0xff, 0xff}
	colors = []color.Color{red, orange, yellow, green, cyan, blue}
)

type Polygon struct {
	ax, ay float64
	bx, by float64
	cx, cy float64
	dx, dy float64
	z      float64
}

func main() {
	polygons := make([]Polygon, 0, cells*cells)
	var zmin, zmax float64 = math.MaxFloat64, 0.0
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)
			if (ax == -1 && ay == -1) ||
				(bx == -1 && by == -1) ||
				(cx == -1 && cy == -1) ||
				(dx == -1 && dy == -1) {
				continue
			}
			z := (az + bz + cz + dz) / 4.0
			zmin = math.Min(z, zmin)
			zmax = math.Max(z, zmax)
			p := Polygon{ax, ay, bx, by, cx, cy, dx, dy, z}
			polygons = append(polygons, p)
		}
	}
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for _, p := range polygons {
		fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
			p.ax, p.ay, p.bx, p.by, p.cx, p.cy, p.dx, p.dy, getColor(p.z, zmin, zmax))
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)
	if math.IsInf(z, 0) || math.IsNaN(z) {
		return -1, -1, 0
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func getColor(f float64, min float64, max float64) string {
	c := colors[0]
	delta := (max - min) / float64(len(colors))
	for i := len(colors) - 1; i > 0; i-- {
		if f <= max-delta*float64(i) {
			c = colors[i]
			break
		}
	}
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf("#%02x%02x%02x", uint8(r), uint8(g), uint8(b))
}
