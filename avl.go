package avl

import (
	"cmp"
	"sync"
)

func New[K cmp.Ordered, V any]() *Tree[K, V] {
	return &Tree[K, V]{
		mutex: &sync.Mutex{},
	}
}

func height[K cmp.Ordered, V any](n *node[K, V]) int {
	if n == nil {
		return -1
	}
	return n.height
}

func balance[K cmp.Ordered, V any](n *node[K, V]) int {
	if n == nil {
		return 0
	}
	return height[K, V](n.left) - height[K, V](n.right)
}

func rightRotate[K cmp.Ordered, V any](n *node[K, V]) *node[K, V] {
	// right rotate
	left := n.left
	n.left = left.right
	left.right = n

	n.calcHeight()
	left.calcHeight()

	return left
}

func leftRotate[K cmp.Ordered, V any](n *node[K, V]) *node[K, V] {
	// left rotate
	right := n.right
	n.right = right.left
	right.left = n

	n.calcHeight()
	right.calcHeight()

	return right
}

func ll[K cmp.Ordered, V any](n *node[K, V]) *node[K, V] {
	return rightRotate[K, V](n)
}
func rr[K cmp.Ordered, V any](n *node[K, V]) *node[K, V] {
	return leftRotate[K, V](n)
}
func lr[K cmp.Ordered, V any](n *node[K, V]) *node[K, V] {
	n.left = leftRotate[K, V](n.left)
	return rightRotate[K, V](n)
}
func rl[K cmp.Ordered, V any](n *node[K, V]) *node[K, V] {
	n.right = rightRotate[K, V](n.right)
	return leftRotate[K, V](n)
}
