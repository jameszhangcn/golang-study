// test-func project main.go
package main

import (
	"fmt"
)

func typedTwoValues() (int, int) {
	return 1, 2
}
func main() {
	fmt.Println("Hello World!")
	a, b := typedTwoValues()
	fmt.Println(a, b)
}
