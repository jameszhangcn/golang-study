// test-quickSort project main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	var array = [10]int{6, 5, 4, 8, 2, 11, 7, 10, 45, 23}
	low := 0
	high := len(array) - 1
	var A = 23
	fmt.Println("low:", low, "high:", high, "Key:", A)
	fmt.Println("\n Before: ")
	for i := 0; i < len(array); i++ {
		fmt.Printf("%d ", array[i])
	}

	for low < high {
		for A < array[high] && (low < high) {
			high--
		}
		if A > array[high] {
			array[low], array[high] = array[high], array[low]
			fmt.Println("\n switch: ")
			fmt.Println("low:", low, "high:", high, "Key:", A)
			for i := 0; i < len(array); i++ {
				fmt.Printf("%d ", array[i])
			}
			low++
		}
		for A > array[low] && (low < high) {
			low++
		}
		if A < array[low] {
			array[low], array[high] = array[high], array[low]
			fmt.Println("\n switch: ")
			for i := 0; i < len(array); i++ {
				fmt.Printf("%d ", array[i])
			}
			high--
		}
	}
	fmt.Println("\n After: ")
	for i := 0; i < len(array); i++ {
		fmt.Printf("%d ", array[i])
	}
}
