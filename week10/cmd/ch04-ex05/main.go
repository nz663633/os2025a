package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/keyboard" // go get github.com/headfirstgo
)

func main() {
	fmt.Print("실수 입력: ")
	number, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	if number >= 80 {
		fmt.Printf("%.1f\n", number)
	}
}
