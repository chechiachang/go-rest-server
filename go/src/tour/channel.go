package main

import (
	"fmt"
	"math"
)

func sum(s []int, c chan int){
	sum := 0
	for _, v := range s{
		sum += v
	}
	c <- sum
}

func main(){
	s := []int{7, 1, 2, 5, 15, 123214215, 24, 124, 345, 463, 35, 1, 23, 1, 9, 6, 7,}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y, z := <-c, <-c, 0

	fmt.Println(x, y, z, x + y)

	c = make(chan int, 10)
	go double(100, c)
	for i := range c{
		fmt.Print(i, " ")
	}
	fmt.Println()
	fmt.Println(math.MaxInt32)
	fmt.Println(1 << 31 - 1)
}

func double(n int, c chan int){
	out := 1
	for i := 0; i < n; i++{
		out *= 2
		c <- out
	}
	close(c)
}
