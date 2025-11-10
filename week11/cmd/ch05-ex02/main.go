package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/datafile"
)

func main() {
	weights, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	sum := 0.0
	// for _, number := range numbers {
	for i := 0; i < len(weights); i++ {
		sum += weights[i]
	}
	weeks := float64(len(weights))
	fmt.Println("주차별 고기 소비량 평균: ", sum/weeks)
}
