package main

import "fmt"
import "unicode/utf8"

func main(){
	var intNum int = 32767
	intNum = intNum + 1
	fmt.Println(intNum)
	
	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1/intNum2)
	fmt.Println(intNum1%intNum2)

	var myString string = "Hello" + " " + "World"
	fmt.Println(myString)

	fmt.Println(utf8.RuneCountInString("Hello World"))

	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBool bool = true
	fmt.Println(myBool)

	var myByte byte = 255
	fmt.Println(myByte)

	var intNum3 rune
	fmt.Println(intNum3)

	var myComplex complex64 = 1 + 2i
	fmt.Println(myComplex)

	myVar := "text"
	fmt.Println(myVar)	

	var var1, var2 = 1, 2
	fmt.Println(var1, var2)

	var var3, var4 int = 3, 4
	fmt.Println(var3, var4)

	var var5, var6 := 5, 6
	fmt.Println(var5, var6)

	const myConst string = "constant"
	fmt.Println(myConst)

	const pi float64 = 3.14159
	fmt.Println(pi)

}
