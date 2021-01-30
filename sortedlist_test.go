// Copyright (c) 2019 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package sortedlist

import (
	"testing"
)

func TestLess(t *testing.T) {
	{
		var a, b int = 0, 1
		if !LessInt(a, b) {
			t.Error()
		}
		if LessInt(b, a) {
			t.Error()
		}
	}
	{
		var a, b int64 = 0, 1
		if !LessInt64(a, b) {
			t.Error()
		}
		if LessInt64(b, a) {
			t.Error()
		}
	}
	{
		var a, b uint64 = 0, 1
		if !LessUint64(a, b) {
			t.Error()
		}
		if LessUint64(b, a) {
			t.Error()
		}
	}
	{
		var a, b string = "a", "b"
		if !LessString(a, b) {
			t.Error()
		}
		if LessString(b, a) {
			t.Error()
		}
	}
}

func TestNode(t *testing.T) {
	var score, value = 1, 1
	var prev, next = &Node{}, &Node{}
	node := &Node{score: score, prev: prev, next: next}
	if node.Score() != score {
		t.Error()
	}
	node.Set(value)
	if node.Value() != value {
		t.Error()
	}
	if node.Prev() != prev {
		t.Error()
	}
	if node.Next() != next {
		t.Error()
	}
}

func TestSortedList(t *testing.T) {
	s := New(LessInt)
	if !s.Ascend() {
		t.Error()
	}
	if s.Length() != 0 {
		t.Error()
	}
	if s.Front() != nil {
		t.Error()
	}
	if s.Rear() != nil {
		t.Error()
	}
	if s.Insert(0, nil) {
		t.Error()
	}
	if s.Top() != nil {
		t.Error()
	}
	if s.Bottom() != nil {
		t.Error()
	}
	if s.Contain(0, 0) {
		t.Error()
	}
	if s.Contain(0, 0) {
		t.Error()
	}
	if s.ContainValue(0) {
		t.Error()
	}
	if s.ContainScore(0) {
		t.Error()
	}
	if s.Remove(0, 0) {
		t.Error()
	}
	if s.RemoveValue(0) {
		t.Error()
	}
	if s.RemoveScore(0) {
		t.Error()
	}
}

func TestAscendSortedList(t *testing.T) {
	s := NewASC(LessInt)
	s.Insert(3, 2)
	s.Insert(1, 1)
	s.Insert(3, 3)
	s.Insert(5, 4)
	s.Insert(6, 4)
	s.Insert(2, 5)
	s.Insert(2, 5)
	s.Insert(2, 6)
	s.Insert(4, 7)
	cur := s.Front()
	for cur != nil && cur != s.Rear() {
		next := cur.Next()
		if next != nil && next != s.Rear() {
			if LessInt(next.Score(), cur.Score()) {
				t.Error(next.Score(), cur.Score())
			}
		}
		cur = next
	}
	{
		if !s.Contain(2, 5) {
			t.Error()
		}
		if !s.Remove(2, 5) {
			t.Error()
		}
		if s.Contain(2, 5) {
			t.Error()
		}
		if s.Remove(2, 5) {
			t.Error()
		}
	}
	{
		if !s.ContainScore(3) {
			t.Error()
		}
		if !s.RemoveScore(3) {
			t.Error()
		}
		if s.ContainScore(3) {
			t.Error()
		}
		if s.RemoveScore(3) {
			t.Error()
		}
	}
	{
		if !s.ContainValue(4) {
			t.Error()
		}
		if !s.RemoveValue(4) {
			t.Error()
		}
		if s.ContainValue(4) {
			t.Error()
		}
		if s.RemoveValue(4) {
			t.Error()
		}
	}
	top := s.Top()
	if top.Score() != 1 || top.Value() != 1 {
		t.Error()
	}
	bottom := s.Bottom()
	if bottom.Score() != 4 || bottom.Value() != 7 {
		t.Error()
	}
	if s.Length() != 1 {
		t.Error()
	}
	if !s.Remove(2, 6) {
		t.Error()
	}
	if s.Length() != 0 {
		t.Error()
	}
}

func TestDescendSortedList(t *testing.T) {
	s := NewDESC(LessInt)
	s.Insert(1, 1)
	s.Insert(3, 2)
	s.Insert(3, 3)
	s.Insert(5, 4)
	s.Insert(6, 4)
	s.Insert(2, 5)
	s.Insert(2, 5)
	s.Insert(2, 6)
	s.Insert(4, 7)
	cur := s.Front()
	for cur != nil && cur != s.Rear() {
		next := cur.Next()
		if next != nil && next != s.Rear() {
			if LessInt(cur.Score(), next.Score()) {
				t.Error(cur.Score(), next.Score())
			}
		}
		cur = next
	}
	{
		if !s.Contain(2, 5) {
			t.Error()
		}
		if !s.Remove(2, 5) {
			t.Error()
		}
		if s.Contain(2, 5) {
			t.Error()
		}
		if s.Remove(2, 5) {
			t.Error()
		}
	}
	{
		if !s.ContainScore(3) {
			t.Error()
		}
		if !s.RemoveScore(3) {
			t.Error()
		}
		if s.ContainScore(3) {
			t.Error()
		}
		if s.RemoveScore(3) {
			t.Error()
		}
	}
	{
		if !s.ContainValue(4) {
			t.Error()
		}
		if !s.RemoveValue(4) {
			t.Error()
		}
		if s.ContainValue(4) {
			t.Error()
		}
		if s.RemoveValue(4) {
			t.Error()
		}
	}
	bottom := s.Bottom()
	if bottom.Score() != 1 || bottom.Value() != 1 {
		t.Error()
	}
	top := s.Top()
	if top.Score() != 4 || top.Value() != 7 {
		t.Error()
	}
	if s.Length() != 1 {
		t.Error()
	}
	if !s.Remove(2, 6) {
		t.Error()
	}
	if s.Length() != 0 {
		t.Error()
	}
}
