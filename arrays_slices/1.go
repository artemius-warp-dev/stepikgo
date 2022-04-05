package main

import "fmt"

func main() {
	var workArray [10]uint8
	var poses [6]uint8

	for i := 0; i < 10; i++ {
		fmt.Scanf("%d", &workArray[i])
		fmt.Print(workArray[i], " ")
	}

	fmt.Println()
	
	for i := 0; i < 6; i++ {
		fmt.Scanf("%d", &poses[i])
		fmt.Print(poses[i], " ")
	}

	for _i, val := range poses {
	  
	}

}
