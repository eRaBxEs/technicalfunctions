package main

import "fmt"

func maxInt(slice []int) int {
	max := 0
	for _, k := range slice {
		if max < k {
			max = k
		}
	}

	return max
}

func main() {
	sliceInt := []int{2, 5, 7, 3, 8, 3, 7, 9}

	fmt.Println(maxInt(sliceInt))

}
