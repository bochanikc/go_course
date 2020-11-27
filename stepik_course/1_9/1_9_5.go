package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	switch {
	case num < 0:
		fmt.Println("Число отрицательное")
	case num == 0:
		fmt.Println("Ноль")
	case num > 0:
		fmt.Println("Число положительное")
	}
}