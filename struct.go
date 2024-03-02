package main

import "fmt"

type gasEngine struct {
	mpg uint8
	gallons uint8
}

func (e gasEngine) milesLeft() string {
	return fmt.Sprintf("%d miles left", e.mpg * e.gallons)
}

func main() {
	var myEngine gasEngine
	fmt.Println(myEngine)
	fmt.Println(myEngine.mpg, myEngine.gallons)

	var myEngine2 gasEngine = gasEngine{30, 10}
	fmt.Println(myEngine2)
	fmt.Println(myEngine2.mpg, myEngine2.gallons)
	fmt.Println(myEngine2.milesLeft())

	var myEngine3 = struct{mpg uint8; gallons uint8}{30, 10}

	fmt.Println(myEngine3)

}