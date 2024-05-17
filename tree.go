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
		t.root = deletion(t.root, k)
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
		n, err := find(t.root, k)
		if err == nil {
			result = append(result, GetResult[K, V]{Key: k, IsExist: true, Value: n.value})
		} else {
			result = append(result, GetResult[K, V]{Key: k})
		}

	}
	return result
}

func (t *Tree[K, V]) Put(key K, value V) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.root = put(t.root, key, value)
}
