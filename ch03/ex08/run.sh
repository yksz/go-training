#!/bin/sh

go run mandelbrot_bigfloat.go   > bigfloat.png
#go run mandelbrot_bigrat.go     > bigrat.png
go run mandelbrot_complex128.go > complex128.png
go run mandelbrot_complex64.go  > complex64.png
