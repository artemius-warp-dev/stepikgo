package main

import "fmt"

func main() {

	var N int

	Amount := 0

	fmt.Scanf("%d", &N)
	array := make([]int, N)
	if N >= 1 && N <= 100 {

		for i := 0; i < N; i++ {
			fmt.Scanf("%d", &array[i])
		}

		for i := 0; i < N; i++ {
			if array[i] > 0 {
				Amount = Amount + 1
			}

		}
		fmt.Println(Amount)
	}
}
