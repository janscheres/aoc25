package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	rows :=	strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]string, len(rows))
	for i, row := range rows {
		grid[i]=strings.Fields(row)
	}

	sum := 0

	for col := 0; col<len(grid[0]); col++ {
		operation := grid[len(rows)-1][col]
		var current int
		if operation == "*" {
			current = 1
		} else {
			current = 0
		}

		for i:=0; i<len(rows)-1; i++ {
			intnum, _ := strconv.Atoi(grid[i][col])
			if operation == "*" {
				current *= intnum
			} else {
				current += intnum
			}
		}
		sum+=current
	}
	fmt.Println("Part 1:", sum)
}

const input = ``
