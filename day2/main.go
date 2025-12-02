package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	sum := 0
	for _, r := range strings.Split(input, ",") {
		dash := strings.Index(r, "-")
		before, _ := strconv.Atoi(r[:dash])
		after, _ := strconv.Atoi(r[dash+1:])

		for i := before; i<=after; i++ {
			str := strconv.Itoa(i)
			l := len(str)

			if l%2==0 {
				if str[:l/2] == str[l/2:] {
					sum +=i
				}
			}
		}
	}
	fmt.Println("hello", sum)
}

const input = `INSERT HERE`
