package main

import "fmt"


// Variables in Go
//  --String boolean int float64 uint8 int8 int16 int32 int64 float32 float64 complex.

//    var variableName variableType = value
 

const token1 = "1234567890abcdef"
// const token1 :="1234567890abcdef"  cannot use := in global scope, must be inside a function /  method


func main(){
	fmt.Println("Variables in Go")
	var username string= "john_doe"
	var age int = 30	
	fmt.Println("Username:", username)
	fmt.Printf("variable type: %T",username)
	fmt.Println("\nAge",age)
	fmt.Printf("variable type: %T", age)


	var isActive bool = true
	fmt.Println("\nIs Active:", isActive)
	fmt.Printf("variable type: %T", isActive)

	var newAge uint8= 255
	fmt.Println("\nNew Age:", newAge)
	fmt.Printf("variable type: %T", newAge)

	var float1 float64 =259.4554454543
	fmt.Println("\nFloat1:", float1)
	fmt.Printf("variable type: %T", float1)

	// defining multiple variables
	var firstName, lastName string = "John", "Doe"	
	fmt.Println(lastName,firstName)

	// no var keyword
	lastName2 := "Smith"
	fmt.Println(lastName2)


	// implict type inference - automatically defines the type based on the value
	var city = "New York"		
	fmt.Println(city)

	fmt.Println(token1)



}