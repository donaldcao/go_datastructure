package Linear

import "fmt"

type Node struct {
	Val int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (list *LinkedList) AddToHeader(val int) {
	newNode := Node{}
	newNode.Val = val
	tmp := list.Head
	list.Head = &newNode
	newNode.Next = tmp
}

func (list *LinkedList) IterateList() {
	var node *Node
	for node = list.Head; node != nil; node = node.Next {
		fmt.Println(node.Val)
	}
}

func (list *LinkedList) LastNode() *Node {
	var node *Node
	for node = list.Head; node != nil; node = node.Next {
		if node.Next == nil {
			return node
		}
	}
	return nil
}

func (list *LinkedList) AddToEnd(val int) {
	newNode := &Node{Val:val}
	var end *Node;
	for end = list.Head; end != nil; end = end.Next {
		if end.Next == nil {
			end.Next = newNode
			return
		}
	}
}

func (list *LinkedList) Find(val int) *Node {
	var cur *Node
	for cur = list.Head; cur != nil; cur = cur.Next {
		if cur.Val == val {
			return cur
		}
	}
	return nil
}

func(list* LinkedList) AddAfter(after int, val int) {
	afterNode := list.Find(after)
	if afterNode == nil {
		list.AddToEnd(val)
	} else {
		newNode := &Node{Val: val}
		tmp := afterNode.Next
		afterNode.Next = newNode
		newNode.Next = tmp
	}
}





