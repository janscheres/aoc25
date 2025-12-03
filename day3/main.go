package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	sum := 0

	for _, l := range strings.Split(strings.TrimSpace(input), "\n") {
		//fmt.Println(l)
		largest := 0
		largestIndex := 0
		nums := strings.Split(l, "")
		for k, i := range nums {
			num, _ := strconv.Atoi(i)
			if num > largest && k != len(l)-1 {
				largest = num	
				largestIndex = k
			}	
		}
		//fmt.Println(largest)
		
		secondlargest := 0
		for _, j := range nums[largestIndex+1:] {
			num, _ := strconv.Atoi(j)
			if num > secondlargest {
				 secondlargest = num
			}
		}

		//fmt.Println(secondlargest)

		sum += 10*largest + secondlargest
	}

	fmt.Println("Part 1:", sum)
}


const input = `436
1925
25453`
