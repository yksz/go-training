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
	real *big.Float
	imag *big.Float
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
			z := complex{big.NewFloat(x), big.NewFloat(y)}
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex) color.Color {
	const iterations = 20
	const contrast = 10

	v := complex{float(), float()}
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
		float().Add(x.real, y.real),
		float().Add(x.imag, y.imag)}
}

func mul(x, y complex) complex {
	return complex{
		float().Sub(float().Mul(x.real, y.real), float().Mul(x.imag, y.imag)),
		float().Add(float().Mul(x.real, y.imag), float().Mul(y.real, x.imag))}
}

func squareAbs(x complex) float64 {
	f, _ := float().Add(float().Mul(x.real, x.real), float().Mul(x.imag, x.imag)).Float64()
	return f
}

func float() *big.Float {
	return big.NewFloat(0)
}
