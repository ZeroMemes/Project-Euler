package main

import (
	"math"
)

type Node struct {
	value  int64
	factor int64
	next   *Node
}

func main() {
	input := &Node{600851475143, 1, nil}
	largest := int64(0)

	// Calculate the whole tree
	current := input
	for ok := true; ok; ok = current != nil {
		current = resolveNode(current)
	}

	// Find the largest gamer
	current = input.next
	for true {
		largest = max(current.factor, largest)
		if current.next == nil {
			largest = max(current.value, largest)
			break
		}
		current = current.next
	}

	println(largest)
}

func resolveNode(node *Node) *Node {
	for i := int64(2); i < int64(math.Ceil(math.Sqrt(float64(node.value)))); i++ {
		if node.value%i == 0 {
			node.next = &Node{node.value / i, i, nil}
			return node.next
		}
	}
	return nil
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
