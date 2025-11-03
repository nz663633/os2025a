package main

import (
	"fmt"
	"log"
	"week10/pkg/keyboard"
)

func main() {
	fmt.Print("점수 입력: ")
	score, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	if score >= 80 {
		fmt.Printf("%.1f점은 합격!", score)
	} else {
		fmt.Printf("%.1f는 불합격...", score)
	}
}
