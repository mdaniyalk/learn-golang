package main

import "fmt"
import "strings"

func main() {
	var myString = "rêsume"
	var indexed = myString[3]
	fmt.Println(indexed)

	fmt.Printf("%v, %T\n", indexed, indexed)

	for i, v := range myString {
		fmt.Println(i, v)
	}

	fmt.Printf("The length of the string is %d\n", len(myString))

	fmt.Println("String with runes")

	var myString2 = []rune("résumé")

	var indexed = myString2[3]
	fmt.Println(indexed)

	fmt.Printf("%v, %T\n", indexed, indexed)

	for i, v := range myString2 {
		fmt.Println(i, v)
	}

	fmt.Printf("The length of the string is %d\n", len(myString2))

	var myRune = 'a'
	fmt.Printf("%v, %T\n", myRune, myRune)

	var strSlice = []string{"a", "b", "c", "d", "e"}
	var catStr = ""
	for i := 0; i < len(strSlice); i++ {
		catStr += strSlice[i]
	}
	fmt.Println(catStr)

	fmt.Println("using strings package")

	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr2 = strBuilder.String()
	fmt.Println(catStr2)
}