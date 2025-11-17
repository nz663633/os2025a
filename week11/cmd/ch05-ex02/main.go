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
	weeks := float64(len(weights)) // 나눌 때 분자와 분모의 타입이 같아야 계산 가능
	fmt.Println("주차별 고기 소비량 평균: ", sum/weeks)
}
