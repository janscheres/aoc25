package main

import (
	"fmt"
	"sort"
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

	fmt.Println("Part 1:", sum)

	freshSum := 0
	sort.Slice(freshArr, func(i, j int) bool {//sort by start
    	return freshArr[i][0] < freshArr[j][0]
	})
	var noOverlap [][2]int
	currentStart := freshArr[0][0]
	currentEnd := freshArr[0][1]

	for _, r := range freshArr[1:] {
		start := r[0]
		end := r[1]

		if start <= currentEnd+1 {
			if end > currentEnd {
				currentEnd = end
			}
		} else {
			noOverlap = append(noOverlap, [2]int{currentStart, currentEnd})
			currentStart=start
			currentEnd=end
		}
	}
	noOverlap=append(noOverlap, [2]int{currentStart,currentEnd})

	for _, r := range noOverlap {
		freshSum+= r[1]-r[0]+1
	}
	fmt.Println(freshSum)
}

const input = ``
