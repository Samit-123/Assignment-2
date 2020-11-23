package main

import "fmt"

const capacity = 100

type Queue struct {
	size  int
	data  [capacity]int
	front int
	back  int
}

func (q *Queue) Add(value int) bool {
	if q.size >= capacity {
		return false
	}
	q.size++
	q.data[q.back] = value
	q.back = (q.back + 1) % (capacity - 1)
	return true
}
func (q *Queue) Remove() (int, bool) {
	var value int
	if q.size <= 0 {
		return 0, false
	}
	q.size--
	value = q.data[q.front]
	q.front = (q.front + 1) % (capacity - 1)
	return value, true
}
func (q *Queue) IsEmpty() bool {
	return q.size == 0
}
func (q *Queue) Size() int {
	return q.size
}
func main() {
	q := &Queue{size: 0, front: 0, back: 0}
	q.Add(10)
	q.Add(20)
	q.Add(30)
	q.Add(40)
	q.Add(50)
	res1, res2 := q.Remove()
	fmt.Print(res1, " ")
	fmt.Println(res2)
	fmt.Println(q.IsEmpty())
}
