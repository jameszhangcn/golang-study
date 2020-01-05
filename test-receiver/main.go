// test-receiver project main.go
package main

import (
	"fmt"
)

type Bag struct {
	items []int
}

func Insert(b *Bag, itemid int) {
	b.items = append(b.items, itemid)
}

func (b *Bag) Insert2(itemid int) {
	b.items = append(b.items, itemid)
}

type Point struct {
	X int
	Y int
}

type MyInt int

func (m MyInt) IsZero() bool {
	return m == 0
}
func (m MyInt) Add(other int) int {
	return other + int(m)
}

func (p Point) Add(other Point) Point {
	return Point{p.X + other.X, p.Y + other.Y}
}
func main() {
	fmt.Println("Hello World!")
	bag := new(Bag)
	Insert(bag, 1001)
	bag.Insert2(200)
	for _, v := range bag.items {
		println(v)
	}

	p1 := Point{1, 1}
	p2 := Point{2, 2}

	result := p1.Add(p2)
	fmt.Println(result)

	var b MyInt
	fmt.Println(b.IsZero())
	b = 1
	fmt.Println(b.Add(2))
}
