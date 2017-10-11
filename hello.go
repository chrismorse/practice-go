package main

import "fmt"


// to run this from command line - go run hello.go
func main() {
	
	fmt.Println("hello, world\n")
	
	// numbers
	var age int = 40
	var favNum float64 = 1.6180339
	randNum := 1
	fmt.Println(age, favNum, randNum)
	
	// constant
	const pi float64 = 3.14159265
	fmt.Println(pi)
	fmt.Printf("%.3f \n", pi)   //print float - only 3 decimals

	// multiple vars
	var (
		varA = 2
		varB = 3
	)
	fmt.Println(varA, varB)

	// strings
	var myName string = "Chris Morse"
	fmt.Println(myName, "length->", len(myName))

	// bool
	var isOver10 bool = true
	fmt.Println(isOver10)

	// logical operators
	fmt.Println("true && false = ", true && false)
	fmt.Println("true || false = ", true || false)
	fmt.Println("!true = ", !true)
	fmt.Println()

	// for loops
	i := 1
	for i < 10 {
		fmt.Println(i);
		i++
	}
	fmt.Println()
	
	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}
	fmt.Println()
	
	// if statement
	yourAge := 18

	if yourAge >= 16 {
		fmt.Println("You can drive")
	} else {
		fmt.Println("Sorry, you can't drive")
	}
	fmt.Println()
	
	// switch
	switch yourAge {
		case 16 : fmt.Println("Go Drive")
		case 18 : fmt.Println("Go Vote")
		default : fmt.Println("Go Have Fun")
	}
	fmt.Println()
	
	// arrays
	var favNums[2] float64
	favNums[0] = 163
	favNums[1] = 1.618
	fmt.Println(favNums[1])
	fmt.Println()
	
	favNums2 := [5]float64 {1,2,3,4,5}
	for i, value := range favNums2 {      // iterate through an array!
		fmt.Println(i,"->",value)
	}


}

