package main

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main(){
	m = map[string]Vertex{
		"bell":Vertex{1,3},
		"google":{3,4},
	}
}
