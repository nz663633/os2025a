package main

import "fmt"

type subscriber struct {
	name  string
	price float32
}

func applyPrice(s *subscriber) {
	s.price = 10000
	s.name = "Park Inha"
}

func main() {
	var s1 subscriber
	// s1.name = "Kim Inha"
	applyPrice(&s1)
	fmt.Println(s1.name, s1.price)
}
