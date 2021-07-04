package cxmath

import (
	"errors"
)

//not concurrent
type QueueI []int

func (q *QueueI) Push(n int) {
	*q = append(*q, n)
}

func (q *QueueI) Pop() (int, error) {
	if q.Len() == 0 {
		return 0, errors.New("queue is empty")
	}

	n := (*q)[0]
	*q = (*q)[1:]
	return n, nil
}

func (q *QueueI) Len() int {
	return len(*q)
}
