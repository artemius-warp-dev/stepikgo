package main

import "fmt"

//func print (array)

func main() {
	var workArray [10]uint8
	var poses [6]uint8

	for i := 0; i < 10; i++ {
		fmt.Scanf("%d", &workArray[i])
		// 	fmt.Print(workArray[i], " ")
	}

	for i := 0; i < 6; i++ {
		fmt.Scanf("%d", &poses[i])
		//	fmt.Print(poses[i], " ")
	}

	// 3 7 1 0 2 5 6 8
	for i := 0; i <= 4; i = i + 2 {

		index_one := poses[i]
		index_two := poses[i+1]
		//fmt.Print(index_one, index_two)

		tmp := workArray[index_one]

		workArray[index_one] = workArray[index_two]
		workArray[index_two] = tmp

	}
	fmt.Println()

	for i := 0; i < 10; i++ {

		fmt.Print(workArray[i], " ")
	}

}
