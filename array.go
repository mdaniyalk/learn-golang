package main

import "fmt"

func main() {
	var intArray [5]int32 
	intArray[1] = 123
	fmt.Println(intArray[0])
	fmt.Println(intArray[1:5])

	fmt.Println(&intArray[0])
	fmt.Println(&intArray[1])
	fmt.Println(&intArray[2])


	var intArray2 [5]int32 = [5]int32{1, 2, 3, 4, 5}
	fmt.Println(intArray2)

	intArray3 := [...]int32{1, 2, 3, 4, 5}
	fmt.Println(intArray3)

}