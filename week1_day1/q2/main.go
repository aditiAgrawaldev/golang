package main

import "fmt"

type Node struct {
	value string
	left  *Node
	right *Node
}

func Preorder(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(root.value)
	Preorder(root.left)
	Preorder(root.right)
}
