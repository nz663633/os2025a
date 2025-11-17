package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func GetFloats(fileName string) ([3]float64, error) {
	var numbers [3]float64
	file, err := os.Open(fileName)
	if err != nil {
		return numbers, err
	}
	i := 0
	scanner := bufio.NewScanner(file) // 파일을 한 줄씩 읽기 쉽게 해줌
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		i++ // 다음 인덱스로 이동
	}
	err = file.Close()
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	return numbers, nil
}

func main() {
	weights, err := GetFloats("meatWeight.txt")
	if err != nil {
		log.Fatal(err)
	}
	hap := 0.0
	for i := 0; i < len(weights); i++ {
		hap = hap + weights[i]
	}
	weeks := float64(len(weights))
	fmt.Println("주차별 고기 소비량 평균: ", hap/weeks)
}
