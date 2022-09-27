package main

import "fmt"

var (
	m = map[int]int{
		1: 1,
		2: 2,
	}
	// Z Expose to outside world
	Z = 42
)

func main() {
	fmt.Println("Hello, playground")

	var j int
	i := 27
	fmt.Println(i, j)
	fmt.Printf("%v, %T", m, m)
}
