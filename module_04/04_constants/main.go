package main

import "fmt"

func main() {
	const pi = 3.1415926535897932
	fmt.Println(pi)

	const c = 3
	fmt.Println("c =", c)

	// constant implicitly converted to an integer
	fmt.Println("c + 3 =", c+3)
	// constant implicitly converted to an float
	fmt.Println("c + 0.81 =", c+0.81)
}
