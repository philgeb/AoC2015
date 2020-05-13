package main

import (
	"fmt"
)

const input = "vzbxkghz"

func incrString(str []rune) {
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == 'z' {
			str[i] = 'a'
		} else {
			str[i]++
			return
		}
	}
}

func isValid(pass []rune) bool {
	ruleOneOk := false
	ruleThreeOk := false

	firstPairChar := ' '
	prev := ' '
	prev2 := ' '

	for _, c := range pass {
		// rule 1
		if prev == c-1 && prev2 == c-2 {
			ruleOneOk = true
		}

		// rule 2
		if c == 'i' || c == 'o' || c == 'l' {
			return false
		}

		// rule 3
		if c == prev {
			if firstPairChar == ' ' {
				firstPairChar = c
			} else {
				if c != firstPairChar {
					ruleThreeOk = true
				}
			}
		}

		prev2 = prev
		prev = c
	}

	return ruleOneOk && ruleThreeOk
}

func main() {

	curr := []rune(input)

	for {
		incrString(curr)
		if isValid(curr) {
			fmt.Printf("Part1: %v\n", string(curr))
			break
		}
	}

	for {
		incrString(curr)
		if isValid(curr) {
			fmt.Printf("Part2: %v\n", string(curr))
			break
		}
	}
}
