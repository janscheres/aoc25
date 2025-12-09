package main

import (
	"fmt"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x >= 0{
		return x
	} else {
		return -x
	}
}

func main() {
	sum :=0
	lines := strings.Split(strings.TrimSpace(input), "\n")
	coords := make([][2]int, len(lines))
	for k, l := range lines {
		s := strings.Split(l, ",")
		fmt.Println(s)
		x, _ := strconv.Atoi(s[0])
		y, _ := strconv.Atoi(s[1])

		for i:=0; i<k; i++{
			compare := coords[i]
			area := (abs(compare[0]-x)+1)*(abs(compare[1]-y)+1)
			if area > sum {
				fmt.Println(x, y, compare, area)
				sum = area
			}
		}
		coords[k]=[2]int{x, y}
	}
	fmt.Println("Part 1:", sum)
}

const input = ``
