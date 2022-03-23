package main

import "fmt"

func main() {
	var name1 = "Soni"
	var name2 = "Soni"

	var checkName = name1 == name2
	fmt.Println(checkName)

	var value1 = 5
	var value2 = 3

	fmt.Println(value1 > value2)  //true
	fmt.Println(value1 < value2)  //false
	fmt.Println(value1 == value2) //false
	fmt.Println(value1 != value2) //true
}
