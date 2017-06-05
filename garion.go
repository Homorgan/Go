package main

import (
	"fmt"
	"math"
	"strconv"
)

const DELTA = 0.00000000000001
const INITIAL_Z = 100.0

func Sqrt(x float64) float64 {
	z := 1.0

	step := func() float64 {
		return z - (z*z -x) / (2 * z)
	}

	for zz := step(); math.Abs(zz - z) > DELTA
	{
		z = zz
		zz = step()
	}
	return z
}

func main() {
	
	fmt.Println("Please enter a number")
	var i string
	fmt.Scan(&i)

	f, _ := strconv.ParseFloat(i, 64)
	
	fmt.Println(Sqrt(f))
	fmt.Println(math.Sqrt(f))
}
