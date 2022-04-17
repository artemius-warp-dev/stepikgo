package main

import "fmt"

func main() {
	var n, value uint64
	var result int

	fmt.Scan(&n)

	//	f := make([]uint64, 0, n+2)
	var f []uint64
	if n < 2 {

		result = int(n)
	} else {
		f = append(f, 0)
		f = append(f, 1)
		var i uint64
		i = 2
		for {
			value = f[i-1] + f[i-2]
			f = append(f, value)
			// fmt.Println(f)
			if f[i] == n {
				result = int(i)
				break
			} else if i > n {
				result = -1
				break
			}
			i = i + 1
		}

	}

	fmt.Print(result)

}

//to use append function for adding element to slice
// that was declared without len
