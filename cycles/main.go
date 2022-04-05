package main

import "fmt"

func main() {
	var x, p, y int
	fmt.Scan(&x, &p, &y)
	// fmt.Println(x, p, y)
	x_f := float64(x)
	sum := x_f
	p_f := float64(p)
	y_f := float64(y)
	years := 0

	for {
		sum = sum + (sum/100)*p_f
		// 	fmt.Println(sum)
		years++

		if sum >= y_f {
			break
		}
	}
	fmt.Println(years)
}
