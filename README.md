# sortedlist
A sortedlist library written in Golang .

## Get started

### Install
```
go get github.com/hslam/sortedlist
```
### Import
```
import "github.com/hslam/sortedlist"
```
### Usage
#### Example
```
package main
import (
	"github.com/hslam/sortedlist"
	"fmt"
)
func main()  {
	s:=sortedlist.New(sortedlist.LessInt)
	fmt.Printf("ASC\t\t\t%t\n",s.ASC())
	fmt.Printf("DESC\t\t%t\n",s.DESC())
	s.Insert(1,1)
	s.Insert(1,1)
	s.Insert(3,3)
	s.Insert(3,3)
	s.Insert(2,2)
	s.Insert(2,2)
	s.Insert(5,5)
	s.Insert(6,6)
	s.Insert(4,4)
	s.Insert(4,4)
	printList(s)
	fmt.Printf("Contain\tScore:2 Value:2\t%t\n",s.Contain(2,2))
	fmt.Printf("Remove\tScore:2 Value:2\t%t\n",s.Remove(2,2))
	fmt.Printf("ContainScore\tScore:3\t%t\n",s.ContainScore(3))
	fmt.Printf("RemoveScore\t\tScore:3\t%t\n",s.RemoveScore(3))
	fmt.Printf("ContainValue\tValue:4\t%t\n",s.ContainValue(4))
	fmt.Printf("RemoveValue\t\tValue:4\t%t\n",s.RemoveValue(4))
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
Traverse	Score:1	Value:1
Traverse	Score:1	Value:1
Traverse	Score:2	Value:2
Traverse	Score:2	Value:2
Traverse	Score:3	Value:3
Traverse	Score:3	Value:3
Traverse	Score:4	Value:4
Traverse	Score:4	Value:4
Traverse	Score:5	Value:5
Traverse	Score:6	Value:6
=============End==============
Contain	Score:2 Value:2	true
Remove	Score:2 Value:2	true
ContainScore	Score:3	true
RemoveScore		Score:3	true
ContainValue	Value:4	true
RemoveValue		Value:4	true
===========Traverse===========
Traverse	Score:1	Value:1
Traverse	Score:1	Value:1
Traverse	Score:5	Value:5
Traverse	Score:6	Value:6
=============End==============
Top			Score:1	Value:1
Bottom		Score:6	Value:6
===========Traverse===========
Traverse	Score:1	Value:1
Traverse	Score:5	Value:5
=============End==============
```

### Licence
This package is licenced under a MIT licence (Copyright (c) 2019 Meng Huang)


### Authors
sortedlist was written by Meng Huang.


