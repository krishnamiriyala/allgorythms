package main

import (
	"fmt"
)

func countdown(count int) {
	for num := count; num >= 1; num-- {
		fmt.Println(num)
	}
}

func main() {
	countdown(10)
}
