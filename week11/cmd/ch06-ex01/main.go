package main

import "fmt"

func main() {
	var subjects []string
	subjects = make([]string, 3)
	subjects[0] = "Go"
	subjects[2] = "Python"

	for _, subject := range subjects {
		fmt.Println(subject)
	}
}
