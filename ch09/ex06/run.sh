#!/bin/bash -x

go build mandelbrot.go
GOMAXPROCS=1 time ./mandelbrot > mandelbrot.png
GOMAXPROCS=2 time ./mandelbrot > mandelbrot.png
GOMAXPROCS=4 time ./mandelbrot > mandelbrot.png
GOMAXPROCS=8 time ./mandelbrot > mandelbrot.png
