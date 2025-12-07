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

	//part 2: instead of a bool, consider a spot as being an integer, how many possible ways to get there
	// we will then sum along the bottom row
	sum = 0

	ago := make([]int, width)
	now := make([]int, width)

	for _, l := range lines {
		chars := strings.Split(l, "")

		for k, c := range chars {
			switch c {
			case "S":
				now[k]=1
			case "^":
				if ago[k]>=1 {
					now[k-1]+=now[k]
					now[k+1]+=now[k]
					now[k]=0
				}
			}
		}
		ago =now 
		fmt.Println(now)	
	}
	for _, i := range now {
		sum +=i
	}
	fmt.Println("Part 2:", sum)


}

const input = ``
