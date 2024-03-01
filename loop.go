package main

import "fmt"

func main() {

	var myMap2 = map[string]int{"adam": 1, "bob": 2, "charlie": 3}

	for name := range myMap2{
		fmt.Printf("Name: %s, Value: %d\n", name, myMap2[name])
	}

	var intArray2 = [5]int32{1, 2, 3, 4, 5}

	for i, v, := range intArray2{
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	var i int = 0
	for i<10{
		fmt.Println(i)
		i++
	}

	var i int = 0
	for {
		if i>=10{
			break
		}
		fmt.Println(i)
		i++
	}

	for i:=0; i<10; i++{
		fmt.Println(i)
	}

}