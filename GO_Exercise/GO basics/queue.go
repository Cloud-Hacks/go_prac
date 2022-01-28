package main

import "fmt"

type queue []int

func (q *queue) enqueue(element int) []int {
	*q = append(*q, element)
	fmt.Println("Enqueued:", element)
	return *q
}

func dequeue(queue []int) []int {
	element := queue[0]
	fmt.Println("Dequeued:", element)
	return queue[1:]
}

func main() {
	var q queue

	q = q.enqueue(1)
	q = q.enqueue(5)
	q = q.enqueue(3)
	q = q.enqueue(11)

	fmt.Println("Queue:", q)

	q = dequeue(q)
	fmt.Println("Queue:", q)

	q = q.enqueue(2)
	fmt.Println("Queue:", q)
}
