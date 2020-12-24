package main

import "fmt"

func main() {
	var check int
	var max int
	count := 1
	fmt.Scan(&check)
	max = check
	for check != 0 {
		fmt.Scan(&check)
		if check > max {
			max = check
			count = 1
		} else if check == max {
			count += 1
		}
	}
	fmt.Print(count)
}
