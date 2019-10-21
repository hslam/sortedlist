package sortedlist

type Less func(i Score,j Score)bool
type Score interface{}
type Value interface{}

var LessInt= func(i Score,j Score)bool{
	if i.(int)<j.(int){return true} else {return false}
}

var LessInt64= func(i Score,j Score)bool{
	if i.(int64)<j.(int64){return true} else {return false}
}

var LessUint64= func(i Score,j Score)bool{
	if i.(uint64)<j.(uint64){return true} else {return false}
}

var LessString= func(i Score,j Score)bool{
	if i.(string)<j.(string){return true} else {return false}
}


type Node struct {
	score		Score
	value		Value
	prev		*Node
	next		*Node
}

func (n *Node) Score()  Score {
	return n.score
}

func (n *Node) Value()  Value {
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
	asc 		bool
	less		Less
	front		*Node
	rear		*Node
	length		int
}
func New(less Less) (*SortedList) {
	return newSortedList(less,true)
}
func NewASC(less Less) (*SortedList) {
	return newSortedList(less,true)
}
func NewDESC(less Less) (*SortedList) {
	return newSortedList(less,false)
}
func newSortedList(less Less,asc bool) (*SortedList) {
	front := &Node{
		value:	nil,
		prev: 	nil,
	}
	rear := &Node{
		value:	nil,
		prev:	nil,
	}
	front.next = rear
	rear.prev=front
	return &SortedList{
		asc:		asc,
		less:		less,
		front:		front,
		rear:		rear,
	}
}
func (l *SortedList) ASC() bool {
	return l.asc
}
func (l *SortedList) DESC() bool {
	return !l.asc
}
func (l *SortedList) Length() int {
	return l.length
}
//read front data node
func (l *SortedList) Front() *Node {
	if l.length == 0 {
		return nil
	}
	return l.front.next
}
//read head node(nil)
func (l *SortedList) Head() *Node {
	if l.length == 0 {
		return nil
	}
	return l.front
}
//read rear data node
func (l *SortedList) Rear() *Node {
	if l.length == 0 {
		return nil
	}
	return l.rear.prev
}
//read tail node(nil)
func (l *SortedList) Tail() *Node {
	if l.length == 0 {
		return nil
	}
	return l.rear
}
//Insert
func (l *SortedList) Insert(score Score,value Value) bool {
	if value == nil {
		return false
	}
	node := &Node{
		score:score,
		value: value,
	}
	if l.length == 0 {
		node.next=l.front.next
		node.prev=l.front
		l.front.next.prev = node
		l.front.next=node
		l.length=1
		return true
	}
	cur := l.front.next
	if l.asc{
		for l.less(cur.score,score){
			cur = cur.next
			if cur == l.rear{
				break
			}
		}
	}else {
		for l.less(score,cur.score)&& cur != l.rear {
			cur = cur.next
		}
	}
	prev:=cur.prev
	node.next=cur
	node.prev=prev
	prev.next=node
	cur.prev=node
	l.length++
	return true
}
//read and remove
func (l *SortedList) Top() *Node{
	if l.length == 0 {
		return nil
	}
	result := l.front.next
	l.front.next = result.next
	result.next.prev=l.front
	result.next = nil
	result.prev = nil
	l.length--
	return result
}
//read and remove
func (l *SortedList) Bottom() *Node{
	if l.length == 0 {
		return nil
	}
	result := l.rear.prev
	l.rear.prev = result.prev
	result.prev.next=l.rear
	result.next = nil
	result.prev = nil
	l.length--
	return result
}
//read
func (l *SortedList) Contain(score Score,value Value) bool {
	if l.length == 0 {
		return false
	}
	cur := l.front.next
	for cur != l.rear {
		if cur.value == value && cur.score == score{
			return true
		}
		cur = cur.next
	}
	return false
}
//read
func (l *SortedList) ContainValue(value Value) bool {
	if l.length == 0 {
		return false
	}
	cur := l.front.next
	for cur != l.rear {
		if cur.value == value {
			return true
		}
		cur = cur.next
	}
	return false
}
//read
func (l *SortedList) ContainScore(score Score) bool {
	if l.length == 0 {
		return false
	}
	cur := l.front.next
	for cur != l.rear {
		if cur.score == score {
			return true
		}
		cur = cur.next
	}
	return false
}
//remove
func (l *SortedList) Remove(score Score,value Value) bool {
	if l.length == 0 {
		return false
	}
	cur := l.front.next
	for cur != l.rear {
		if cur.value == value && cur.score == score{
			cur.next.prev=cur.prev
			cur.prev.next=cur.next
			l.length--
			if cur.next!=l.rear{
				if cur.next.value==value&&cur.next.score == score{
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
//remove
func (l *SortedList) RemoveScore(score Score) bool {
	if l.length == 0 {
		return false
	}
	cur := l.front.next
	for cur != l.rear {
		if cur.score == score{
			cur.next.prev=cur.prev
			cur.prev.next=cur.next
			l.length--
			if cur.next!=l.rear{
				if cur.next.score == score{
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
//remove
func (l *SortedList) RemoveValue(value Value) bool {
	if l.length == 0 {
		return false
	}
	cur := l.front.next
	for cur != l.rear {
		if cur.value == value{
			cur.next.prev=cur.prev
			cur.prev.next=cur.next
			l.length--
			if cur.next!=l.rear{
				if cur.next.value==value {
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
