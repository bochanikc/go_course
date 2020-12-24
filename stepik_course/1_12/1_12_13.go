package main

import "fmt"

func main() {
	var num int
	var addNum int

	fmt.Scan(&num)

	var slise []int = make([]int, num)

	//fmt.Println(slise)
	//fmt.Println(len(slise))
	for i := 0; i < len(slise); i++ {
		fmt.Scan(&addNum)
		slise[i] = addNum
		//fmt.Println(slise)
	}

	fmt.Print(slise[3])
}
