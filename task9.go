package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var distByRoute map[string]int // e.g. distByRoute["TristramAlphaCentauri"] = 5
var nodes map[string]bool      // map instead of list, because go doesn't have a native list.contains function

var minDist = 99999999 // for part1, nitally really big
var maxDist = -1       // for part2, initally really small

func parseInput() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		reg := regexp.MustCompile(`(.*) to (.*) = (\d*)`)
		match := reg.FindStringSubmatch(line)

		name1 := match[1]
		name2 := match[2]
		dist, _ := strconv.Atoi(match[3])

		distByRoute[name1+name2] = dist
		distByRoute[name2+name1] = dist

		nodes[name1] = true
		nodes[name2] = true
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func goToNextNode(origin string, currDist int) {
	for n := range nodes {
		if !nodes[n] {
			continue
		}

		nodes[n] = false
		goToNextNode(n, currDist+distByRoute[origin+n])
		nodes[n] = true
	}

	for _, visited := range nodes {
		if visited {
			return
		}
	}

	// all nodes visited (i.e. map only contains "false")

	if currDist < minDist {
		minDist = currDist
	}

	if currDist > maxDist {
		maxDist = currDist
	}
}

func main() {
	distByRoute = make(map[string]int)
	nodes = make(map[string]bool)

	parseInput()

	for n := range nodes {
		nodes[n] = false // node visited -> deactivate
		goToNextNode(n, 0)
		nodes[n] = true // reactivate node for next try
	}

	fmt.Printf("Part1: %v\n", minDist)
	fmt.Printf("Part2: %v\n", maxDist)
}
