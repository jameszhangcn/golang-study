// test-unamed-func project main.go
package main

import (
	"flag"
	"fmt"
)

var skillParam = flag.String("skill", "", "skill to perform")

func main() {
	fmt.Println("Hello World!")
	f := func(data int) {
		fmt.Println("hello", data)
	}
	f(100)

	visit([]int{1, 2, 3, 4}, func(v int) {
		fmt.Println(v)
	})

	flag.Parse()

	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angel fly")
		},
	}

	if f, ok := skill[*skillParam]; ok {
		f()
	} else {
		fmt.Println("skill not found")
	}
}

func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}
