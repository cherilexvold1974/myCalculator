package main

import(
	"fmt"
)

func calculate(a int, b int) (out int){
	fmt.Println(a+b)
	return a+b
}

func main() {
	fmt.Println("Hello world")
	calculate(12,2)
}

		
