package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64{
	// Zn+1 = Zn - (Zn^2 - x) / 2 * Zn
	if x == 0{
		return 0
	}else if x == 1{
		return 1
	}else{
		z := (x - 0.01) / 2
		delta := 1.0
		for delta > 0.01 {
			next := z - (math.Pow(z, 2) - x) / (2 * z)
			delta = math.Abs(z - next)
			z = next
		}
		return z
	}
}

func main(){
	fmt.Println(math.Sqrt(2))
	fmt.Println(Sqrt(2))

	fmt.Println(Sqrt(10000))
}
