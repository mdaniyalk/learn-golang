package main

import "fmt"

func main() {
	var p *int32 = new(int32)
	var i int32 

	fmt.Printf("The value of p point to is %v\n", *p)
	fmt.Printf("The value of p is %v\n", p)
	fmt.Printf("The value of i is %v\n", i)

	*p = 123
	fmt.Printf("The value of p point to is %v\n", *p)

	p = &i
	*p = 456
	fmt.Printf("The value of p point to is %v\n", *p)
	fmt.Printf("The value of i is %v\n", i)
	fmt.Printf("The value of p is %v\n", p)


	var slice = []int32{1, 2, 3, 4, 5}
	var sliceCopy = slice
	sliceCopy[0] = 3
	fmt.Println(slice)
	fmt.Println(sliceCopy)
	
}