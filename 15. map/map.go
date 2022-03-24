package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Soni",
		"address": "Bogor",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])
	delete(person, "name")
	fmt.Println(len(person))
	fmt.Println(person)
}
