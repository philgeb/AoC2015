package main

import (
	"fmt"
	"strconv"
)

const input = "1321131112"

func main() {

	// working with strings is really slow because they're UTF-8 encoded (enc/dec routines are huge performance bottlenecks)
	// []rune is just a array of uint32 which is much faster
	nextSequence := []rune(input)

	for i := 0; i < 50; i++ {

		currSequence := nextSequence
		nextSequence = nil

		currChr := ' '
		currCount := 0

		for _, c := range currSequence {
			if c == currChr {
				currCount++
			} else {
				if currChr != ' ' {
					nextSequence = append(nextSequence, []rune(strconv.Itoa(currCount))[0], currChr)
				}

				currChr = c
				currCount = 1
			}
		}

		nextSequence = append(nextSequence, []rune(strconv.Itoa(currCount))[0], currChr)
	}

	fmt.Printf("Sequence length: %v\n", len(nextSequence))
}
