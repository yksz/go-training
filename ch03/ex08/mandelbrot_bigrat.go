// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/big"
	"os"
)

type complex struct {
	real *big.Rat
	imag *big.Rat
}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 512, 512
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex{rat().SetFloat64(x), rat().SetFloat64(y)}
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex) color.Color {
	const iterations = 20
	const contrast = 10

	v := complex{rat(), rat()}
	for n := uint8(0); n < iterations; n++ {
		v = add(mul(v, v), z)
		if squareAbs(v) > 4 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func add(x, y complex) complex {
	return complex{
		rat().Add(x.real, y.real),
		rat().Add(x.imag, y.imag)}
}

func mul(x, y complex) complex {
	return complex{
		rat().Sub(rat().Mul(x.real, y.real), rat().Mul(x.imag, y.imag)),
		rat().Add(rat().Mul(x.real, y.imag), rat().Mul(y.real, x.imag))}
}

func squareAbs(x complex) float64 {
	f, _ := rat().Add(rat().Mul(x.real, x.real), rat().Mul(x.imag, x.imag)).Float64()
	return f
}

func rat() *big.Rat {
	return big.NewRat(0, 1)
}
