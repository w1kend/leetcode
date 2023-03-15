package queue

type Queue struct {
	i     int
	items []int
}

func NewQueue() Queue {
	return Queue{
		items: make([]int, 10),
	}
}

func (q *Queue) Push(item int) {
	if q.i > len(q.items)/2 {
		q.items = q.items[q.i:]
		q.i = 0
	}
	q.items = append(q.items, item)
}

func (q *Queue) Pop() int {
	q.i++
	return q.items[q.i-1]
}

func (q *Queue) IsEmpty() bool {
	return q.i == len(q.items)
}

func (q *Queue) Peek() int {
	return q.items[q.i]
}
