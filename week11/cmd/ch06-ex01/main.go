package main

import "fmt"

func main() {
	subjects := []string{"Go", "", "Python"} // initialized by slice literal

	for _, subject := range subjects {
		fmt.Println(subject)
	}
}
