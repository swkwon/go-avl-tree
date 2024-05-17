package avl

import "cmp"

type node[K cmp.Ordered, V any] struct {
	left   *node[K, V]
	right  *node[K, V]
	height int
	key    K
	value  V
}

func (n *node[K, V]) calcHeight() {
	n.height = max(height[K, V](n.left), height[K, V](n.right)) + 1
}

func (n *node[K, V]) free() {
	var k K
	var v V
	n.key = k
	n.value = v
	n.height = 0
	n.left = nil
	n.right = nil
}
