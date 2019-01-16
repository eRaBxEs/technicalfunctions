package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	d := []int{3, 4, 5, 6, 7}

	fmt.Println("Initial arrangement of d:", d)

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]

	}

	fmt.Println("Arrangement of d after swapping:", d)

}
