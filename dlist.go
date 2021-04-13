package Dlist

type Dlist struct {
	Item interface{}
	prev *Dlist
	next *Dlist
}

// Swap in O(1) two sequences. Return l
func (l *Dlist) Swap(rhs interface{}) interface{} {

	link := rhs.(*Dlist)

	if l.IsEmpty() && link.IsEmpty() {
		return l
	}

	if l.IsEmpty() {
		link.next.prev = l
		link.prev.next = l
		l.next = link.next
		l.prev = link.prev
		link.Reset()
		return l
	}

	if link.IsEmpty() {
		l.next.prev = link
		l.prev.next = link
		link.next = l.next
		link.prev = l.prev
		l.Reset()
		return l
	}

	l.prev.next, link.prev.next = link.prev.next, l.prev.next
	l.next.prev, link.next.prev = link.next.prev, l.next.prev
	l.prev, link.prev = link.prev, l.prev
	l.next, link.next = link.next, l.next

	return l
}

func (l *Dlist) Reset() *Dlist {
	l.next, l.prev = l, l // circular
	return l
}

func NewDnode(item interface{}) *Dlist {
	p := &Dlist{Item: item}
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

// First Return the first element of the list
func (l *Dlist) First() interface{} {
	if l == nil {
		return nil
	}
	return l.next.Item
}

// Last Return the last element of the list
func (l *Dlist) Last() interface{} {
	if l == nil {
		return nil
	}
	return l.prev.Item
}

func (l *Dlist) __append(node *Dlist) *Dlist {
	node.next = l
	node.prev = l.prev
	l.prev.next = node
	l.prev = node
	return l
}

func (l *Dlist) Append(item interface{}, items ...interface{}) interface{} {
	l.__append(NewDnode(item))
	for _, i := range items {
		l.__append(NewDnode(i))
	}
	return l
}

func (l *Dlist) __insert(p *Dlist) *Dlist {
	p.prev = l
	p.next = l.next
	l.next.prev = p
	l.next = p
	return l
}

func (l *Dlist) Insert(item interface{}) *Dlist {
	return l.__insert(NewDnode(item))
}

// Del autoremove the node p from the linked list. Return an interface to the remove node
func (l *Dlist) Del() interface{} {
	l.prev.next = l.next
	l.next.prev = l.prev
	l.Reset()
	return l
}

// RemoveFirst remove the first item of the list whose header node is h. Return an interface to the
// removed node
func (l *Dlist) RemoveFirst() interface{} {
	return l.next.Del()
}

// RemoveLast remove the first item of the list whose header node is h. Return an interface to the
// removed node
func (l *Dlist) RemoveLast() interface{} {
	return l.prev.Del()
}

func New(items ...interface{}) *Dlist {

	head := NewDnode(nil)
	for _, item := range items {
		head.Append(item)
	}
	return head
}

func (l *Dlist) Create(items ...interface{}) interface{} {
	return New(items...)
}

func (l *Dlist) Clone() *Dlist {
	ret := NewDnode(nil)
	for it := NewIterator(l); it.HasCurr(); it.Next() {
		ret.Append(it.GetCurr())
	}
	return ret
}

// InsertList return ll + l. ll is emptied
func (l *Dlist) InsertList(ll *Dlist) *Dlist {

	if ll.IsEmpty() {
		return l
	}

	ll.prev.next = l.next
	ll.next.prev = l
	l.next.prev = ll.prev
	l.next = ll.next
	ll.Reset()

	return l
}

func (l *Dlist) AppendList(head *Dlist) *Dlist {

	if head.IsEmpty() {
		return l
	}

	head.next.prev = l.prev
	head.prev.next = l
	l.prev.next = head.next
	l.prev = head.prev
	head.Reset()

	return l
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
	return it.curr.Item
}

// Next Advance the iterator to the next element of the list
func (it *Iterator) Next() interface{} {
	it.curr = it.curr.next
	if it.curr == it.listPtr {
		return nil
	}
	return it
}

// DelCurr delete delete current item and advance iterator to next in the sequence. Return the deleted item
func (it *Iterator) DelCurr() interface{} {
	ret := it.curr
	it.curr = it.curr.next
	return ret.Del()
}

func (l *Dlist) Size() int {
	n := 0
	for it := NewIterator(l); it.HasCurr(); it.Next() {
		n++
	}
	return n
}

// Traverse the list and execute operation. It stops if the operation returns false. Return true if
// all the elements of the list were traversed
func (l *Dlist) Traverse(operation func(interface{}) bool) bool {

	for it := NewIterator(l); it.HasCurr(); it.Next() {
		if !operation(it.GetCurr()) {
			return false
		}
	}

	return true
}

// ToSlice Return a slice with the elements of the list
func (l *Dlist) ToSlice() []interface{} {

	ret := make([]interface{}, 0, 4)
	for it := NewIterator(l); it.HasCurr(); it.Next() {
		ret = append(ret, it.GetCurr())
	}

	return ret
}

// ReverseInPlace Reverse the list in place
func (l *Dlist) ReverseInPlace() *Dlist {

	tmp := New()
	for !l.IsEmpty() {
		tmp.Insert(l.RemoveFirst())
	}

	return l.Swap(tmp).(*Dlist)
}

// Reverse Return a reversed copy of seq
func (l *Dlist) Reverse() *Dlist {
	return l.Clone().ReverseInPlace()
}

// RotateLeftInPlace Rotate in place n positions to left
func (l *Dlist) RotateLeftInPlace(n int) *Dlist {

	if l.IsEmpty() || n == 0 {
		return l
	}

	for i := 0; i < n; i++ {
		l.Append(l.RemoveFirst().(*Dlist).Item)
	}

	return l
}

// RotateLeft Return a copy of seq rotated n positions to left
func (l *Dlist) RotateLeft(n int) *Dlist {
	return l.Clone().RotateLeftInPlace(n)
}

// RotateRightInPlace Rotate in place n positions to right
func (l *Dlist) RotateRightInPlace(n int) *Dlist {

	if l.IsEmpty() || n == 0 {
		return l
	}

	for i := 0; i < n; i++ {
		l.Insert(l.RemoveLast().(*Dlist).Item)
	}

	return l
}

// RotateRight Return a copy of seq rotated n positions to right
func (l *Dlist) RotateRight(n int) *Dlist {
	return l.Clone().RotateRightInPlace(n)
}
