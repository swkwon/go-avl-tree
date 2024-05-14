package avl

import "cmp"

type node[K cmp.Ordered, V any] struct {
	left      *node[K, V]
	right     *node[K, V]
	height    int
	key       K
	value     V
	isDeleted bool
}

func (n *node[K, V]) calcHeight() {
	n.height = max(height[K, V](n.left), height[K, V](n.right)) + 1
}
