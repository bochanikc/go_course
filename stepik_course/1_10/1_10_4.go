package main

import "fmt"

func main() {
	var num int
	var dig int
	summ := 0

	fmt.Scan(&num)

	for i := 0; i < num; i++ {
		fmt.Scan(&dig)
		if dig >= 10 && dig <= 100 && (dig%8) == 0 {
			summ += dig
		}
	}
	fmt.Print(summ)
}
