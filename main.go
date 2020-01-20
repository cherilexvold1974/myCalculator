package main

import(
	"fmt"
)

func calculate(a int, operator string, b int){
	if operator == "+"{
	     fmt.Println(a+b)
	}
}
                 
func main() {
	fmt.Println("Hello world")
	calculate(12, "+", 2)
}

		
