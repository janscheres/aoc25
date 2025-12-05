package main

import (
	"fmt"
	"strings"
)

func main() {
	prevSum := -1
	sum := 0

	lines := strings.Split(strings.TrimSpace(input), "\n")
	arr := make([][]string, len(lines))
	var toCheck = [][2]int{}
	for i, l := range lines {
		arr[i] = make([]string, len(lines[i]))
		for j, c := range strings.Split(l, "") {
			arr[i][j] = c
		}
	}

	dirs := [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	for sum != prevSum {
		prevSum = sum
		for i := range arr {
			for j := range arr[i] {
				if arr[i][j] == "@" {
					toCheck = append(toCheck, [2]int{i,j})
				}	
			}
		}

		for _, pos := range toCheck {
			ns :=0
			for _, v := range dirs {
				x := pos[0]+v[0]
				y := pos[1]+v[1]
				
				if x >= 0 && x < len(lines) && y >= 0 && y < len(lines[0]) {
					if arr[x][y] == "@" {
						ns++
					}
				}
			}
			
			if ns < 4 {
				sum++
				arr[pos[0]][pos[1]] = "x"
			}
		}

		toCheck = [][2]int{}
	}

	fmt.Println(sum)
}

const input = ``
