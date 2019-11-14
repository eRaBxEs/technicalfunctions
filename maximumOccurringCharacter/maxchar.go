package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {

	res := maxOccur("aaaabbbccccccc")
	fmt.Printf("Most Occurring: %s", res)
}

func maxOccur(str string) string {
	str = strings.ToLower(str)

	runes := bytes.Runes([]byte(str)) // transform the strings to a slice of rune
	var charOccur []int               // to store the no of occurrence for a character
	for i := 0; i < len(runes); i++ {
		var counter int

		for j := i + 1; j < len(runes); j++ {
			if runes[i] == runes[j] {
				counter++
			}
		}

		charOccur = append(charOccur, counter) // assign the counter as an

	}

	large := 0
	var charVal rune

	for i := 0; i < len(charOccur); i++ {
		if charOccur[i] > large {
			charVal = runes[i]
		}
	}

	return string(charVal)
}
