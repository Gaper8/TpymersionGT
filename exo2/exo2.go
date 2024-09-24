package main

import (
	"fmt"
)

func Ft_missing(nums []int) int {
	n := len(nums)

	result := n * (n + 1) / 2

	resultlist := 0
	for _, nombre := range nums {
		resultlist += nombre
	}
	return result - resultlist
}

func main() {
	fmt.Println(Ft_missing([]int{3, 1, 2}))
	fmt.Println(Ft_missing([]int{0, 1}))
	fmt.Println(Ft_missing([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}
