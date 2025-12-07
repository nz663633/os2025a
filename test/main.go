package main

import (
	"fmt"
	"log"
	calendar "week13/pkg/calendar"
)

func main() {
	// today := calendar.Date{year: 2025, month: 11, day: 24}
	today := calendar.Date{}
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
	fmt.Println(today.year(), "년 ", today.Month(), "월 ", today.Day(), "일")
}
