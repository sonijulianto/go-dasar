package main

import "fmt"

func main() {
	name := "budi"

	switch name := "soni"; name {
	case "soni":
		fmt.Println("apa kabar soni")
	case "ucup":
		fmt.Println("apa bakar ucup")
	case "budi":
		fmt.Println("apa kabar budi")
	default:
		fmt.Println("apa kabar kalian")
	}

	length := len(name)

	switch {
	case length > 10:
		fmt.Println("nama terlalu panjang")
	case length > 5:
		fmt.Println("nama cukup panjang")
	default:
		fmt.Println("nama sudah benar")

	}
}
