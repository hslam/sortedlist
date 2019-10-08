# sortedlist
A sortedlist library written in Golang .

## Get started

### Install
```
go get hslam.com/mgit/Mort/sortedlist
```
### Import
```
import "hslam.com/mgit/Mort/sortedlist"
```
### Usage
#### Example
```
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
```

### Output
```
===========Traverse===========
Traverse		5
Traverse		4
Traverse		3
Traverse		3
Traverse		2
Traverse		2
Traverse		1
Traverse		1
=============End==============
Contain			3	true
ContainScore	3	true
ContainValue	3	true
Remove			3	true
RemoveScore		2	true
RemoveValue		1	true
Dequeue			5
===========Traverse===========
Traverse		4
=============End==============
```

### Licence
This package is licenced under a MIT licence (Copyright (c) 2019 Mort Huang)


### Authors
sortedlist was written by Mort Huang.


