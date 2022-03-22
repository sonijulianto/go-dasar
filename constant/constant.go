package main

import "fmt"

func main() {
	const baseUrl string = "kursusteknologi.com" //tipe data tidak wajib di delar
	const apiKey = "fbn872fb2t28bf2289"

	// baseUrl = "https://kursusteknologi.com" tidak dapat merubah value const

	fmt.Println(baseUrl)
	fmt.Println(apiKey)

	const (
		firstName string = "Soni" //tipe data string tidak wajib
		lastName         = "julianto"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
