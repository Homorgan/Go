package main

import (
	"fmt"
	"math"
)

const DELTA = 0.00000000000000001
const INITIAL_Z = 100.0

//func Sqrt(x float64) float64 {
//
//}

func Sqrt(x float64) float64 {
    z := 1.0 
    
    step := func() float64 {
    	return z - (z*z - x) / (2 * z)
    }
    
    for zz := step(); math.Abs(zz - z) > DELTA
    {
    	z = zz
	zz = step()
    }
    return z
}

func main() {
	fmt.Println(Sqrt(20))
}
