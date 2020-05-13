package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type reindeer struct {
	name      string
	speed     int
	speedTime int
	pauseTime int

	// only for part2
	currDist        int
	resting         bool // true when flying, false when resting
	timeUntilSwitch int
	points          int
}

var reindeers []reindeer

func main() {

	const time = 2503

	parseInput()

	part1(time)
	part2(time)
}

func parseInput() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, " ")

		speed, _ := strconv.Atoi(values[3])
		speedTime, _ := strconv.Atoi(values[6])
		pauseTime, _ := strconv.Atoi(values[13])

		r := reindeer{values[0], speed, speedTime, pauseTime, 0, false, speedTime, 0}
		reindeers = append(reindeers, r)
	}
}

func part1(time int) {
	maxDist := 0

	for _, r := range reindeers {
		fullRounds := time / (r.speedTime + r.pauseTime)
		remainder := time % (r.speedTime + r.pauseTime)

		temp := r.speed * r.speedTime
		if remainder < r.speedTime {
			temp = r.speed * remainder
		}

		distance := fullRounds*r.speed*r.speedTime + temp

		if distance > maxDist {
			maxDist = distance
		}
	}

	fmt.Printf("Part1: %v\n", maxDist)
}

func part2(time int) {
	for t := 0; t < time; t++ {
		maxDist := 0
		for i := range reindeers {
			r := &reindeers[i] // range operator creates a copy, so to modify the elements we need this

			if !r.resting {
				r.currDist += r.speed
			}

			if r.currDist > maxDist {
				maxDist = r.currDist
			}

			r.timeUntilSwitch--

			if r.timeUntilSwitch == 0 {
				r.resting = !r.resting

				if r.resting {
					r.timeUntilSwitch = r.pauseTime
				} else {
					r.timeUntilSwitch = r.speedTime
				}
			}
		}

		// give points to reindeers with highest dist in this second
		for i := range reindeers {
			r := &reindeers[i]
			if r.currDist == maxDist {
				r.points++
			}
		}
	}

	maxPoints := 0

	for _, r := range reindeers {
		if r.points > maxPoints {
			maxPoints = r.points
		}
	}

	fmt.Printf("Part2: %v\n", maxPoints)
}
