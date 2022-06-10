package other

import "fmt"

type Cat struct {
	Name    string
	Age     int
	Message string
}

func (c Cat) Battlecry() {
	fmt.Println(c.Name, "says", c.Message, "!")
}
