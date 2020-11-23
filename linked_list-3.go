package main

import "fmt"

type List struct {
	head  *Node
	count int
}

type Node struct {
	value int
	next  *Node
}

func (list *List) Size() int {
	return list.count
}

func (list *List) isEmpty() bool {
	return (list.count == 0)
}

func (list *List) addHead(value int) {
	list.head = &Node{value, list.head}
	list.count++
}

func (list *List) addTail(value int) {
	curr := list.head
	newNode := &Node{value, nil}
	if curr == nil {
		list.head = newNode
	}
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = newNode
	list.count++
}

func (list *List) Print() {
	temp := list.head
	for temp != nil {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
	fmt.Println("")
}

func (list *List) SortedInsert(value int) {
	newNode := &Node{value, nil}
	curr := list.head
	if curr == nil || curr.value > value {
		newNode.next = list.head
		list.head = newNode
		list.count++
		return
	}
	for curr.next != nil && curr.next.value < value {
		curr = curr.next
	}
	newNode.next = curr.next
	curr.next = newNode
	list.count++
}

func (list *List) IsPresent(data int) bool {
	temp := list.head
	for temp != nil {
		if temp.value == data {
			return true
		}
		temp = temp.next
	}
	return false
}

func (list *List) RemoveHead() int {
	if list.isEmpty() {
		fmt.Println("EmptyListError")
		return 0
	}
	value := list.head.value
	list.head = list.head.next
	list.count--
	return value
}

func (list *List) DeleteNode(delValue int) bool {
	temp := list.head
	if list.isEmpty() {
		fmt.Println("EmptylistError")
		return false
	}
	if delValue == list.head.value {
		list.head = list.head.next
		list.count--
		return true
	}
	for temp.next != nil {
		if temp.next.value == delValue {
			temp.next = temp.next.next
			list.count--
			return true
		}
		temp = temp.next
	}
	return false
}

func (list *List) FreeList() {
	list.head = nil
	list.count = 0
}

func (list *List) Reverse() {
	curr := list.head
	var prev, temp *Node
	for curr != nil {
		temp = curr.next
		curr.next = prev
		prev = curr
		curr = temp
	}
	list.head = prev
}

func (list *List) ReverseRecurse() {
	list.head = list.reverseRecurseUtil(nil, list.head)

}
func (list *List) reverseRecurseUtil(prev *Node, curr *Node) *Node {
	if curr != nil {
		list.reverseRecurseUtil(curr, curr.next)
		curr.next = prev
	} else {
		list.head = prev
	}
	return list.head
}

func (list *List) RemoveDuplicate() {
	curr := list.head
	for curr != nil {
		if curr.next != nil && curr.value == curr.next.value {
			curr.next = curr.next.next
		} else {
			curr = curr.next
		}
	}
}
func main() {
	lst := &List{}
	lst.addHead(1)
	lst.addHead(2)
	lst.addTail(4)
	lst.SortedInsert(3)
	fmt.Println("Value search in Linked list:")
	fmt.Println(lst.IsPresent(4))
	fmt.Println(lst.Size())
	fmt.Println("Before remove head")
	lst.Print()
	res1 := lst.RemoveHead()
	fmt.Println("After removing:")
	fmt.Println("Removing element is:", res1)
	lst.Print()
	fmt.Println("Delete particular value:")
	lst.DeleteNode(3)
	lst.Print()
	fmt.Println("==========================")
	fmt.Println("Before Reversing:")
	lst.Print()
	lst.ReverseRecurse()
	fmt.Println("After reversing:")
	lst.addTail(4)
	lst.addHead(1)
	lst.addTail(4)
	lst.addHead(1)
	fmt.Println("Before Remoing Duplicate:")
	lst.Print()
	fmt.Println("After Remoing Duplicate:")
	lst.RemoveDuplicate()
	lst.Print()
}
