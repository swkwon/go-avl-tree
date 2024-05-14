package avl

import (
	"cmp"
	"fmt"
	"math/rand"
	"testing"
)

func TestAVL(t *testing.T) {
	myTree := New[string, string]()
	var numbers []int
	for i := 0; i < 200; i++ {
		numbers = append(numbers, rand.Intn(100))
	}

	fmt.Printf("numbers: %v\n", numbers)
	for _, v := range numbers {
		value := fmt.Sprintf("%d", v)
		myTree.Put(value, value)

	}
	printTree(myTree)
}

func printTree[K cmp.Ordered, V any](t *Tree[K, V]) {
	fmt.Println("start")
	if t.root != nil {
		print(t.root, 0)
	}
	fmt.Print("end\n\n")
}

func print[K cmp.Ordered, V any](n *node[K, V], depth int) {
	if n.right != nil {
		print[K, V](n.right, depth+1)
	}

	for i := 0; i < depth; i++ {
		fmt.Print("        ")
	}
	fmt.Printf("-> %v\n", n.key)

	if n.left != nil {
		print[K, V](n.left, depth+1)
	}
}
