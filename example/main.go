package main

import (
	"fmt"
	"github.com/hslam/sortedlist"
)

func main() {
	s := sortedlist.New(sortedlist.LessInt)
	fmt.Printf("ASC\t\t%t\n", s.Ascend())
	s.Insert(1, 1)
	s.Insert(1, 1)
	s.Insert(3, 3)
	s.Insert(3, 3)
	s.Insert(2, 2)
	s.Insert(2, 2)
	s.Insert(5, 5)
	s.Insert(6, 6)
	s.Insert(4, 4)
	s.Insert(4, 4)
	printList(s)
	fmt.Printf("Contain\t\tScore:2 Value:2\t%t\n", s.Contain(2, 2))
	fmt.Printf("Remove\t\tScore:2 Value:2\t%t\n", s.Remove(2, 2))
	fmt.Printf("ContainScore\tScore:3\t%t\n", s.ContainScore(3))
	fmt.Printf("RemoveScore\tScore:3\t%t\n", s.RemoveScore(3))
	fmt.Printf("ContainValue\tValue:4\t%t\n", s.ContainValue(4))
	fmt.Printf("RemoveValue\tValue:4\t%t\n", s.RemoveValue(4))
	printList(s)
	top := s.Top()
	fmt.Printf("Top\t\tScore:%d\tValue:%d\n", top.Score(), top.Value())
	bottom := s.Bottom()
	fmt.Printf("Bottom\t\tScore:%d\tValue:%d\n", bottom.Score(), bottom.Value())
	printList(s)
}

func printList(s *sortedlist.SortedList) {
	fmt.Println("\n===========Traverse===========")
	cur := s.Front()
	for cur != nil && cur != s.Rear() {
		fmt.Printf("Traverse\tScore:%d\tValue:%d\n", cur.Score(), cur.Value())
		cur = cur.Next()
	}
	fmt.Println("=============End==============\n")
}
