package main

import "fmt"

func main()  {
	var deg int
	fmt.Scan(&deg)
	hours := deg / 30
	mins :=  (deg - (30 * hours)) * 2
	fmt.Println("It is", hours, "hours", mins, "minutes.")
}
