package Linear

type DoubleNode struct {
	 Val int
	 Next *DoubleNode
	 Prev *DoubleNode
}

type DoubleLinkedList struct {
	Head *DoubleNode
}

func (list *DoubleLinkedList) NodeBetweenValues(first , second int) *DoubleNode {
	var firstNode, secondNode *DoubleNode
	for cur := list.Head; cur != nil; cur = cur.Next {
		if cur.Val == first {
			firstNode = cur
		} else if cur.Val == second {
			secondNode = cur
		} else if firstNode != nil && secondNode != nil {
			return cur
		}
	}
	return nil
}

func (list *DoubleLinkedList) AddToHead(val int) {
	newNode := &DoubleNode{Val: val}
	if list.Head == nil {
		list.Head = newNode
	} else {
		list.Head.Prev = newNode
		newNode.Next = list.Head
		list.Head = newNode
	}
}

func (list *DoubleLinkedList) AddAfter(after, val int) {
	var cur *DoubleNode
	for cur = list.Head; cur != nil; cur = cur.Next {
		if cur.Val == after {
			newNode := &DoubleNode{Val: val}
			if cur.Next == nil {
				cur.Next = newNode
				newNode.Prev = cur
			} else {
				tmp := cur.Next
				cur.Next = newNode
				newNode.Prev = cur
				newNode.Next = tmp
				tmp.Prev = newNode
			}
			return
		}
	}
	// list is empty
	if cur == nil {
		list.Head = &DoubleNode{Val: val}
	}
	// TODO: if there is no candidate 'after' in the list, how to handle it?
}

func (list *DoubleLinkedList) AddToEnd(val int) {
	newNode := &DoubleNode{Val: val}
	var cur *DoubleNode
	for cur = list.Head; cur != nil; cur = cur.Next {
		if cur.Next == nil {
			cur.Next = newNode
			newNode.Prev = cur
			return
		}
	}
	if cur == nil {
		list.Head = &DoubleNode{Val: val}
	}
}


