package main

import (
	"fmt"
	"go_course/learngolang/incapsulation/calendar"
	"log"
)

func main() {
	date := calendar.Date{}

	err := date.SetDay(13)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(9)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetYear(2021)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(date)
	fmt.Println(date.Day())
}
