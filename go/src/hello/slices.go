package main


func Pic(dx, dy int) [][]uint8 {
	r := make([][]uint8, dy)
	for i := 0; i < dy;i++ {
		r[i] = make([]uint8, dx)
	}

	return r
}

func main() {
}
