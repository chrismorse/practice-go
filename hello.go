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
	fmt.Println()
	
	// slice ... like an array but leave out size
	numSlice := []int {5,4,3,2,1}
	numSlice2 := numSlice[3:5]   // take a slice of a slice
	fmt.Println(numSlice2[0])
	fmt.Println(numSlice[2:])

	numSlice3 := make([]int, 5, 10)     // default the first 5 to zero, max size is 10
	fmt.Println(numSlice3)
	copy(numSlice3, numSlice)
	fmt.Println(numSlice3)
	fmt.Println()
	

	// map - collection of key value pairs
	presAge := make(map[string] int)
	presAge["TheodoreRoosevelt"] = 42
	presAge["John F. Kennedy"] = 43
	fmt.Println(presAge)
	delete(presAge, "John F. Kennedy")
	fmt.Println(presAge)
	fmt.Println()
	



	// call a function
	listNums := []float64{1,2,3,4,5}
	fmt.Println("Sum", addThemUp(listNums))
	
	// return 2 values :-)
	num1, num2 := next2Values(5)
	fmt.Println(num1,num2)
	
	// accepts multiple values
	fmt.Println(subtractThem(1,2,3,4,5))
	fmt.Println()

	// closures
	num3 := 3
	doubleNum := func() int {
		num3 *= 2
		return num3
	}
	fmt.Println(doubleNum())
	fmt.Println(doubleNum())
	fmt.Println(num3)
	fmt.Println()
	
	// recursion
	fmt.Println(factorial(5))
	
	// defer - it will be called after main is finished.  used a lot for cleanup.
	defer printTwo()
	printOne()



	fmt.Println()
}






// functions
func addThemUp(numbers []float64) float64 {
	sum := 0.0
	for _,value := range numbers {
		sum += value
	}
	return sum
}

// function that returns 2 values
func next2Values(number int)(int, int) {
	return number+1, number+2
}

// variadic  -- multiple inputs
func subtractThem(args ...int) int {
	finalValue := 0
	for _, value := range args {
		finalValue -= value
	}
	return finalValue
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}

	return num * factorial(num - 1)
}

func printOne() { fmt.Println(1) }
func printTwo() { fmt.Println(2) }

