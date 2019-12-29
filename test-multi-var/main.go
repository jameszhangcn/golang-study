// test-multi-var project main.go
package main

import (
	"fmt"
)

func myfunc(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func rawPrint(rawList ...interface{}) {
	for _, a := range rawList {
		fmt.Println(a)
	}
}

func print(slist ...interface{}) {
	rawPrint(slist...)
}
func main() {
	fmt.Println("Multi vars! \n")
	myfunc(1)
	myfunc(1, 2, 3, 4)
	myfunc(6, 5, 4)

	print(1, 2, 3)
	print(8, 9, 10)
	print(5, 6, "testjames")
	rawPrint(4, 5, "test2")
}
