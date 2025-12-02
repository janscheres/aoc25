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
	fmt.Println("part 1", sum)

	sum = 0
	for _, r := range strings.Split(input, ",") {
		dash := strings.Index(r, "-")
		before, _ := strconv.Atoi(r[:dash])
		after, _ := strconv.Atoi(r[dash+1:])

		for i := before; i<=after; i++ {
			str := strconv.Itoa(i)
			l := len(str)

			for j := range l/2 +1 {// this +1 is the bane of my existance
				count := strings.Count(str, str[:j])
				//fmt.Println(j, str[:j], str, count, l)
				if count*j==l {
					sum += i
					break
				}
			}
		}
	}
	fmt.Println("part 2", sum)
}


const input = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`
