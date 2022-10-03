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

	var n bool = true
	fmt.Printf("%v, %T\n", n, n)

	var x uint16 = 42
	x = x << 3
	fmt.Printf("%v, %T\n", x, x)

	var s1 = "this is a string"
	b := []byte(s)
	fmt.Printf("%v, %T\n", s1[2], s1[2])
	fmt.Printf("%v, %T\n", b, b)

	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r)

}
