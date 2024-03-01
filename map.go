package main

import "fmt"

func main() {
	var myMap map[string]int = make(map[string]int)
	fmt.Println(myMap)

	var myMap2 = map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(myMap2)
	fmt.Println(myMap2["one"])
	fmt.Println(myMap2["four"])

	var num, ok = myMap2["five"]

	if ok {
		fmt.Println(num)
	} else {
		fmt.Println("five is not in the map")
	}

	delete(myMap2, "one")
	fmt.Println(myMap2)


}