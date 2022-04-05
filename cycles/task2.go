package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1, num2 int
	fmt.Scan(&num1, &num2)
	//584 8111
	numstr1 := strconv.Itoa(num1)
	numstr2 := strconv.Itoa(num2)

	for _, digit := range numstr1 {

		value, _ := strconv.Atoi(string(digit))

		for _, digit1 := range numstr2 {
			value1, _ := strconv.Atoi(string(digit1))

			if value == value1 {
				fmt.Print(value, " ")
			}

		}
	}
}
