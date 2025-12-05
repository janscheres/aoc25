package main

import (
	"fmt"
	"strings"
)

func main() {
	sum :=0

	lines := strings.Split(strings.TrimSpace(input), "\n")
	arr := make([][]string, len(lines))
	var toCheck = [][2]int{}
	for i, l := range lines {
		arr[i] = make([]string, len(lines[i]))
		for j, c := range strings.Split(l, "") {
			arr[i][j] = c
			if c == "@" {
				toCheck=append(toCheck, [2]int{i, j})
			}
		}
	}

	dirs := [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	for _, pos := range toCheck {
		neighbours := 0

		for _, v := range dirs {
			x := pos[0]+v[0]
			y := pos[1]+v[1]

			if x >=0 && x < len(lines) &&y >= 0 && y<len(lines[0]) {
				if arr[x][y] == "@" {
					neighbours++
				}
			}
		}
		if neighbours < 4 {
			sum++
		}
	 }
	 fmt.Println(sum)
}
const input = ``
