package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var ticketTapeStats map[string]int

func part1() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		r := regexp.MustCompile(`Sue (\d*): (.*): (\d*), (.*): (\d*), (.*): (\d*)`)
		match := r.FindStringSubmatch(line)

		sueNr := match[1]
		stat1 := match[2]
		amount1, _ := strconv.Atoi(match[3])
		stat2 := match[4]
		amount2, _ := strconv.Atoi(match[5])
		stat3 := match[6]
		amount3, _ := strconv.Atoi(match[7])

		if ticketTapeStats[stat1] == amount1 &&
			ticketTapeStats[stat2] == amount2 &&
			ticketTapeStats[stat3] == amount3 {
			fmt.Printf("Part1: %v\n", sueNr)
		}
	}
}

func part2() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
nextSue:
	for scanner.Scan() {
		line := scanner.Text()
		r := regexp.MustCompile(`Sue (\d*): (.*): (\d*), (.*): (\d*), (.*): (\d*)`)
		match := r.FindStringSubmatch(line)

		var things []string
		var amounts []int

		sueNr := match[1]

		things = append(things, match[2], match[4], match[6])

		amount1, _ := strconv.Atoi(match[3])
		amount2, _ := strconv.Atoi(match[5])
		amount3, _ := strconv.Atoi(match[7])

		amounts = append(amounts, amount1, amount2, amount3)

		for i, t := range things {
			if t == "cats" || t == "trees" {
				if amounts[i] <= ticketTapeStats[t] {
					continue nextSue
				}
			} else if t == "pomeranians" || t == "goldfish" {
				if amounts[i] >= ticketTapeStats[t] {
					continue nextSue
				}
			} else if amounts[i] != ticketTapeStats[t] {
				continue nextSue
			}
		}

		fmt.Printf("Part1: %v\n", sueNr)
	}
}

func main() {

	ticketTapeStats = make(map[string]int)

	ticketTapeStats["children"] = 3
	ticketTapeStats["cats"] = 7
	ticketTapeStats["samoyeds"] = 2
	ticketTapeStats["pomeranians"] = 3
	ticketTapeStats["akitas"] = 0
	ticketTapeStats["vizslas"] = 0
	ticketTapeStats["goldfish"] = 5
	ticketTapeStats["trees"] = 3
	ticketTapeStats["cars"] = 2
	ticketTapeStats["perfumes"] = 1

	part1()
	part2()

}
