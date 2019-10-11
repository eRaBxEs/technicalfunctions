package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
)

func sortRunes(str string) string {
	runes := bytes.Runes([]byte(str))
	var temp rune

	for i := 0; i < len(runes); i++ {

		for j := i + 1; j < len(runes); j++ {

			if runes[j] < runes[i] {
				temp = runes[j]

				runes[i], runes[j] = temp, runes[i]
			}
		}
	}

	return string(runes)
}

func load(fname string) ([]string, error) {

	if fname == "" {
		return nil, errors.New(
			"Dictionary cannot be empty",
		)
	}

	file, err := os.Open(fname)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, err

}

func main() {

	fname := "dict.txt"
	words, err := load(fname)
	if err != nil {
		fmt.Println("Unable to load file:", err)
		os.Exit(1)
	}

	anagrams := make(map[string][]string)

	for _, word := range words {
		wordSig := sortRunes(word)

		anagrams[wordSig] = append(anagrams[wordSig], word)
	}

	for k, v := range anagrams {
		fmt.Println(k, "-->", v)
	}

}
