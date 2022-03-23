package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var NoKtpSoni NoKTP = "23946234"
	var getMarried Married = true

	fmt.Println(NoKtpSoni)
	fmt.Println(getMarried)
}
