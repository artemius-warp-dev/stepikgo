package main

import "fmt"

func main() {
	var N int
	var finalSlice []int

	fmt.Scanf("%d", &N)
	inputArray := make([]int, N)
	if N >= 1 && N <= 100 {

		for i := 0; i < N; i++ {
			fmt.Scanf("%d", &inputArray[i])
		}
		for i := 0; i < N; i++ {
			reminder := i % 2
			if reminder == 0 {
				finalSlice = append(finalSlice, inputArray[i])
			}
		}
		for i := 0; i < len(finalSlice); i++ {
			fmt.Print(finalSlice[i], " ")

		}

	}
}
