package main

import "fmt"

func main() {
	var name [3]string

	name[0] = "soni"
	name[1] = "julianto"
	name[2] = "bogor"
	// name[3] = "bogor" <== ini tidak bisa karena melebihi kapasitas

	fmt.Println(name[0], name[1], name[2])

	var value = [3]int{
		90,
		90,
		90,
	}

	fmt.Println(value[0])
	fmt.Println(value[1])
	fmt.Println(value[2])

	fmt.Println(len(name))
	fmt.Println(len(value))

	var panjaneArray [10]int

	fmt.Println(len(panjaneArray))
}
