package main

import (
	"fmt"
	"math"
)

func main(){
	sum := 0
	for i:=0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	for ; sum < 10000; {
		sum += sum
	}
	fmt.Println(sum)

	for sum < 1000000 {
		sum += sum
	}
	fmt.Println(sum)

	fmt.Println(
		powUntil(3, 2, 10),
		powUntil(3, 3, 20),
	)
}

func powUntil(x, n, lim float64) float64 {
	if v:= math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
