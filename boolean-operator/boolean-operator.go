package main

import "fmt"

func main() {
	var ujian = 80
	var absensi = 80

	// var lulusUjian = ujian >= 80     //true
	// var lulusAbsensi = absensi >= 80 //true

	// var lulus = lulusAbsensi && lulusUjian

	// fmt.Println(lulusUjian)
	// fmt.Println(lulusAbsensi)

	var lulus = ujian >= 80 && absensi >= 80
	fmt.Println(lulus)
}
