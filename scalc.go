package main

import (
	"fmt"
	"math"
)

const bufsize = 10

type llnode struct {
	car float64
	cdr *llnode
}

type llist *llnode

func (l llist) push(x float64) {
	nlist := new(llnode)
	*nlist.car = x
	*nlist.cdr = l
	l = nlist
}

func (l llist) pop() (float64, error) {
	if l == nil {
		return 0, errors.New("List Underrun: Popping empty list")
	}
	x := l.car
	l = l.cdr
	return x, nil
}

type Stack struct {
	head  [bufsize]float64
	tail  llist
	front int
}

type Dictionary map[string]func()

func (s Stack) push(x float64) {
	if s.front == bufsize {
		ntail := s.tail
		for i := 0; i < bufsize; i++ {
			ntail = ntail.push(s.head[i])
		}
		s.tail = ntail
		s.front = 0
	}
	s.head[s.front] = x
	s.front++
}

func (s Stack) pop() (float64, error) {
	if s.front == 0 {
		//figure out how to cache the top of list
	}
	return 0, nil //dummy implementation
}
