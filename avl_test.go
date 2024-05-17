package avl

import (
	"cmp"
	"fmt"
	"math/rand"
	"testing"
)

func TestAVL(t *testing.T) {
	myTree := New[int, string]()
	var numbers []int
	for i := 0; i < 20; i++ {
		numbers = append(numbers, i)
	}
	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})

	fmt.Println(printKeys(numbers))
	for _, v := range numbers {
		value := fmt.Sprintf("%d", v)
		myTree.Put(v, value)

	}
	printTree(myTree)

	resultList := myTree.Gets(1, 100)
	for _, v := range resultList {
		fmt.Printf("%#v\n", v)
	}
	fmt.Println("======================================================")

	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})

	for _, v := range numbers {
		fmt.Printf("delete: %d\n", v)
		myTree.Delete(v)
		printTree(myTree)
	}
}

const (
	Center int = 0
	Left   int = 1
	Right      = 2
)

func printKeys(arr []int) string {
	str := ""
	str += "["
	for i, v := range arr {
		str += fmt.Sprintf("%d", v)
		if i != len(arr)-1 {
			str += ", "
		}
	}
	str += "]"
	return str
}

func printTree[K cmp.Ordered, V any](t *Tree[K, V]) {
	fmt.Println("======================================================")
	if t.root != nil {
		print(t.root, 0, Center)
	}
	fmt.Printf("======================================================\n")
}

func print[K cmp.Ordered, V any](n *node[K, V], depth int, pos int) {
	if n.right != nil {
		print[K, V](n.right, depth+1, Right)
	}

	for i := 0; i < depth; i++ {
		if i == depth-1 {
			if pos == Right {
				fmt.Print("     ┌──")
			} else if pos == Left {
				fmt.Print("     └──")
			}
		} else {
			fmt.Print("        ")
		}
	}
	fmt.Printf("─ %v(%d)\n", n.key, n.height)

	if n.left != nil {
		print[K, V](n.left, depth+1, Left)
	}
}
