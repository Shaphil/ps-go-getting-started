package main

import "fmt"

const (
	first = iota
	second
	// pi = 3.1415
)

const (
	// `iota` resets at every new `const` block
	third = iota
	fourth
)

func main() {
	fmt.Println(first, second, third, fourth)
}
