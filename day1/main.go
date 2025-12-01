package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	current := 50
	numzeros := 0

	for l := range strings.Lines(input) {
		l = strings.TrimSpace(l)
		if l == "" {
			continue
		}
		i, _ := strconv.Atoi(l[1:])
		//fmt.Println(i)
		if l[0]=='R' {
			current = (current+i)%100
		} else {
			current = (current-i)%100
		}

		if current == 0 {
			numzeros++
		}
	}
	fmt.Println("Part 1:", numzeros)

	current = 50
	numzeros = 0

	for l := range strings.Lines(input) {
		l = strings.TrimSpace(l)
		if l == "" {
			continue
		}

		i, _ := strconv.Atoi(l[1:])

		for a:=0; a<i; a++ {
			if l[0] == 'R' {
				current++

				if current > 99 {
					current = 0
				}
			} else {
				current --

				if current < 0 {
					current = 99
				}
			}

			if current==0 {
				numzeros++
			}
		}
	}
	fmt.Println("Part 2:", numzeros)
}

/*const input = `L1
R1
R2`*/
