package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")

	sliceInt := []int{2, 5, 7, 3, 8, 3, 7, 9}

	fmt.Println(delElement(3, sliceInt))

}

func delElement(index int, slice []int) []int {

	fresh := []int{}

	for i, k := range slice {
		if i != index {
			fresh = append(fresh, k)
		}
	}

	return fresh

}
