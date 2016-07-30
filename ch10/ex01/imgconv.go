package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

var format string

func init() {
	flag.StringVar(&format, "out", "png", "output format")
	flag.Parse()
}

func main() {
	if err := convertTo(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "imgconv: %v\n", err)
		os.Exit(1)
	}
}

func convertTo(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	switch format {
	case "gif":
		return gif.Encode(out, img, nil)
	case "jpeg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "png":
		return png.Encode(out, img)
	default:
		return fmt.Errorf("Unsupported format = %s", format)
	}
}
