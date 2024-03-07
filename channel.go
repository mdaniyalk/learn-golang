package main 

import "fmt"

func main() {
	var c = make(chan int)	
	go process(c)
	fmt.Println(<-c)

	var c2 = make(chan int)
	go process2(c2)
	for i:= range c2{
		fmt.Println(i)
	
	}
}

func process(c chan int) {
	c <- 123
}


func process2(c chan int) {
	for i:=0; i<10; i++{
		c <- i
	}
	close(c)
}