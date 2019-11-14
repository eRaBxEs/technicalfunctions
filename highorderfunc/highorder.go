package main

import "fmt"

func apply(nums []int, f func(int) int) func() {

	for i, v := range nums {
		nums[i] = f(v)
	}

	return func() {
		fmt.Println(nums)
	}

}

func main() {

	nums := []int{4, 2, 8, 12, 6, 10, 22}

	result := apply(nums, func(i int) int {

		return i / 2
	})

	result()
}
