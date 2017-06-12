package main

import "fmt"

func main(){
	i, j := 42, 2901

	p := &i
	fmt.Println(i)
	fmt.Println(*p)

	fmt.Println(&i)
	fmt.Println(p)

	*p = 21
	fmt.Println(i)
	fmt.Println(&i)

	q := &j
	fmt.Println(j)
	fmt.Println(*q)

	fmt.Println(Vertex{1, 2})
	fmt.Println("===")

	v := Vertex{1, 2}
	v.X = 5
	fmt.Println(&v)
	fmt.Println(&v.X)

	r := &v
	fmt.Println(r)
	fmt.Println("r.Y", r.Y)

	fmt.Println(*(&v.X))
	fmt.Println(v)

}

type Vertex struct {
	X int
	Y int
}


