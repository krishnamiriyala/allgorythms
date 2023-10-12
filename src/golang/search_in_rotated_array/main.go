package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func searchInRotatedArray(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := (low + high) / 2

		if arr[mid] == target {
			return mid
		}

		// Check if the left half is sorted
		if arr[low] <= arr[mid] {
			// Check if the target is within the range of the sorted left half
			if arr[low] <= target && target <= arr[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else { // Right half is sorted
			// Check if the target is within the range of the sorted right half
			if arr[mid] <= target && target <= arr[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	return -1 // Element not found
}

func main() {
	var arrStr, targetStr string

	// Check if array is provided as a command-line argument
	if len(os.Args) > 1 {
		arrStr = os.Args[1]
	} else {
		// If not provided, read from standard input
		fmt.Print("Enter the array (comma-separated): ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		arrStr = scanner.Text()
	}

	// Check if target is provided as a command-line argument
	if len(os.Args) > 2 {
		targetStr = os.Args[2]
	} else {
		// If not provided, read from standard input
		fmt.Print("Enter the target: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		targetStr = scanner.Text()
	}

	// Parse the array
	var arr []int
	for _, numStr := range strings.Split(arrStr, ",") {
		num, err := strconv.Atoi(strings.TrimSpace(numStr))
		if err != nil {
			fmt.Println("Error parsing array:", err)
			return
		}
		arr = append(arr, num)
	}

	// Parse the target
	target, err := strconv.Atoi(targetStr)
	if err != nil {
		fmt.Println("Error parsing target:", err)
		return
	}

	// Call the search function
	result := searchInRotatedArray(arr, target)
	fmt.Println(result)
}
