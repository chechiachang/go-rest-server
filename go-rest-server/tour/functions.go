package main

import "fmt"

func main(){
	f := func(x, y float64) float64{
		return x * y
	}

	fmt.Println(f(1,3))
}

func fibonacci() func() int {
	i := 0
	a := 0
	b := 1
	return func() int {
		i = a
		a = a + b
		b = 1
		return i
	}
}