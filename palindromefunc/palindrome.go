package main

import (
	"bytes"
	"fmt"
	"strings"
)

func scanWord(str string) bool {
	str = strings.ToLower(str)
	var newStr string
	runes := bytes.Runes([]byte(str))
	var newRune []rune

	for i := len(runes) - 1; i >= 0; i-- {

		newRune = append(newRune, runes[i])

	}

	newStr = string(newRune)

	if str == newStr {
		return true
	}
	return false
}

func main() {
	word := "Madam"

	fmt.Println(scanWord(word))
}
