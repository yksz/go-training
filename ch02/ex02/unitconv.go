package main

import (
	"../ex01/tempconv"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Meter float64
type Feet float64
type Pound float64
type Kilogram float64

func (m Meter) String() string    { return fmt.Sprintf("%gm", m) }
func (f Feet) String() string     { return fmt.Sprintf("%gft", f) }
func (p Pound) String() string    { return fmt.Sprintf("%glb", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%gkg", k) }

func MToF(m Meter) Feet     { return Feet(m * 1200 / 3937) }
func FToM(f Feet) Meter     { return Meter(f * 3937 / 1200) }
func PToK(p Pound) Kilogram { return Kilogram(p * 0.45359237) }
func KToP(k Kilogram) Pound { return Pound(k / 0.45359237) }

func main() {
	var args []string
	if len(os.Args) > 1 {
		args = os.Args[1:]
	} else {
		r := bufio.NewReader(os.Stdin)
		s, _ := r.ReadString('\n')
		args = []string{strings.TrimSpace(s)}
	}

	for _, arg := range args {
		v, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unitconv: %v\n", err)
			os.Exit(1)
		}
		{
			f := tempconv.Fahrenheit(v)
			c := tempconv.Celsius(v)
			fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
		}
		{
			m := Meter(v)
			f := Feet(v)
			fmt.Printf("%s = %s, %s = %s\n", m, MToF(m), f, FToM(f))
		}
		{
			p := Pound(v)
			k := Kilogram(v)
			fmt.Printf("%s = %s, %s = %s\n", p, PToK(p), k, KToP(k))
		}
	}
}
