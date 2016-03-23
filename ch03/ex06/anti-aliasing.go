package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	if len(os.Args) <= 2 {
		fmt.Printf("usage: %s <src> <dst>\n", os.Args[0])
		os.Exit(0)
	}
	src := os.Args[1]
	dst := os.Args[2]
	img, _ := readPNG(src)
	writePNG(smooth(img), dst)
}

func readPNG(filename string) (image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return png.Decode(file)
}

func writePNG(img image.Image, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	return png.Encode(file, img)
}

func smooth(img image.Image) image.Image {
	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y
	doubleImg := resize(img, width*2, height*2)
	return resize(average(doubleImg), width, height)
}

func resize(img image.Image, newwidth, newheight int) image.Image {
	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y
	hrate := float32(width) / float32(newwidth)
	vrate := float32(height) / float32(newheight)
	newimg := image.NewRGBA(image.Rect(0, 0, newwidth, newheight))
	for newy := 0; newy < newheight; newy++ {
		for newx := 0; newx < newwidth; newx++ {
			x := (int)(float32(newx) * hrate)
			y := (int)(float32(newy) * vrate)
			x = within(x, 0, width-1)
			y = within(y, 0, height-1)
			newimg.Set(newx, newy, img.At(x, y))
		}
	}
	return newimg
}

func average(img image.Image) image.Image {
	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y
	newimg := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			var rsum, gsum, bsum uint32
			for j := -1; j <= 1; j++ {
				for i := -1; i <= 1; i++ {
					px := within(x+i, 0, width-1)
					py := within(y+j, 0, width-1)
					r, g, b, _ := img.At(px, py).RGBA()
					rsum += r >> 8
					gsum += g >> 8
					bsum += b >> 8
				}
			}
			r := rsum / 9
			g := gsum / 9
			b := bsum / 9
			c := color.RGBA{uint8(r), uint8(g), uint8(b), 255}
			newimg.Set(x, y, c)
		}
	}
	return newimg
}

func within(x, min, max int) int {
	if x < min {
		return min
	} else if x > max {
		return max
	} else {
		return x
	}
}
