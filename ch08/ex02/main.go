package main

import (
	"flag"

	"./ftp"
)

var port int

func init() {
	flag.IntVar(&port, "port", 21, "port number")
	flag.Parse()
}

func main() {
	ftp.ListenAndServe(port)
}
