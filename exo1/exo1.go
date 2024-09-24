package main

import (
	"fmt"
	"sort"
)

func Ft_coin(coins []int, amount int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(coins)))

	count := 0

	for _, coin := range coins {
		for amount >= coin {
			amount -= coin
			count++
		}
	}

	if amount > 0 {
		return -1
	}

	return count
}

func main() {
	fmt.Println(Ft_coin([]int{1, 2, 5}, 16))
	fmt.Println(Ft_coin([]int{2}, 3))
	fmt.Println(Ft_coin([]int{1}, 0))
}
