// test-sort1 project main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	//sort
	arr := [...]int{21, 32, 12, 33, 34, 34, 87, 24}
	fmt.Println("before the sort \n", arr)
	var n = len(arr)
	for i := 0; i < n; i++ {
		fmt.Println("the ", i+1, "time:")
		for j := i; j < n; j++ {
			if arr[i] > arr[j] {
				t := arr[i]
				arr[i] = arr[j]
				arr[j] = t

			}
			fmt.Println(arr)
		}
	}
	fmt.Println("Finally: \n", arr)
}
