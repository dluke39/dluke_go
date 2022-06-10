package main

import (
	"fmt"
	"ninjalevel/dog"
	"ninjalevel/other"
)

type canine struct {
	name string
	age  int
}

func main() {
	r := canine{
		name: "Rock",
		age:  dog.Years(10),
	}

	g := other.Cat{
		Name:    "Goldie",
		Age:     2,
		Message: "Meow",
	}

	g.Battlecry()

	fmt.Println(r)
	fmt.Println("This is the updated version of the project.")
	fmt.Println("Do you think I should add a bunny rabbit next?")
}
