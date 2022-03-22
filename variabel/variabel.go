package main

import "fmt"

func main() {
	var name string

	name = "soni julianto"
	fmt.Println(name)

	name = "julianto soni"
	fmt.Println(name)

	var age = 20
	fmt.Println(age)

	hobby := "coding"
	fmt.Println(hobby)

	var (
		firstName = "soni"
		lastName  = "julianto"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	dec := []byte{
		115,
		97,
		121,
		97,
	}

	fmt.Println(string(dec))

}
