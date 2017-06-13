package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
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

func printOS(){
	switch os:= runtime.GOOS; os {
	case "darwan":
	case "linux":
	default:
		fmt.Print("%s", os)
	}
}

func printDay(){
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
	case today + 1:
	case today + 2:
	default:

	}
}

func printTime(){
	t := time.Now()
	switch{
	case t.Hour() < 12:
	case t.Hour() < 17:
	default:

	}
}

func deferFunc(){
	defer fmt.Println("world")
	fmt.Println("hello")

	for i := 0; i < 10; i ++{
		defer fmt.Println(i)
	}

	fmt.Println("done")
}