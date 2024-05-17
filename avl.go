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

func find[K cmp.Ordered, V any](n *node[K, V], key K) (*node[K, V], error) {
	for n != nil {
		if n.key == key {
			return n, nil
		} else if n.key < key {
			n = n.right
		} else {
			n = n.left
		}
	}
	return nil, ErrNil
}

func put[K cmp.Ordered, V any](n *node[K, V], key K, value V) *node[K, V] {
	if n == nil {
		return &node[K, V]{key: key, value: value}
	}

	if key < n.key {
		n.left = put(n.left, key, value)
	} else if key > n.key {
		n.right = put(n.right, key, value)
	} else if key == n.key {
		n.value = value
		return n
	}

	n.calcHeight()

	balanceFactor := balance(n)

	if balanceFactor > 1 {
		// LL
		if key < n.left.key {
			return ll(n)
		}
		// LR
		if key > n.left.key {
			return lr(n)
		}
	}

	if balanceFactor < -1 {
		// RL
		if key < n.right.key {
			return rl(n)
		}
		// RR
		if key > n.right.key {
			return rr(n)
		}
	}

	return n
}

func deletion[K cmp.Ordered, V any](n *node[K, V], key K) *node[K, V] {
	if n == nil {
		return n
	}
	if key < n.key {
		n.left = deletion(n.left, key)
	} else if key > n.key {
		n.right = deletion(n.right, key)
	} else {
		// found delete node.
		// get the smallest value from the right.
		// if none, get the largest value from the left.
		// if neither exists, exit.
		if s := successor(n.right); s != nil {
			n.key = s.key
			n.value = s.value
			n.right = deletion(n.right, s.key)
		} else if p := predecessor(n.left); p != nil {
			n.key = p.key
			n.value = p.value
			n.left = deletion(n.left, p.key)
		} else {
			n.free()
			return nil
		}
	}

	// calc height
	n.calcHeight()
	// rebalancing
	balanceFactor := balance(n)
	if balanceFactor < -1 {
		// right heavy
		rightBalanceFactor := balance(n.right)
		if rightBalanceFactor > 0 {
			// rl
			return rl(n)
		} else {
			// rr
			return rr(n)
		}
	} else if balanceFactor > 1 {
		// left heavy
		leftBalanceFactor := balance(n.left)
		if leftBalanceFactor < 0 {
			// lr
			return lr(n)
		} else {
			// ll
			return ll(n)
		}
	}
	return n
}

func successor[K cmp.Ordered, V any](n *node[K, V]) *node[K, V] {
	if n == nil {
		return nil
	}

	for n.left != nil {
		n = n.left
	}
	return n
}

func predecessor[K cmp.Ordered, V any](n *node[K, V]) *node[K, V] {
	if n == nil {
		return nil
	}
	for n.right != nil {
		n = n.right
	}
	return n
}
