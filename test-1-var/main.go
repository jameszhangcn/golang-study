// test-1-var project main.go
package main

import (
	"fmt"
	"math"
)

var g_a float32 = 3.14

const Avogadro = 6.02214129e23
const Planck = 6.62606957e-34

func main() {
	//var {
	//	a int
	//	b int
	//	c string
	//	d []float32
	//	e func() bool
	//	f struct {
	//		x int
	//	}
	//}
	a := 3
	var b int = 100
	var c int = 200

	c, b = b, c
	fmt.Println("Hello World!")
	fmt.Printf("a= %d", a)
	fmt.Printf("b= %d c = %d", b, c)

	a1, _ := getData()
	_, b1 := getData()
	fmt.Println(a1, b1)

	fmt.Printf("\n Global a: %f", g_a)

	sum := sum(g_a, g_a)
	fmt.Printf("\n Sum: %f", sum)
}

func getData() (int, int) {
	return 100, 200
}

func sum(a float32, b float32) float32 {
	num := a + b
	return num
}
