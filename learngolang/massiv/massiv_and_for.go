package main

import "fmt"

func main() {
	notes := [7]string{
		"do",
		"re",
		"mi",
	}

	// for i := 0; i < len(notes); i++ {
	// 	fmt.Println(i, " ", notes[i])
	// }
	for index, value := range notes {
		fmt.Println(index, " ", value)
	}
}
