package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	slice := arr[:]

	arr[1] = 42
	slice[2] = 27

	fmt.Println("arr:", arr)
	fmt.Println("slice:", slice)

	new_slice := []int{4, 5, 6}
	fmt.Println("new_slice:", new_slice)

	new_slice = append(new_slice, 11, 22, 33, 44, 55)
	fmt.Println("new_slice (after append):", new_slice)

	// slice of a slice
	s1 := new_slice[1:]
	s2 := new_slice[:2]
	s3 := new_slice[1:2]

	fmt.Println(s1, "\n", s2, "\n", s3)
}
