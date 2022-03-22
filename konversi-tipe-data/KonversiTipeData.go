package main

import (
	"fmt"
)

func main() {
	var int32 int32 = 123456
	var int64 int64 = int64(int32)
	var int8 int8 = int8(int32) // hasilnya (64) karena panjang dari int8 adalah 127 kalau melebihi itu maka akan dihitung mundur dari -127

	fmt.Println(int32)
	fmt.Println(int64)
	fmt.Println(int8)

	var name string = "soni"
	var s byte = name[0]
	var eString = string(s)

	fmt.Println(name)
	fmt.Println(eString)
}
