package main

import (
	"errors"
	"fmt"
	"time"
	// "unicode/utf8"
)

func main() {
	/*
	var intNum int16 = 32767
	intNum = intNum + 1
	
	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = float32(intNum32) + floatNum32


	var intNum1 int = 3
	var intNum2 int = 2

	fmt.Println(intNum1/intNum2)
	fmt.Println(intNum1%intNum2)

	var myString string = "Hello " + "" + "fen"

	fmt.Println(myString)

	fmt.Println("Result: ", result)

	fmt.Println("Hello") 

	fmt.Println(len("test"))

	fmt.Println(utf8.RuneCountInString("N"))

	var myBoolean bool
	fmt.Println(myBoolean)

	var myVar = "text"
	fmt.Println(myVar)

	myVar1 := "text1"
	fmt.Println(myVar1)

	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	const myName string = "man minh"
	fmt.Println(myName)

	printMe("Man Minh")

	var numberator int = 10
	var denominator int = 5
	var result1, remainder, err = intDivision(numberator, denominator)*/
	
	/*
	if (err != nil) {
		fmt.Println(err.Error())
	} else if (remainder == 0) {
		fmt.Printf("The result of the int division is %v", result1)
	} else {
		fmt.Printf("Result1: %v. Remainder: %v", result1, remainder)
	}
	switch {
		case err != nil:
			fmt.Println(err.Error())
		case remainder == 0:
			fmt.Printf("The result of the int division is %v", result1)
		default:
			fmt.Printf("Result1: %v. Remainder: %v", result1, remainder)
	}
	
	switch remainder {
	case 0:
		fmt.Println("\nThe division was exact")
	case 1, 2:
		fmt.Println("\nThe division was close")
	default:
		fmt.Println("\nThe division was not close")
	}
		*/

	// var intArr [3]int32
	
	// intArr[1] = 123
	// fmt.Println(intArr[0])
	// fmt.Println(intArr[1:3])

	// fmt.Println(&intArr[0])
	// fmt.Println(&intArr[1])
	// fmt.Println(&intArr[2])

	/*
	intVar := [...]int32{1,2,3,4,5}
	// intVar := [3]int32{1,2,5} 
	fmt.Println(intVar)

	var intSlice []int32 = []int32{1,2,3}
	fmt.Printf("The length is %v with capacity %v \n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 4)
	fmt.Printf("The length is %v with capacity %v \n", len(intSlice), cap(intSlice))
	var intSlice1 []int32 = []int32{5, 6}
	intSlice = append(intSlice, intSlice1...)
	fmt.Printf("The length is %v with capacity %v \n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)
	
	var intSlice2 []int32 = make([]int32, 3, 6)
	fmt.Println(intSlice2)
	intSlice2 = append(intSlice2, 1,2,3,4)
	fmt.Println(intSlice2)
	intSliceFirstEl := []int32{5, 6}
	copy(intSlice2, intSliceFirstEl)
	fmt.Println(intSlice2)
	*/

	/*
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)
	var myMap1 = map[string]uint8{"Man": 21, "Minh": 22}
	fmt.Println(myMap1["Man"])
	var age, ok = myMap1["Thienn"]
	
	if ok {
		fmt.Printf("The age id %v", age)
	} else {
		fmt.Println("Invalid Name")
	}

	for name, age:= range myMap1{
		fmt.Printf("Name: %v , Age: %v\n", name, age)
	}

	intArr := [...]int32{1,2,3,4,5}
	for i,v := range intArr{
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}
	*/
	

	// var i int = 0

	// for {
	// 	if i >= 10 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	// for i := 0; i<10; i++{
	// 	fmt.Println(i)
	// }
	


	var n int = 1000000
	var testSlice = []int{}
	var testSlice2 = make([]int, 0, n)

	fmt.Printf("Total time without preallocation: %v \n", timeLoop(testSlice, n))
	fmt.Printf("Total time with preallocation: %v \n", timeLoop(testSlice2, n))
}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numberator int, denominator int) (int, int, error) {
	var err error
	if (denominator == 0) {
		err = errors.New("Cannot Divide by Zero")
		return 0, 0, err
	}
	var result int = numberator / denominator
	var remainder int = numberator % denominator
	return result, remainder, err
}