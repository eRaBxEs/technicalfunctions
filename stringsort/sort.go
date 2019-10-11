package main

import (
	"bytes"
	"fmt"
)

func sortAsc(str string) string {

	runes := bytes.Runes([]byte(str))
	var temp rune

	for i := 0; i < len(runes); i++ {

		for j := i + 1; j < len(runes); j++ {
			if runes[i] > runes[j] {
				temp = runes[j]
				runes[i], runes[j] = temp, runes[i]
			}
		}
	}
	return string(runes)
}

func sortDesc(str string) string {
	runes := bytes.Runes([]byte(str))
	var temp rune

	for i := 0; i < len(runes); i++ {

		for j := i + 1; j < len(runes); j++ {

			if runes[i] < runes[j] {

				temp = runes[j]

				runes[i], runes[j] = temp, runes[i]
			}
		}
	}

	return string(runes)
}

func main() {

	str := "goodmorning"
	st := "abcdefghijklmnopqrstuvwxyz"
	s := "zyxwvutsrqponmlkjihgfedcba"

	// sort in ascending order
	newStr := sortAsc(str)
	fmt.Println(newStr)

	// sort in descending order
	newSt := sortDesc(st)
	fmt.Println(newSt)

	// sort in ascending order
	newS := sortAsc(s)
	fmt.Println(newS)
	fmt.Println("length: ", len(newS))

}
