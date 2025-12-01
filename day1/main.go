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
	fmt.Println(numzeros)
}

/*const input = `L1
R1
R2`*/
