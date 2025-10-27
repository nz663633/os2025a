package main

import "fmt"

func swap(first *int, second *int) {
	temp := 0
	temp = *first
	*first = *second
	*second = temp
	fmt.Println(*first, *second)
}

func main() {
	var a, b int = 10, 20
	fmt.Println(a, b)
	swap(&a, &b) // call by pointer
	fmt.Println(a, b)
	// fmt.Printf("%.2f\n", math.Sqrt(-25.0))
}
