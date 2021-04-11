package Dlist

type Dlist struct {
	item interface{}
	prev *Dlist
	next *Dlist
}

func (p *Dlist) Reset() *Dlist {
	p.next, p.prev = p, p // circular
	return p
}

func NewDnode(item interface{}) *Dlist {
	p := &Dlist{item: item}
	return p.Reset()
}

func (l *Dlist) IsEmpty() bool {
	return l == l.next && l == l.prev
}

func (l *Dlist) IsUnitarian() bool {
	return l != l.next && l.next == l.prev
}

func (l *Dlist) IsUnitarianOrEmpty() bool {
	return l.next == l.prev
}

func (p *Dlist) __append(node *Dlist) *Dlist {
	node.next = p
	node.prev = p.prev
	p.prev.next = node
	p.prev = node
	return p
}

func (p *Dlist) Append(item interface{}) *Dlist {
	return p.__append(NewDnode(item))
}

func New(items ...interface{}) *Dlist {
	
	head := NewDnode(nil)
	for item := range items {
		head.Append(item)
	}
	return head
}


type Iterator struct {
	listPtr *Dlist
	curr    *Dlist
}

// NewIterator Return an iterator to the list
func NewIterator(l *Dlist) *Iterator {
	it := new(Iterator)
	it.listPtr = l
	it.curr = l.next
	return it
}

// CreateIterator Return an iterator to the list
func (l *Dlist) CreateIterator() interface{} {
	return NewIterator(l)
}

// ResetFirst Reset the iterator to the first element of the list
func (it *Iterator) ResetFirst() interface{} {
	it.curr = it.listPtr.next
	return it
}

// HasCurr Return true if the iterator is positioned on a valid element
func (it *Iterator) HasCurr() bool {
	return it.curr != it.listPtr
}

// IsLast Return true if the current element of the list is the last of the list
func (it *Iterator) IsLast() bool {
	return it.curr == it.listPtr.prev
}

// GetCurr Return the current element of the list
func (it *Iterator) GetCurr() interface{} {
	if it.curr == nil {
		return nil
	}
	return it.curr.item
}

// Next Advance the iterator to the next element of the list
func (it *Iterator) Next() interface{} {
	it.curr = it.curr.next
	if it.curr == it.listPtr {
		return nil
	}
	return it
}
