package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type ingredient struct {
	cap int
	dur int
	fla int
	tex int
	cal int
}

var ingredients []ingredient

func parseInput() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		r := regexp.MustCompile(`.*: capacity (-?\d*), durability (-?\d*), flavor (-?\d*), texture (-?\d*), calories (-?\d*)`)
		match := r.FindStringSubmatch(line)

		cap, _ := strconv.Atoi(match[1])
		dur, _ := strconv.Atoi(match[2])
		fla, _ := strconv.Atoi(match[3])
		tex, _ := strconv.Atoi(match[4])
		cal, _ := strconv.Atoi(match[5])

		ingr := ingredient{cap, dur, fla, tex, cal}
		ingredients = append(ingredients, ingr)
	}
}

// assumption: there are exactly 4 ingredients, sorry :(
func computeScore(exactCalories int) int {
	maxScore := 0

	for i := 0; i <= 100; i++ {
		for j := 0; j <= 100-i; j++ {
			for k := 0; k <= 100-i-j; k++ {
				l := 100 - i - j - k
				capTotal := ingredients[0].cap*i + ingredients[1].cap*j + ingredients[2].cap*k + ingredients[3].cap*l
				durTotal := ingredients[0].dur*i + ingredients[1].dur*j + ingredients[2].dur*k + ingredients[3].dur*l
				flaTotal := ingredients[0].fla*i + ingredients[1].fla*j + ingredients[2].fla*k + ingredients[3].fla*l
				texTotal := ingredients[0].tex*i + ingredients[1].tex*j + ingredients[2].tex*k + ingredients[3].tex*l

				var fullScore int
				if capTotal < 0 || durTotal < 0 || flaTotal < 0 || texTotal < 0 {
					fullScore = 0
				} else {
					fullScore = capTotal * durTotal * flaTotal * texTotal
				}

				calTotal := ingredients[0].cal*i + ingredients[1].cal*j + ingredients[2].cal*k + ingredients[3].cal*l

				if fullScore > maxScore {
					if exactCalories <= 0 || calTotal == exactCalories {
						maxScore = fullScore
					}
				}
			}
		}
	}

	return maxScore
}

func main() {

	parseInput()

	fmt.Printf("Part1: %v\n", computeScore(-1))
	fmt.Printf("Part2: %v\n", computeScore(500))
}
