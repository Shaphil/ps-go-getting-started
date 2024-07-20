package main

import "fmt"

func main() {
	// declaration and initialization separately
	var i int
	i = 42
	fmt.Println(i)

	// declaration and initialization on the same line
	var f float32 = 3.1415926535897932
	fmt.Println(f)

	// implicit initialization syntax
	firstName := "Shaphil"
	fmt.Println(firstName)

	// boolean
	b := true
	fmt.Println(b)

	// complex data type
	c := complex(3, 4)
	fmt.Println(c)

	real, imag := real(c), imag(c)
	fmt.Println(real, imag)
}
