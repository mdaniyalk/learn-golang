package main

import "fmt"

func main() {
	var intSlice []int32 = []int32{1, 2, 3, 4, 5}
	fmt.Println(intSlice)
	fmt.Printf("Length: %d, Capacity: %d\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 6)
	fmt.Println(intSlice)
	fmt.Printf("Length: %d, Capacity: %d\n", len(intSlice), cap(intSlice))

	var intSlice2 []int32 = make([]int32, 5)
	fmt.Println(intSlice2)

	var intSlice3 []int32 = make([]int32, 5, 10)
	fmt.Println(intSlice3)

	var intSlice4 []int32 = []int32{1, 2, 3, 4, 5}
	intSlice = append(intSlice, intSlice4...)
	fmt.Println(intSlice)

}