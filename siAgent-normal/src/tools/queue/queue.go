package queue

import "fmt"

type Queue struct {
	Elems []interface{}
	Top    int
	Bottom int
	Maxlen int
	Full   bool
}

// top ---> bottom
func New(maxLen int) *Queue {
	q := &Queue{
		Elems: make([]interface{}, maxLen),
		Top: 0,
		Bottom: 0,
		Maxlen: maxLen,
	}

	fmt.Printf("Queue: %d\n", q.Maxlen)
	return q
}

func (q *Queue) EnQueue(e interface{}) {

	if q.Maxlen == q.Bottom - q.Top {
		q.Bottom = q.Top
		q.Full = true
	}
	q.Elems[q.Bottom] = e
	q.Bottom++

}

func (q *Queue) DeQueue() {

	if q.Bottom > q.Top {
		q.Top++
	}

}

func (q *Queue) Len() int {
	return len(q.Elems)
}

