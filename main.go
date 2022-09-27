package main

import (
	"fmt"
	"strconv"
)

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

	i := 27
	fmt.Printf("%v, %T\n", m, m)

	//convert
	j := float32(i)
	fmt.Printf("%v, %T\n", j, j)

	s := strconv.Itoa(i)
	fmt.Printf("%v, %T\n", s, s)

}
