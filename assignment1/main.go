package main

import (
	"assignment1/constants"
	"fmt"
	"reflect"
)

var sum int
var sub int
var mul int
var quo float32

func main() {
	fmt.Printf("Number1 is of type %T\n", constants.Number1)
	fmt.Println("Number2 is of type",reflect.TypeOf(constants.Number2))
	fmt.Println(constants.Phrase)
	fmt.Println("Phrase is of type",reflect.TypeOf(constants.Phrase))
	sum, sub, mul, quo = calculate(constants.Number1, constants.Number2)
	fmt.Println("Sum:",sum,"Difference:", sub,"Product:", mul,"Quotient:", quo)
}

func calculate(a , b int) (sum, sub, mul int, quo float32){
	sum = (a + b)
	sub = (a - b) 
	mul = (a * b)
	quo = (float32(a) / float32(b))
	return
}