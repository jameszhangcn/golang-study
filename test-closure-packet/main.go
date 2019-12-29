// test-closure-packet project main.go
package main

import (
	"fmt"
)

func playerGen(name string) func() (string, int) {
	hp := 150
	return func() (string, int) {
		return name, hp
	}
}
func main() {
	fmt.Println("Hello World!")
	generator := playerGen("sand and wood")
	name, hp := generator()
	fmt.Println(name, hp)
	name2, hp2 := generator()
	fmt.Println(name2, hp2)
}
