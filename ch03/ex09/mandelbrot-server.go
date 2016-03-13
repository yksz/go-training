package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

var (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		for k, v := range r.Form {
			switch k {
			case "x":
				x, _ := strconv.Atoi(v[0])
				xmin, xmax = x-2, x+2
			case "y":
				y, _ := strconv.Atoi(v[0])
				ymin, ymax = y-2, y+2
			case "zoom":
				zoom, _ := strconv.ParseFloat(v[0], 64)
				width, height = int(float64(width)*zoom), int(float64(height)*zoom)
			}
		}
		writeFractal(w)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func writeFractal(out io.Writer) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*float64(ymax-ymin) + float64(ymin)
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*float64(xmax-xmin) + float64(xmin)
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(out, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
