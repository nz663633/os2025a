package main

import "fmt"

func main() {
	var arrayBool [3]bool
	var arrayInt [3]int
	fmt.Println(arrayBool[1]) // zero value (false)
	arrayInt[1]++
	arrayInt[1] = arrayInt[1] + 1
	fmt.Println(arrayInt[1]) // zero value + 2
}
