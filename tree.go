package avl

import (
	"cmp"
	"errors"
	"sync"
)

var ErrNil = errors.New("nil")

type Tree[K cmp.Ordered, V any] struct {
	mutex *sync.Mutex
	root  *node[K, V]
}

type GetResult[K cmp.Ordered, V any] struct {
	Key     K
	IsExist bool
	Value   V
}

func (t *Tree[K, V]) Delete(keys ...K) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	for _, k := range keys {
		n, err := t.find(t.root, k)
		if err == nil {
			n.isDeleted = true
			var empty V
			n.value = empty
		}
	}
}

func (t *Tree[K, V]) Gets(keys ...K) []GetResult[K, V] {
	var result []GetResult[K, V]

	if len(keys) <= 0 {
		return result
	}
	t.mutex.Lock()
	defer t.mutex.Unlock()

	for _, k := range keys {
		n, err := t.find(t.root, k)
		exist := false
		if err == nil {
			exist = true
		}
		result = append(result, GetResult[K, V]{Key: k, IsExist: exist, Value: n.value})
	}
	return result
}

func (t *Tree[K, V]) Put(key K, value V) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.root = t.put(t.root, key, value)
}

func (t *Tree[K, V]) find(n *node[K, V], key K) (*node[K, V], error) {
	if n == nil {
		return nil, ErrNil
	}

	if n.key == key {
		if n.isDeleted {
			return nil, ErrNil
		} else {
			return n, nil
		}
	}

	if n.key < key {
		t.find(n.right, key)
	} else if n.key > key {
		t.find(n.left, key)
	}

	return nil, ErrNil
}

func (t *Tree[K, V]) put(n *node[K, V], key K, value V) *node[K, V] {
	if n == nil {
		return &node[K, V]{key: key, value: value}
	}

	if key < n.key {
		n.left = t.put(n.left, key, value)
	} else if key > n.key {
		n.right = t.put(n.right, key, value)
	} else if key == n.key {
		n.value = value
		n.isDeleted = false
		return n
	}

	n.calcHeight()

	balanceValue := balance(n)

	if balanceValue > 1 {
		// LL
		if key < n.left.key {
			return ll(n)
		}
		// LR
		if key > n.left.key {
			return lr(n)
		}
	}

	if balanceValue < -1 {
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
