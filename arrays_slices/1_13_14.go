package main

import "fmt"

func main() {

	var num int
	var digit int
	var remove uint
	fmt.Scan(&num, &remove)
	var num_array []int
	var sign_flag bool
	if num < 0 {
		sign_flag = true
		num = num * -1
	} else {
		sign_flag = false

	}
	for {

		digit = num % 10
		num = num / 10
		if digit != int(remove) {
			num_array = append(num_array, digit)

		}
		if num == 0 {
			break
		}
		//fmt.Println(digit)
	}
	if sign_flag {
		fmt.Print("-")
	}
	for i := len(num_array) - 1; i >= 0; i-- {

		fmt.Printf("%d", num_array[i])

	}

}
