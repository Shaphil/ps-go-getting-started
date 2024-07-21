package main

func main() {
	println("Iterating over a slice")
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for index, value := range slice {
		println(index, value)
	}

	println("\nIterating over a map")
	wellKnownPorts := map[string]int{"http": 80, "https": 443}
	for key, value := range wellKnownPorts {
		println(key, value)
	}

	println("\nIterating over the keys of a map (ignoring the values)")
	for key := range wellKnownPorts {
		println(key)
	}

	println("\nIterating over the values of a map (ignoring the keys)")
	for _, value := range wellKnownPorts {
		println(value)
	}
}
