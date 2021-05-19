package main

import (
	"fmt"
	linear "go-datastructure/Linear"
)

func main() {
	var set *linear.Set = &linear.Set{}
	set.New()
	set.AddElement(1)
	set.AddElement(2)
	fmt.Println(set)
	fmt.Println(set.ContainsElement(1))

}
