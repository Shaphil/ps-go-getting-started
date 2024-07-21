package main

func main() {
	var i int
	for i < 5 {
		i++
		println(i)
		if i == 3 {
			continue
		}
		println("continue...")
	}
}
