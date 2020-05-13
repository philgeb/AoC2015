package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

func walkJSON(data interface{}) int {
	sum := 0

	switch typedData := data.(type) {
	case float64:
		return int(typedData)
	case []interface{}:
		for _, d := range typedData {
			sum += walkJSON(d)
		}
	case map[string]interface{}:
		for _, v := range typedData {

			if vStr, ok := v.(string); ok {
				if vStr == "red" {
					return 0
				}
			}

			sum += walkJSON(v)
		}
	default:
		return 0
	}

	return sum
}

func main() {
	file, _ := os.Open("input.txt")
	b, _ := ioutil.ReadAll(file)
	reg := regexp.MustCompile(`(-?\d+)`)
	match := reg.FindAll(b, -1)

	// Part 1
	sum := 0
	for _, m := range match {
		num, _ := strconv.Atoi(string(m))
		sum += num
	}
	fmt.Printf("Part1: %v\n", sum)

	// Part 2
	var initData interface{}
	json.Unmarshal(b, &initData)

	sum = walkJSON(initData)

	fmt.Printf("Part2: %v\n", sum)
}
