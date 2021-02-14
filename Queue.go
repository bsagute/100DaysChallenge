package main

import "fmt"

type Queue struct {
	arrr []int
}

func (q *Queue) EnQueue(a int) {
	q.arrr = append(q.arrr, a)
}

func (q *Queue) DeQueue() (removedItem int) {

	if len(q.arrr) > 0 {

		removedItem = q.arrr[0]
		q.arrr = q.arrr[1:len(q.arrr)]
		return removedItem
	}
	return removedItem

}
func main() {
	q := Queue{}
	q.EnQueue(10)
	q.EnQueue(20)
	q.EnQueue(30)
	q.EnQueue(40)
	fmt.Println("QUEUE", q)
	fmt.Println("REMOVED ", q.DeQueue())
	fmt.Println("QUEUE", q)
	fmt.Println("REMOVED ", q.DeQueue())
	fmt.Println("QUEUE", q)

}
