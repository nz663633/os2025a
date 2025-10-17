package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// rand.Seed(55)
	dice := rand.Intn(6) + 1 // 1 ~ 6
	fmt.Println(dice)
}
