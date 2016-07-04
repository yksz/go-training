#!/bin/bash -x

go build mandelbrot.go
go build mandelbrot_parallel.go
time ./mandelbrot > mandelbrot.png
time ./mandelbrot_parallel > mandelbrot_parallel.png
