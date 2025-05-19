package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")

	total := 0

	for _, num := range nums {
		total += num
	}

	fmt.Println("Total:", total)
}

func main() {
	sum(1, 2, 4)
	sum(1, 2, 4, 6, 7)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
	sum()
}
