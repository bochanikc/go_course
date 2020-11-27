package main

import "fmt"

func main()  {
	fmt.Print("Введите градус Фаренгейта: ")
	var f int
	var c int
	fmt.Scan(&f)
	fmt.Println(f)
	c = (f - 32) * 5/9
	fmt.Println("Градусы цельсия:", c)

}
