package main

import (
	"hslam.com/mgit/Mort/sortedlist"
	"fmt"
)

func main()  {
	less:= func(i sortedlist.Score,j sortedlist.Score)bool{
		if i.(int)<j.(int){return true} else {return false}
	}
	s:=sortedlist.NewSortedList(less)
	s.Insert(1,1)
	s.Insert(3,3)
	s.Insert(5,5)
	s.Insert(3,3)
	s.Insert(2,2)
	s.Insert(4,4)
	printList(s)
	for	s.Length()>0{
		fmt.Printf("Dequeue %d\n",s.Dequeue().Value())
	}
	printList(s)
}

func printList(s *sortedlist.SortedList){
	cur:=s.Front()
	for cur!=nil&&cur!=s.Tail(){
		fmt.Printf("Traverse %d\n",cur.Value())
		cur=cur.Next()
	}
}