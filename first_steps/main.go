package main

import (
	"errors"
	"fmt"
)

func main() {
	var name string = "Aman"
	fmt.Println("Hello, " + name + "!")


	var intNum int64 = 67
	fmt.Println("intNum: ", intNum)

	var floatNum float64 = 6.7
	fmt.Println("floatNum: ", floatNum)

	if intNum > 50 {
		fmt.Println("intNum is greater than 50")
	} else {
		fmt.Println("intNum is less than 50")
	}

	nameNew := "Aman"
	println("nameNew: ", nameNew)

	const pi float64 = 3.1415

	fmt.Println("Dividing 10 by 3, 5 by 0")
	q1, r1, err := intDivider(10, 3)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	} else {
		fmt.Println("Quotient: ", q1)
		fmt.Println("Remainder: ", r1)
	}
	q2, r2, err := intDivider(5, 0)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	} else {
		fmt.Println("Quotient: ", q2)
		fmt.Println("Remainder: ", r2)
	}

	if sayBye(name) {
		println("Closing the program")
	}

	// array 
	arr := [...]int{1, 2, 3}
	fmt.Println(&arr[0])
	fmt.Println(&arr[1])
	fmt.Println(&arr[2])

	arr2 := make([]int, 3, 5)
	arr2[0] = 1
	arr2[1] = 2
	arr2[2] = 3

	fmt.Println(arr2)

	// Slice
	slice := arr2[0:2]
	fmt.Println(slice)

	// Map
	mymap := make(map[string]int)
	mymap["Aman"] = 1
	mymap["Aman"] = 2
	mymap["Aman"] = 3
	fmt.Println(mymap)

	random, exists := mymap["Random"]
	if exists {
		fmt.Println(random)
	} else {
		fmt.Println("Random key does not exist")
	}
}

func sayBye(name string) bool {
	fmt.Println("Bye, " + name + "!")
	return true
}

func intDivider(a int, b int) (int, int, error) {
	var err error
	if b == 0 {
		err = errors.New("division by zero")
		return 0, 0, err
	}
	quotient := a / b
	remainder := a % b
	return quotient, remainder, nil
}