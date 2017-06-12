package main

import "fmt"

func main(){
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	switch v:= i.(type){
	case string:
		fmt.Println("string")
		fmt.Printf("%q %v %T", v, v, v)
	case float64:
		fmt.Println("float")
		fmt.Printf("%q %v %T", v, v, v)
	default:
		fmt.Println("unknown")
	}
}
