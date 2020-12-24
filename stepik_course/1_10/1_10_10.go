package main

import "fmt"
import "strconv"

func main() {
	var num1 int
	var num2 int

	var check1 string
	var check2 string

	fmt.Scan(&num1, &num2)

	check1 = strconv.Itoa(num1)
	check2 = strconv.Itoa(num2)
	//println(check1)
	//println(string(check1[0]))
	//println(len(check1))

	for i := 0; i < len(check1); i++ {
		for j := 0; j < len(check2); j++ {
			if string(check1[i]) == string(check2[j]) {
				fmt.Println(string(check1[i]) + string(' '))
			}
		}
	}
}
