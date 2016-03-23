// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 512, 512
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float32(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float32(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex64) color.Color {
	const iterations = 20
	const contrast = 10

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if squareAbs(v) > 4 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func squareAbs(x complex64) float32 {
	return real(x)*real(x) + imag(x)*imag(x)
}
