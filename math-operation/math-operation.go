package main

import "fmt"

func main() {

	//operasi matematika biasa
	var total = 20 + 20

	var a = 29
	var b = 13
	var c = a * b

	fmt.Println(c)
	fmt.Println(total)

	//augmented operation

	a += 20
	b -= 20
	c /= 20

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//unary operation
	c++
	var aa = -c
	var hungry = true

	fmt.Println(c)
	fmt.Println(aa)
	fmt.Println(!hungry)

}
