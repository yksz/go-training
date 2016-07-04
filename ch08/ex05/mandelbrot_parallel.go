// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"runtime"
	"sync"
)

type pixel struct {
	x   int
	y   int
	val color.Color
}

const nchannels = 1024

func init() {
	ncpus := runtime.NumCPU()
	runtime.GOMAXPROCS(ncpus)
}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	pixels := make(chan *pixel, nchannels)
	var wg sync.WaitGroup
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		wg.Add(1)
		go func(py int) {
			defer wg.Done()
			for px := 0; px < width; px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				pixels <- &pixel{px, py, mandelbrot(z)}
			}
		}(py)
	}

	go func() {
		wg.Wait()
		close(pixels)
	}()

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for p := range pixels {
		// Image point (px, py) represents complex value z.
		img.Set(p.x, p.y, p.val)
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 255
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
