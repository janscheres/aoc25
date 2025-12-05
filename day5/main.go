package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	sum := 0
	trimmed := strings.TrimSpace(input)

	parts := strings.Split(trimmed, "\n\n")

	freshRanges := strings.Split(parts[0], "\n")
	freshArr := make([][2]int, len(freshRanges))

	for k, l := range freshRanges {
		r := strings.Split(l, "-")
		before, _ := strconv.Atoi(r[0])
		after, _ := strconv.Atoi(r[1])

		freshArr[k] = [2]int{before, after}
	}

	for _, j := range strings.Split(parts[1], "\n") {
		intj, _ := strconv.Atoi(j)
		
		for _, r := range freshArr {
			if intj >= r[0] && intj<=r[1] {
				sum++
				break
			}
		}
	}

	fmt.Println(sum)
}

const input = ``
