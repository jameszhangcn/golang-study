// test-list project main.go
package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	test_list := list.New()
	test_list.PushBack("back")
	test_list.PushFront(65)
	fmt.Println("test-list \n")
	for i := test_list.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
