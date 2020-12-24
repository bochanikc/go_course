package main

import "fmt"

func main() {
	var numA int
	var numB int
	summ := 0

	fmt.Scan(&numA, &numB)
	for i := numA; i <= numB; i++ {
		summ += i
	}
	fmt.Print(summ)
}
