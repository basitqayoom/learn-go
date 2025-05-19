package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	c := [...]int{1, 2, 3, 4, 5, 6, 7, 10: 8}
	fmt.Println("dcl:", c)
	fmt.Println("len:", len(c))

	var twoD [3][4]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("twoD", twoD)

	twoD = [3][4]int{
		{0, 0, 0, 0},
		{1, 2, 3, 4},
		{0, 0, 0, 0},
	}
	fmt.Println("twoD", twoD)

	// wrong
	// twoD = {
	// 	{1,2,3,4},
	// 	{1,2,3,4},
	// 	{1,2,3,4},
	// }

}
