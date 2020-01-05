// test-struct project main.go
package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

var p Point

type Player struct {
	Name        string
	HealthPoint int
	MagicPoint  int
}

type Command struct {
	Name    string
	Var     *int
	Comment string
}

func newCommand(name string, varref *int, comment string) *Command {
	return &Command{
		Name:    name,
		Var:     varref,
		Comment: comment,
	}
}

func main() {
	fmt.Println("Test Struct!")
	p.X = 10
	p.Y = 50
	fmt.Println("Px:", p.X, "Py:", p.Y)
	var version int = 1
	tank := new(Player)
	tank.Name = "test1"
	tank.HealthPoint = 300

	tank2 := &Player{}
	tank2.Name = "test2"
	tank2.HealthPoint = 500
	fmt.Println("tank2 name", tank2.Name, "tank2 health", tank2.HealthPoint)
	cmd := newCommand("version", &version, "show version")
	fmt.Println(cmd.Comment)
}
