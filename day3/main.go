package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	sum := 0

	for l := range strings.SplitSeq(strings.TrimSpace(input), "\n") {
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

	sum = 0
	for l := range strings.SplitSeq(strings.TrimSpace(input), "\n") {
		largestArr := make([]int, 12)
		largestIndexArr := make([]int, 12)
		nums := strings.Split(l, "")
		for k := range largestArr {
			allowedFrom := 0
			if k > 0 {
				allowedFrom=largestIndexArr[k-1]+1	
			}

			maxAllowed := len(nums)-(12-k)
			for idx, i := range nums[allowedFrom:] {
				if allowedFrom+idx > maxAllowed {
					break
				}
				num, _ := strconv.Atoi(i)
				if num > largestArr[k] {
					largestArr[k]=num
					largestIndexArr[k]=idx+allowedFrom
				}
			}

		}

		for k := range largestArr {
			sum += int(math.Pow(10, float64(12-k-1)))*largestArr[k]
		}
	}
	fmt.Println("Part 2:", sum)
}

const input = `987654321111111
811111111111119
234234234234278
818181911112111`
