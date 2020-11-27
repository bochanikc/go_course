package main

import "fmt"

func main()  {
	var num int
	fmt.Scan(&num)
	var (
		num1 int
		num2 int
		num3 int
		num4 int
		num5 int
		num6 int
	)
	num1 = num / 100000
	num2 = num / 10000 % 10
	num3 = num / 1000 % 10
	num4 = num / 100 % 10
	num5 = num / 10 % 10
	num6 = num % 10

	var (
		sum1 int
		sum2 int
	)

	sum1 = num1 + num2 + num3
	sum2 = num4 + num5 + num6

	if sum1 == sum2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
