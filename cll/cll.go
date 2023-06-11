package cll

import (
	"fmt"
)

type Node struct {
	Word string
	Next *Node
}

func Init() map[byte]*CLLNode {
	cllColection := make(map[byte]*CLLNode)
	return cllColection
}

// CLLNode represents a circular linked list
type CLLNode struct {
	Head *Node
	Tail *Node
}

// Insert inserts a new node in the circular linked list while maintaining ascending order of string size
func (cll *CLLNode) Insert(data string) *CLLNode {
	newNode := &Node{Word: data}

	if cll.Head == nil {
		cll.Head = newNode
		cll.Tail = newNode
		newNode.Next = newNode
	} else if len(data) <= len(cll.Head.Word) {
		newNode.Next = cll.Head
		cll.Head = newNode
		cll.Tail.Next = newNode
	} else if len(data) >= len(cll.Tail.Word) {
		newNode.Next = cll.Head
		cll.Tail.Next = newNode
		cll.Tail = newNode
	} else {
		current := cll.Head
		for current.Next != cll.Head {
			if len(data) >= len(current.Word) && len(data) < len(current.Next.Word) {
				newNode.Next = current.Next
				current.Next = newNode
				break
			}
			current = current.Next
		}
	}
	return cll
}

// Display displays the elements of the circular linked list
func (cll *CLLNode) Display() {
	if cll.Head == nil {
		fmt.Println("Circular linked list is empty.")
		return
	}

	current := cll.Head
	for {
		fmt.Printf("%s ", current.Word)
		current = current.Next
		if current == cll.Head {
			break
		}
	}
	fmt.Println()
}

// Search searches for a node with the given data in the circular linked list
func (cll *CLLNode) Search(data string) bool {
	if cll.Head == nil {
		return false
	}

	current := cll.Head
	for {
		if current.Word == data {
			return true
		}
		current = current.Next
		if current == cll.Head {
			break
		}
	}

	return false
}

// Delete deletes the node with the given data from the circular linked list
func (cll *CLLNode) Delete(data string) bool {
	if cll.Head == nil {
		return false
	}

	// If the node to be deleted is the head
	if cll.Head.Word == data {
		// If there is only one node in the list
		if cll.Head == cll.Tail {
			cll.Head = nil
			cll.Tail = nil
		} else {
			cll.Head = cll.Head.Next
			cll.Tail.Next = cll.Head
		}
		return true
	}

	// Search for the node to be deleted
	prev := cll.Head
	current := cll.Head.Next
	for current != cll.Head {
		if current.Word == data {
			prev.Next = current.Next
			if current == cll.Tail {
				cll.Tail = prev
			}
			return true
		}
		prev = current
		current = current.Next
	}

	return false
}
