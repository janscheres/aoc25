package main

import (
	"fmt"
	"strings"
)

func main() {
	sum :=0
	lines := strings.Split(strings.TrimSpace(input), "\n")

	width := len(lines[0])

	current := make([]bool, width)
	before := make([]bool, width)
	
	for _, l := range lines {
		chars := strings.Split(l, "")

		for k, c := range chars {
			switch c {
			case "S":
				current[k]=true
			case "^":
				if before[k] {
					sum++
					current[k-1]=true
					current[k]=false
					current[k+1]=true
				}
			}
		}
		before = current
		//fmt.Println(current)	
	}
	fmt.Println("Part 1:", sum)
}

const input = ``
