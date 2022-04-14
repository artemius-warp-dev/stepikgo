package main

import "fmt"

func main() {
	var n uint64
	var result int

	fmt.Scan(&n)

	f := make([]uint64, n)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	var i uint64
	i = 2
	for {
		f[i] = f[i-1] + f[i-2]
		fmt.Println(f)
		if f[i] == n {
			result = int(i)
			break
		} else if i > n {
			result = -1
			break
		}
		i = i + 1
	}
	fmt.Print(result)

}

//to use append function for adding element to slice
// that was declared without len
