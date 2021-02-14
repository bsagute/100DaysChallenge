package main

import "fmt"

type Stack struct {
	arr []int
}

func (s *Stack) Push(element int) {
	s.arr = append(s.arr, element)
}
func (s *Stack) Pop() int {
	removedItem := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return removedItem
}
func main() {
	s := Stack{}
	s.Push(10)
	s.Push(20)
	s.Push(30)
	s.Push(40)
	fmt.Println("SS", s)
	fmt.Println("REMOVED ITEM ", s.Pop())
	fmt.Println("SS", s)
	fmt.Println("REMOVED ITEM ", s.Pop())
	fmt.Println("SS", s)
	fmt.Println("REMOVED ITEM ", s.Pop())
	fmt.Println("SS", s)
}
