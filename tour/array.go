package main

import "fmt"

func main(){
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a)

	a1 := a[1:2]
	fmt.Println(a1)
	fmt.Println(len(a1))

	b := make([]int, 5, 5)
	fmt.Println(b)

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i := range pow {
		fmt.Println(i)
	}
}
