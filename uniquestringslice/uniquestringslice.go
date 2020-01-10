package main

import (
	"fmt"
	"strings"
)

func uniqueNames(a, b []string) []string {
	var result []string

	for _, m := range a {
		result = append(result, m)
	}

	for _, n := range b {

		result = append(result, n)

	}

	temp := len(result)
	fmt.Printf("%v\n", result)
	//var count int;
	for i := 0; i < len(result); i++ {

		for j := i + 1; j < len(result); j++ {

			if strings.EqualFold(result[i], result[j]) {

				fmt.Println(result[i])
				result[i] = result[len(result)-1] // Copy last element to index i.
				result[len(result)-1] = ""        //  Erase last element (write zero value)
				result = result[:len(result)-1]   // Truncate slice.

				// since there is a change in the starting length of the result slice
				// Reset j back to i + 1 to start again
				j = i + 1
			}

		}

		// Reset i back to zero to start again if there is a change in the starting length of the result slice
		if temp > len(result) {
			temp = len(result) // reset temp
			i = 0
		}

	}

	return result
}

func main() {
	// should print Ava, Emma, Olivia, Sophia
	fmt.Println(uniqueNames(
		[]string{"James", "Bolu", "Ava", "Emma", "Olivia", "Joshua", "Grace", "Blake"},
		[]string{"James", "Bolu", "Olivia", "Sophia", "Emma", "Joshua", "Grace", "Blake"}))
}
