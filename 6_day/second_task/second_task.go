package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func reverseList(head *Node) *Node {
	curr := head.Next
	head.Next = nil
	last := head
	for {
		next := curr.Next
		curr.Next = last
		last = curr
		curr = next

		if next == nil {
			break
		}
	}
	return last

}

func main() {

}
