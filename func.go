package main

import "fmt"
import "errors"

func main(){
	var printValue string = "Hello World"
	printMe(printValue)

	var numerator int = 10
	var denominator int = 3
	var result, remainder, err = intDivision(numerator, denominator)
	if err != nil {
		fmt.Println(err)
	} else if remainder == 0 {
		fmt.Printf("%d divided by %d equals %d\n", numerator, denominator, result)
	} else {
		fmt.Printf("%d divided by %d equals %d with a remainder of %d\n", numerator, denominator, result, remainder)
	}
}

func printMe(printValue string){
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) int {
	var err error
	if denominator == 0 {
		err = errors.New("Division by zero")
		fmt.Println(err)
		return 0, 0, err
	}
	var result int = numerator/denominator
	var remainder int = numerator%denominator
	return result, remainder, err
}