package main

import (
	"fmt"
	"ninjalevel/dog"
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

	fmt.Println(r)
	fmt.Println("This is the updated version of the project.")
}
