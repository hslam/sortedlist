// Copyright (c) 2019 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

// Package sortedlist implements sorted list.
package sortedlist

type Less func(i Score, j Score) bool

type Score interface{}

type Value interface{}

var LessInt = func(i Score, j Score) bool {
	if i.(int) < j.(int) {
		return true
	} else {
		return false
	}
}

var LessInt64 = func(i Score, j Score) bool {
	if i.(int64) < j.(int64) {
		return true
	} else {
		return false
	}
}

var LessUint64 = func(i Score, j Score) bool {
	if i.(uint64) < j.(uint64) {
		return true
	} else {
		return false
	}
}

var LessString = func(i Score, j Score) bool {
	if i.(string) < j.(string) {
		return true
	} else {
		return false
	}
}

type Node struct {
	score Score
	value Value
	prev  *Node
	next  *Node
}

func (n *Node) Score() Score {
	return n.score
}

func (n *Node) Value() Value {
	return n.value
}

func (n *Node) Set(value Value) {
	n.value = value
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (n *Node) Next() *Node {
	return n.next
}

type SortedList struct {
	ascend bool
	less   Less
	head   *Node
	tail   *Node
	length int
}

func New(less Less) *SortedList {
	return newSortedList(less, true)
}

func NewASC(less Less) *SortedList {
	return newSortedList(less, true)
}

func NewDESC(less Less) *SortedList {
	return newSortedList(less, false)
}

func newSortedList(less Less, ascend bool) *SortedList {
	head := &Node{
		value: nil,
		prev:  nil,
	}
	tail := &Node{
		value: nil,
		prev:  nil,
	}
	head.next = tail
	tail.prev = head
	return &SortedList{
		ascend: ascend,
		less:   less,
		head:   head,
		tail:   tail,
	}
}

func (l *SortedList) Ascend() bool {
	return l.ascend
}

func (l *SortedList) Length() int {
	return l.length
}

func (l *SortedList) Front() *Node {
	if l.length == 0 {
		return nil
	}
	return l.head.next
}

func (l *SortedList) Head() *Node {
	if l.length == 0 {
		return nil
	}
	return l.head
}

func (l *SortedList) Rear() *Node {
	if l.length == 0 {
		return nil
	}
	return l.tail.prev
}

func (l *SortedList) Tail() *Node {
	if l.length == 0 {
		return nil
	}
	return l.tail
}

func (l *SortedList) Insert(score Score, value Value) bool {
	if value == nil {
		return false
	}
	node := &Node{
		score: score,
		value: value,
	}
	if l.length == 0 {
		node.next = l.head.next
		node.prev = l.head
		l.head.next.prev = node
		l.head.next = node
		l.length = 1
		return true
	}
	cur := l.head.next
	if l.ascend {
		for l.less(cur.score, score) {
			cur = cur.next
			if cur == l.tail {
				break
			}
		}
	} else {
		for l.less(score, cur.score) && cur != l.tail {
			cur = cur.next
		}
	}
	prev := cur.prev
	node.next = cur
	node.prev = prev
	prev.next = node
	cur.prev = node
	l.length++
	return true
}

func (l *SortedList) Top() *Node {
	if l.length == 0 {
		return nil
	}
	result := l.head.next
	l.head.next = result.next
	result.next.prev = l.head
	result.next = nil
	result.prev = nil
	l.length--
	return result
}

func (l *SortedList) Bottom() *Node {
	if l.length == 0 {
		return nil
	}
	result := l.tail.prev
	l.tail.prev = result.prev
	result.prev.next = l.tail
	result.next = nil
	result.prev = nil
	l.length--
	return result
}

func (l *SortedList) Contain(score Score, value Value) bool {
	if l.length == 0 {
		return false
	}
	cur := l.head.next
	for cur != l.tail {
		if cur.value == value && cur.score == score {
			return true
		}
		cur = cur.next
	}
	return false
}

func (l *SortedList) ContainValue(value Value) bool {
	if l.length == 0 {
		return false
	}
	cur := l.head.next
	for cur != l.tail {
		if cur.value == value {
			return true
		}
		cur = cur.next
	}
	return false
}

func (l *SortedList) ContainScore(score Score) bool {
	if l.length == 0 {
		return false
	}
	cur := l.head.next
	for cur != l.tail {
		if cur.score == score {
			return true
		}
		cur = cur.next
	}
	return false
}

func (l *SortedList) Remove(score Score, value Value) bool {
	if l.length == 0 {
		return false
	}
	cur := l.head.next
	for cur != l.tail {
		if cur.value == value && cur.score == score {
			cur.next.prev = cur.prev
			cur.prev.next = cur.next
			l.length--
			if cur.next != l.tail {
				if cur.next.value == value && cur.next.score == score {
					cur = cur.next
					continue
				}
			}
			return true
		}
		cur = cur.next
	}
	return false
}

func (l *SortedList) RemoveScore(score Score) bool {
	if l.length == 0 {
		return false
	}
	cur := l.head.next
	for cur != l.tail {
		if cur.score == score {
			cur.next.prev = cur.prev
			cur.prev.next = cur.next
			l.length--
			if cur.next != l.tail {
				if cur.next.score == score {
					cur = cur.next
					continue
				}
			}
			return true
		}
		cur = cur.next
	}
	return false
}

func (l *SortedList) RemoveValue(value Value) bool {
	if l.length == 0 {
		return false
	}
	cur := l.head.next
	for cur != l.tail {
		if cur.value == value {
			cur.next.prev = cur.prev
			cur.prev.next = cur.next
			l.length--
			if cur.next != l.tail {
				if cur.next.value == value {
					cur = cur.next
					continue
				}
			}
			return true
		}
		cur = cur.next
	}
	return false
}
