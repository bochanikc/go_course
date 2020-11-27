package main

import "fmt"

func main()  {
	var num int
	fmt.Scan(&num)
	last_num := num % 10
	fmt.Println(last_num)
}
