package main

import "fmt"

type BST struct {
	key   int
	left  *BST
	right *BST
}

func (b *BST) Insert(key int) {

	if b.key > key {
		//GO LEFT
		if b.left == nil {
			b.left = &BST{key: key}
		} else {
			b.left.Insert(key)

		}
	} else if b.key < key {
		//GO RIGHT
		if b.right == nil {
			b.right = &BST{key: key}
		} else {
			b.right.Insert(key)
		}
	}

}

func (b *BST) Search(key int) bool {

	if b == nil {
		return false
	}
	if b.key > key {
		//Search LEFT
		b.left.Search(key)
	} else if b.key < key {
		//Search RIGHT
		b.right.Search(key)
	}
	return true
}
func (t *BST) Inorder() {
	if t == nil {
		return
	}
	t.left.Inorder()
	fmt.Println("", t.key)
	t.right.Inorder()
}
func (t *BST) PreOrder() {
	if t == nil {
		return
	}
	fmt.Println("", t.key)
	t.left.PreOrder()
	t.right.PreOrder()

}

func (t *BST) PostOrder() {
	if t == nil {
		return
	}
	t.left.PostOrder()
	t.right.PostOrder()
	fmt.Println("", t.key)
}
func main() {
	var x interface{} = "abc"
	str := fmt.Sprintf("%v", x)
	b := &BST{key: 100}
	b.Insert(50)
	b.Insert(500)
	b.Insert(80)
	b.Insert(900)
	b.Insert(970)
	b.Insert(20)
	// b.Inorder()
	// b.PreOrder()
	b.PostOrder()
	fmt.Println("BST", str)
	list := []int{10, 54, 8, 78, 7, 8, 8, 7, 6, 4, 5, 454, 6, 46, 4, 5, 46478, 865, 98, 7}
	for inn, number := range list {

		fmt.Println("BsssssssssssST", inn)
		if number%2 == 1 {
			fmt.Printf("ODD")
			fmt.Println("", number)

		}
	}
}
