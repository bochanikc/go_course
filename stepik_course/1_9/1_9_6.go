package main

import "fmt"

func main()  {
	var num int
	fmt.Scan(&num)

	var first_num int = num / 100
	var second_num int = num / 10 % 10
	var third_num int = num % 10
	if first_num != second_num && first_num != third_num && second_num != third_num {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
