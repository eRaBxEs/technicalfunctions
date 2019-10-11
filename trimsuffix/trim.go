package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println("Hello, 世界")
	data := []byte{47, 62, 60, 47, 103, 62, 60, 47, 115, 0, 0, 0, 0, 0, 0}
	fmt.Println(data)                              // print the value of data without trimming
	fmt.Println(bytes.TrimSuffix(data, []byte{0})) // print the value after trimming using TrimSuffix
	trimed := bytes.TrimRightFunc(data, func(r rune) bool {
		if r == 0 {
			return true
		}

		return false
	})
	// print the value after trimming using TrimRightFunc
	fmt.Println(trimed)

}
