package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var now time.Time = time.Now()
	var day int = now.Day()
	fmt.Println(day)
	univ := "Go Inha$!"
	changer := strings.NewReplacer("$", "!")
	changed := changer.Replace(univ)
	fmt.Println(changed)
}
