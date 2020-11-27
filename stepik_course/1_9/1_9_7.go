package main

import "fmt"

func main()  {
	var num int
	fmt.Scan(&num)
	switch {
	case num >= 0 && num <= 9:
		fmt.Println(num)
	case num >= 10 && num <= 99:
		num = num / 10
		fmt.Println(num)
	case num >= 100 && num <= 999:
		num = num / 100
		fmt.Println(num)
	case num >= 1000 && num <= 9999:
		num = num / 1000
		fmt.Println(num)
	case num >= 10000 && num <= 99999:
		num = num / 10000
		fmt.Println(num)
	}
}
