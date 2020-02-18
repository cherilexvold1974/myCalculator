package main

import(
	"fmt"
)

func calculate(a int, operator string, b int){
	if operator == "+"{
	     fmt.Println(a+b)
	} else if operator == "-"{
		 fmt.Println(a-b)
	} else if operator == "*"{
		fmt.Println(a*b)
	} else if operator == "/"{
		 if b == 0{
	 	 fmt.Println("error: b cannot = 0")
		 } else {
		 fmt.Println(a/b)
		 }	
	} else if operator != "+"||operator != "-"||operator != "*"||operator !=	  "/" {
		fmt.Println("error: No such operator")
	}
}

func main() {
	fmt.Println("Hello world")
	calculate(12, "+", 2)
	calculate(33, "-", 5)
	calculate(6, "*", 3)
	calculate(12, "/", 4)
	calculate(5, "/", 0)
        calculate(5, "😁", 1)
	calculate(5, "😁", 0)
	calculate(12, "/", 7)
	calculate(3, "-", 5)
	calculate(5, "/", 1)
}		
