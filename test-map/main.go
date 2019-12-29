// test-map project main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	var mapList map[string]int
	mapCreate := make(map[string]float32)
	mapList = map[string]int{"one": 1, "two": 2, "three": 3}
	mapCreate["key1"] = 4.5
	mapCreate["key2"] = 19.3
	mapCreate["key3"] = 3.4556

	mapAssigned := mapList
	mapAssigned["two"] = 4
	fmt.Printf("map list one is %d", mapList["one"])
	fmt.Printf("map create key %f", mapCreate["key3"])

	for k, v := range mapCreate {
		fmt.Printf("\n Key %s value %f ", k, v)
	}
	delete(mapCreate, "key2")
	for k, v := range mapCreate {
		fmt.Printf("\n After delete Key %s value %f ", k, v)
	}

	m := make(map[int]int)
	go func() {
		for {
			m[1] = 1
		}
	}()

	go func() {
		for {
			_ = m[1]
		}
	}()

	for {

	}
}
