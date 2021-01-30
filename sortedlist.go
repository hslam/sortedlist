// Copyright (c) 2019 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

// Package sortedlist implements a sorted list.
package sortedlist

// Less represents a less function.
type Less func(i interface{}, j interface{}) bool

// LessInt implements the Less function for int.
var LessInt = func(i interface{}, j interface{}) bool {
	if i.(int) < j.(int) {
		return true
	}
	return false
}

// LessInt64 implements the Less function for int64.
var LessInt64 = func(i interface{}, j interface{}) bool {
	if i.(int64) < j.(int64) {
		return true
	}
	return false
}

// LessUint64 implements the Less function for uint64.
var LessUint64 = func(i interface{}, j interface{}) bool {
	if i.(uint64) < j.(uint64) {
		return true
	}
	return false
}

// LessString implements the Less function for string.
var LessString = func(i interface{}, j interface{}) bool {
	if i.(string) < j.(string) {
		return true
	}
	return false
}

// Node represents a node.
type Node struct {
	score interface{}
	value interface{}
	prev  *Node
	next  *Node
}

// Score returns the node score.
func (n *Node) Score() interface{} {
	return n.score
}

// Value returns the node value.
func (n *Node) Value() interface{} {
	return n.value
}

// Set sets the node value.
func (n *Node) Set(value interface{}) {
	n.value = value
}

// Prev returns the prev node.
func (n *Node) Prev() *Node {
	return n.prev
}

// Next returns the next node.
func (n *Node) Next() *Node {
	return n.next
}

// SortedList represents a sorted list.
type SortedList struct {
	ascend bool
	less   Less
	head   *Node
	rear   *Node
	length int
}

// New returns a new ascend sorted list with the given Less function.
func New(less Less) *SortedList {
	return NewSortedList(less, true)
}

// NewASC returns a new ascend sorted list with the given Less function.
func NewASC(less Less) *SortedList {
	return NewSortedList(less, true)
}

// NewDESC returns a new descend sorted list with the given Less function.
func NewDESC(less Less) *SortedList {
	return NewSortedList(less, false)
}

// NewSortedList returns a new sorted list.
func NewSortedList(less Less, ascend bool) *SortedList {
	head := &Node{
		value: nil,
		prev:  nil,
	}
	rear := &Node{
		value: nil,
		prev:  nil,
	}
	head.next = rear
	rear.prev = head
	return &SortedList{
		ascend: ascend,
		less:   less,
		head:   head,
		rear:   rear,
	}
}

// Ascend returns ascend.
func (l *SortedList) Ascend() bool {
	return l.ascend
}

// Length returns length of the sorted list.
func (l *SortedList) Length() int {
	return l.length
}

// Front returns the front node.
func (l *SortedList) Front() *Node {
	if l.length == 0 {
		return nil
	}
	return l.head.next
}

// Rear returns the rear node.
func (l *SortedList) Rear() *Node {
	if l.length == 0 {
		return nil
	}
	return l.rear
}

// Insert inserts a value with a score.
func (l *SortedList) Insert(score interface{}, value interface{}) bool {
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
	cur := l.rear.prev
	if l.ascend {
		for l.less(score, cur.score) {
			cur = cur.prev
			if cur == l.head {
				break
			}
		}
	} else {
		for l.less(cur.score, score) {
			cur = cur.prev
			if cur == l.head {
				break
			}
		}
	}
	next := cur.next
	node.next = next
	node.prev = cur
	next.prev = node
	cur.next = node
	l.length++
	return true
}

// Top returns the top node of the sorted list.
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

// Bottom returns the bottom node of the sorted list.
func (l *SortedList) Bottom() *Node {
	if l.length == 0 {
		return nil
	}
	result := l.rear.prev
	l.rear.prev = result.prev
	result.prev.next = l.rear
	result.next = nil
	result.prev = nil
	l.length--
	return result
}

// Contain return whether the sorted list contains the score value.
func (l *SortedList) Contain(score interface{}, value interface{}) bool {
	if l.length == 0 {
		return false
	}
	cur := l.head.next
	for cur != l.rear {
		if cur.value == value && cur.score == score {
			return true
		}
		cur = cur.next
	}
	return false
}

// ContainValue return whether the sorted list contains the value.
func (l *SortedList) ContainValue(value interface{}) bool {
	if l.length == 0 {
		return false
	}
	cur := l.head.next
	for cur != l.rear {
		if cur.value == value {
			return true
		}
		cur = cur.next
	}
	return false
}

// ContainScore return whether the sorted list contains the score.
func (l *SortedList) ContainScore(score interface{}) bool {
	if l.length == 0 {
		return false
	}
	cur := l.head.next
	for cur != l.rear {
		if cur.score == score {
			return true
		}
		cur = cur.next
	}
	return false
}

// Remove removes the value with the score.
func (l *SortedList) Remove(score interface{}, value interface{}) bool {
	if l.length == 0 {
		return false
	}
	cur := l.head.next
	for cur != l.rear {
		if cur.value == value && cur.score == score {
			cur.next.prev = cur.prev
			cur.prev.next = cur.next
			l.length--
			if cur.next != l.rear {
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

// RemoveScore removes the score.
func (l *SortedList) RemoveScore(score interface{}) bool {
	if l.length == 0 {
		return false
	}
	cur := l.head.next
	for cur != l.rear {
		if cur.score == score {
			cur.next.prev = cur.prev
			cur.prev.next = cur.next
			l.length--
			if cur.next != l.rear {
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

// RemoveValue removes the value.
func (l *SortedList) RemoveValue(value interface{}) bool {
	if l.length == 0 {
		return false
	}
	cur := l.head.next
	for cur != l.rear {
		if cur.value == value {
			cur.next.prev = cur.prev
			cur.prev.next = cur.next
			l.length--
			if cur.next != l.rear {
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
