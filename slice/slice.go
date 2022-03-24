package main

import "fmt"

func main() {
	var month = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = month[4:7]

	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// month[5] = "Diubah"
	// fmt.Println(slice1)

	// slice1[0] = "Mei lagi"
	// fmt.Println(month)

	var slice2 = month[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Soni")
	fmt.Println(slice3)
	slice3[1] = "Bukan desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(month)

	newSlice := make([]string, 2, 5)

	newSlice[0] = "Soni"
	newSlice[1] = "Julianto"

	fmt.Println(newSlice)

	copySlice := make([]string, 5, cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
