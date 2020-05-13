package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// e.g. if Alice's happiness increases by 23 when she's next to Bob, then happinessMap["Alice"]["Bob"] = 23
var happinessMap map[string]map[string]int
var persons []string

func parseInput() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	lastPerson := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, " ")

		person := values[0]
		increase := false
		if values[2] == "gain" {
			increase = true
		}
		amount, _ := strconv.Atoi(values[3])
		neighbor := strings.TrimRight(values[10], ".")

		// new person, initialize nested map and add to person list
		if person != lastPerson {
			happinessMap[person] = make(map[string]int)
			persons = append(persons, person)
			lastPerson = person
		}

		if !increase {
			amount *= -1
		}

		happinessMap[person][neighbor] += amount
	}
}

// https://stackoverflow.com/questions/30226438/generate-all-permutations-in-go
func permutations(arr []string) [][]string {
	var helper func([]string, int)
	res := [][]string{}

	helper = func(arr []string, n int) {
		if n == 1 {
			tmp := make([]string, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func computeHappiness() int {
	maxHappiness := -9999

	numPersons := len(persons)
	for _, perm := range permutations(persons) {
		currHappiness := 0
		for i := 0; i < numPersons; i++ {
			person := perm[i]
			leftN := perm[(i-1+numPersons)%numPersons]
			rightN := perm[(i+1)%numPersons]
			currHappiness += happinessMap[person][leftN]
			currHappiness += happinessMap[person][rightN]
		}

		if currHappiness > maxHappiness {
			maxHappiness = currHappiness
		}
	}

	return maxHappiness
}

func main() {
	happinessMap = make(map[string]map[string]int)

	parseInput()

	// part 1
	fmt.Printf("Part1: %v\n", computeHappiness())

	// part 2
	happinessMap["Me"] = make(map[string]int)
	for _, p := range persons {
		happinessMap["Me"][p] = 0
	}
	persons = append(persons, "Me")

	fmt.Printf("Part2: %v\n", computeHappiness())
}
