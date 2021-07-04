package cxmath

import (
	"errors"
)

//not concurrent
type StackI []int

func (q *StackI) Push(n int) {
	*q = append(*q, n)
}

func (q *StackI) Pop() (int, error) {
	x := q.Len() - 1
	if x <= 0 {
		return 0, errors.New("stack is empty")
	}
	n := (*q)[x]
	*q = (*q)[:x]
	return n, nil
}
func (q *StackI) Len() int {
	return len(*q)
}
