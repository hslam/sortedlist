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
	s:=sortedlist.New(less)
	fmt.Printf("ASC\t\t\t%t\n",s.ASC())
	fmt.Printf("DESC\t\t%t\n",s.DESC())
	s.Insert(10,1)
	s.Insert(15,1)
	s.Insert(30,3)
	s.Insert(35,3)
	s.Insert(20,2)
	s.Insert(25,2)
	s.Insert(50,5)
	s.Insert(60,6)
	s.Insert(40,4)
	s.Insert(45,4)
	printList(s)
	fmt.Printf("Contain\tScore:20 Value:2\t%t\n",s.Contain(20,2))
	fmt.Printf("Remove\tScore:20 Value:2\t%t\n",s.Remove(20,2))
	fmt.Printf("ContainScore\tScore:30\t%t\n",s.ContainScore(30))
	fmt.Printf("RemoveScore\t\tScore:30\t%t\n",s.RemoveScore(30))
	fmt.Printf("ContainValue\tValue:4\t\t%t\n",s.ContainValue(4))
	fmt.Printf("RemoveValue\t\tValue:4\t\t%t\n",s.RemoveValue(4))
	printList(s)
	top:=s.Top()
	fmt.Printf("Top\t\t\tScore:%d\tValue:%d\n",top.Score(),top.Value())
	bottom:=s.Bottom()
	fmt.Printf("Bottom\t\tScore:%d\tValue:%d\n",bottom.Score(),bottom.Value())
	printList(s)
}
func printList(s *sortedlist.SortedList){
	fmt.Println("===========Traverse===========")
	cur:=s.Front()
	for cur!=nil&&cur!=s.Tail(){
		fmt.Printf("Traverse\tScore:%d\tValue:%d\n",cur.Score(),cur.Value())
		cur=cur.Next()
	}
	fmt.Println("=============End==============")
}
```

### Output
```
ASC			true
DESC		false
===========Traverse===========
Traverse	Score:10	Value:1
Traverse	Score:15	Value:1
Traverse	Score:20	Value:2
Traverse	Score:25	Value:2
Traverse	Score:30	Value:3
Traverse	Score:35	Value:3
Traverse	Score:40	Value:4
Traverse	Score:45	Value:4
Traverse	Score:50	Value:5
Traverse	Score:60	Value:6
=============End==============
Contain	Score:20 Value:2	true
Remove	Score:20 Value:2	true
ContainScore	Score:30	true
RemoveScore		Score:30	true
ContainValue	Value:4		true
RemoveValue		Value:4		true
===========Traverse===========
Traverse	Score:10	Value:1
Traverse	Score:15	Value:1
Traverse	Score:25	Value:2
Traverse	Score:35	Value:3
Traverse	Score:50	Value:5
Traverse	Score:60	Value:6
=============End==============
Top			Score:10	Value:1
Bottom		Score:60	Value:6
===========Traverse===========
Traverse	Score:15	Value:1
Traverse	Score:25	Value:2
Traverse	Score:35	Value:3
Traverse	Score:50	Value:5
=============End==============
```

### Licence
This package is licenced under a MIT licence (Copyright (c) 2019 Mort Huang)


### Authors
sortedlist was written by Mort Huang.


