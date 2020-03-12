package main

import(
	"fmt"
	"math"
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
	} else {
		fmt.Println("error: No such operator")
	}
}

func cosine(operator string, x float64) {
	if operator == "Cos" {
 		fmt.Println(math.Cos(x))
	}
}

func anticosine(operator string, y float64) {
	if operator == "Acos" {
		fmt.Println(math.Acos(y))
	}
}

func sine (operator string, z float64) {
	if operator == "Sin" {
		fmt.Println(math.Sin(z))
	}
}

func antisine (operator string, q float64) {
	if operator == "Asin" {
		fmt.Println(math.Asin(q))
	}
}

func pwr(a float64, b float64){
	c := math.Pow(a, b)
        	 fmt.Println(" a power of b is : ", c)
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
	pwr(5, 2)
	pwr(-5, 2)
	pwr(5, -2)
	cosine("Cos", 25)
	anticosine("Acos", 0.5)
	anticosine("Acos", 0.25)
	anticosine("Acos", 0.90630779)
	sine("Sin", 25)
	antisine("Asin", 0.5)
	antisine("Asin", 0.25)
        antisine("Asin", 0.90630779)
}		
