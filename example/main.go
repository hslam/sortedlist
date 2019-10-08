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
	s.Insert(1,1)
	s.Insert(3,3)
	s.Insert(3,3)
	s.Insert(2,2)
	s.Insert(2,2)
	s.Insert(5,5)
	s.Insert(5,5)
	s.Insert(6,6)
	s.Insert(4,4)
	printList(s)
	fmt.Printf("Contain\t\t\t3\t%t\n",s.Contain(3,3))
	fmt.Printf("ContainScore\t3\t%t\n",s.ContainScore(3))
	fmt.Printf("ContainValue\t3\t%t\n",s.ContainValue(3))
	fmt.Printf("Remove\t\t\t3\t%t\n",s.Remove(3,3))
	fmt.Printf("RemoveScore\t\t2\t%t\n",s.RemoveScore(2))
	fmt.Printf("RemoveValue\t\t1\t%t\n",s.RemoveValue(1))
	fmt.Printf("Dequeue\t\t\t%d\n",s.Dequeue().Value())
	printList(s)
}
func printList(s *sortedlist.SortedList){
	fmt.Println("===========Traverse===========")
	cur:=s.Front()
	for cur!=nil&&cur!=s.Tail(){
		fmt.Printf("Traverse\t\t%d\n",cur.Value())
		cur=cur.Next()
	}
	fmt.Println("=============End==============")
}