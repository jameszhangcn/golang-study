// test-syncMap project main.go
package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Hello World!")
	var scene sync.Map
	scene.Store("")
}
