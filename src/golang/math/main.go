package main

import (
	"fmt"
	"math"
)

func main() {
	out := make(chan float64)
	for n := 1.0; n <= 10; n++ {
		go square(n, out)
		go cube(n, out)
		go squareRoot(n, out)
		fmt.Println(<-out)
		fmt.Println(<-out)
		fmt.Println(<-out)
	}
}

func squareRoot(num float64, out chan<- float64) {
	result := math.Sqrt(num)
	out <- result
}
func square(num float64, out chan<- float64) {
	result := num * num
	out <- result
}
func cube(num float64, out chan<- float64) {
	result := num * num * num
	out <- result
}
