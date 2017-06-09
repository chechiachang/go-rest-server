package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(x int, y int) int {
	return x + y
}

func subtract(x int, y int) int {
	return x - y
}

func subtract1(x, y int) int {
	return x - y
}

func product(x int, y int) (z int) {
	z = x + y
	return
}

func swap(x, y string) (string, string) {
	return y, x
}

func swapAdd(x, y int) (int, int) {
	return x + y, y + x
}

func three(x, y, z int) (int, int, int) {
	return z, y, x
}

var packageVar bool;

func main() {
	fmt.Printf("hello, world\n")
	fmt.Println("Pi = ", math.Pi)
	fmt.Println("Cos Pi/2 = 3^-2 / 2 = ", math.Cos(math.Pi/6))
	fmt.Println("3^-2 / 2 = ", math.Sqrt(3) / 2)
	fmt.Println("Rand Int =", rand.Intn(10))

	fmt.Println(math.Pi)
	fmt.Println(math.Cos(0))
	fmt.Println(math.Cos(math.Pi/3))

	// Cos(Pi/6) = 3^-2 / 2
	fmt.Println(math.Cos(math.Pi/6))
	fmt.Println(math.Sqrt(3))
	fmt.Println(math.Sqrt(3)/2)

	fmt.Println("42 + 13 = ", add(42, 13))
	fmt.Println("42 - 13 = ", subtract(42, 13))
	fmt.Println("42 * 13 = ", product(42, 13))

	a, b := swap("hello", "world")
	fmt.Println("hello", "world")

	fmt.Println(a, b)
	c, d, e := three(1, 3, 5)
	fmt.Println(c, d, e)
}
