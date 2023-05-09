package main

import (
	"fmt"
)

func main() {
	tickets := make([]string, 1000000)
	for i := 0; i < 1000000; i++ {
		tickets[i] = fmt.Sprintf("%06d", i)
	}

	count := 0
	for _, v := range tickets {
		if isLucky(v) {
			count++
		}
	}

	fmt.Println(count)
}

func isLucky(number string) bool {
	start, end := 0, 0
	for _, num := range number[:3] {
		start += int(num - '0')
	}

	for _, num := range number[len(number)-3:] {
		end += int(num - '0')
	}

	return start == end
}
