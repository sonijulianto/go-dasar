package main

import "fmt"

func main() {
	person := "Soni"
	age := 22

	if person == "Soni" && age == 22 {
		fmt.Println("hello kakak")
	} else if person == "Soni" {
		fmt.Println("hello " + person)
	} else {
		fmt.Println("hello kalian")
	}

	if angka := 40; angka == 20 {
		fmt.Println("angkanya 20")
	} else if angka > 20 {
		fmt.Println("angka lebih dari 20")
	} else {
		fmt.Println("angka kurang dari 20")
	}
}
