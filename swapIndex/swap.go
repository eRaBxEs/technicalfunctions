package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	d := []int{3, 4, 5, 6, 7}

	fmt.Println("Initial arrangement of d:", d)

	// Need to ensure that we use a seed that is constantly changing e.g time
	// and that is what enables us to get a different sequence of random number

	source := rand.NewSource(time.Now().UnixNano())

	//After obtaining a source of type Source, get a Rand by using func New from package rand
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]

	}

	fmt.Println("Arrangement of d after swapping:", d)

}
