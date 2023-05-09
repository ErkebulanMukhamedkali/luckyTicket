package main

import (
	"fmt"
	"strconv"
)

func main() {
	tickets := make([]string, 1000000)
	for i := 0; i < 1000000; i++ {
		temp := strconv.Itoa(i)

		zeros := ""
		for i := 0; i < 6-len(temp); i++ {
			zeros += "0"
		}

		if len(temp) < 6 {
			temp = fmt.Sprintf("%s%s", zeros, temp)
		}
		tickets[i] = temp
	}

	fmt.Println(tickets[:100])

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
