package main

import (
	"fmt"
	"log"
	calender "week13/pkg/calender"
)

func main() {
	// today := calender.Date{year: 2025, month: 11, day: 24}
	today := calender.Date{}
	err := today.SetYear(2025)
	if err != nil {
		log.Fatal(err)
	}
	// err = today.SetMonth(15)
	err = today.SetMonth(11)
	if err != nil {
		log.Fatal(err)
	}
	err = today.SetDay(24)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(today.Year(), "년 ", today.Month(), "월 ", today.Day(), "일")
}
