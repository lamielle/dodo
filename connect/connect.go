package main

import (
	"fmt"
	"strconv"
)

// 1 2 3
// 4 5 6
// 7 8 9
//   0

// nums is a 10 x 10 adjacency matrix that represents the connectivity of the numbers of a phone's dial pad.

func connected(digits string) bool {
	adj := [][]int{[]int{0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, // 0
		[]int{0, 0, 1, 0, 1, 0, 0, 0, 0, 0}, // 1
		[]int{0, 1, 0, 1, 0, 1, 0, 0, 0, 0}, // 2
		[]int{0, 0, 1, 0, 0, 0, 1, 0, 0, 0}, // 3
		[]int{0, 1, 0, 0, 0, 1, 0, 1, 0, 0}, // 4
		[]int{0, 0, 1, 0, 1, 0, 1, 0, 1, 0}, // 5
		[]int{0, 0, 0, 1, 0, 0, 1, 0, 0, 1}, // 6
		[]int{0, 0, 0, 0, 1, 0, 0, 0, 1, 0}, // 7
		[]int{1, 0, 0, 0, 0, 1, 0, 1, 0, 1}, // 8
		[]int{0, 0, 0, 0, 0, 0, 1, 0, 1, 0}, // 9
	}

	nums := []int{}
	for _, digit := range digits {
		num, err := strconv.Atoi(string(digit))
		if err != nil {
			return false
		}
		nums = append(nums, num)
	}
	if 0 == len(nums) {
		return true
	}
	fmt.Println(nums)
	curr := nums[0]
	for _, num := range nums[1:] {
		fmt.Println(curr, num, adj[curr][num])
		if adj[curr][num] == 0 {
			return false
		}
		curr = num
	}
	return true
}

func main() {
	fmt.Println("empty string: ", connected(""))
	fmt.Println("1: ", connected(""))
	fmt.Println("1234: ", connected("1234"))
	fmt.Println("1236980: ", connected("1236980"))
}
