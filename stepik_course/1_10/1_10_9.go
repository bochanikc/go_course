package main

import "fmt"

func main() {
	var (
		x    int
		p    float64
		y    int
		summ int
	)
	years := 0

	fmt.Scan(&x, &p, &y)
	p = p / 100

	for {
		years++
		summ = int(float64(x) * p)
		x = x + summ
		if x < y {
			continue
		} else if x >= y {
			fmt.Println(years)
			break
		}
	}
}
