package sortedlist

type Less func(i Score,j Score)bool
type Score interface{}
type Value interface{}

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
	less		Less
	front		*Node
	rear		*Node
	length		int
}
func NewSortedList(less Less) (*SortedList) {
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
		less:		less,
		front:		front,
		rear:		rear,
	}
}

func (l *SortedList) Length() int {
	return l.length
}

func (l *SortedList) Front() *Node {
	if l.length == 0 {
		return nil
	}
	return l.front.next
}

func (l *SortedList) Head() *Node {
	if l.length == 0 {
		return nil
	}
	return l.front
}

func (l *SortedList) Rear() *Node {
	if l.length == 0 {
		return nil
	}
	return l.rear.prev
}
func (l *SortedList) Tail() *Node {
	if l.length == 0 {
		return nil
	}
	return l.rear
}
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
	for l.less(score,cur.score)&& cur != l.rear {
		cur = cur.next
	}
	prev:=cur.prev
	node.next=cur
	node.prev=prev
	prev.next=node
	cur.prev=node
	l.length++
	return true
}

func (l *SortedList) Dequeue() *Node{
	if l.length == 0 {
		return nil
	}
	result := l.front.next
	l.front.next = result.next
	result.next = nil
	result.prev = nil
	l.length--
	return result
}

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