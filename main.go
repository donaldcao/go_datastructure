package main

import (
	"go-datastructure/Linear"
)

func main() {


	var list Linear.LinkedList
	list.AddToHeader(1)
	list.AddToHeader(3)
	list.AddToEnd(5)
	list.AddAfter(1, 7)
	list.IterateList()
}
