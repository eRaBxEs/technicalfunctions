package main

import (
	"fmt"
	"unicode"
)

// UserInput ...
type UserInput interface {
	Add(rune)
	GetValue() string
}

// NumericInput ...
type NumericInput struct {
	input string
}

// Add ...
func (r *NumericInput) Add(t rune) {

	if unicode.IsDigit(t) {
		s := string(t)
		r.input = r.input + s
	}

}

// GetValue ...
func (r *NumericInput) GetValue() string {

	return r.input
}

func main() {
	var input UserInput = &NumericInput{}
	input.Add('1')
	input.Add('a')
	input.Add('0')
	fmt.Println(input.GetValue())
}
