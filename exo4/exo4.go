package main

import (
	"fmt"
)

func Ft_profit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	min := prices[0]
	max := 0

	for _, prix := range prices {
		if prix < min {
			min = prix
		}
		profit := prix - min
		if profit > max {
			max = profit
		}
	}

	return max
}

func main() {
	fmt.Println(Ft_profit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(Ft_profit([]int{7, 6, 4, 3, 1}))
}
