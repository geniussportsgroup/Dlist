package Dlist

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"

	Fct "github.com/geniussportsgroup/FunctionalLib"
)

func TestNew(t *testing.T) {

	{
		l := New()
		assert.True(t, l.IsEmpty())
		assert.True(t, l.IsUnitarianOrEmpty())
	}

	{
		l := New(1)
		assert.False(t, l.IsEmpty())
		assert.True(t, l.IsUnitarian())
		assert.True(t, l.IsUnitarianOrEmpty())
		assert.Equal(t, l.First().(int), 1)
	}

	{
		l := New(1, 2, 3, 4, 5)
		assert.False(t, l.IsEmpty())
		assert.False(t, l.IsUnitarian())
		assert.False(t, l.IsUnitarianOrEmpty())
		assert.Equal(t, l.First().(int), 1)
		assert.Equal(t, l.Last().(int), 5)
		for it := NewIterator(l); it.HasCurr(); it.Next() {
			fmt.Println(it.GetCurr().(int))
		}
	}
}

func TestDlist_Insert(t *testing.T) {

	l := New()
	assert.Nil(t, l.First())
	assert.Nil(t, l.Last())

	for i := 1; i <= 5; i++ {
		l.Insert(i)
	}

	assert.Equal(t, l.First().(int), 5)
	assert.Equal(t, l.Last().(int), 1)
	assert.Equal(t, l.Size(), 5)
}

func TestDlist_InsertList(t *testing.T) {

	l0 := New()
	l1 := New(1, 2, 3, 4, 5)

	l0.InsertList(l1)
	assert.True(t, l1.IsEmpty())
	assert.Equal(t, l0.Size(), 5)
	assert.False(t, l0.IsEmpty())
	assert.Equal(t, l0.First().(int), 1)
	assert.Equal(t, l0.Last().(int), 5)

	l1.Swap(l0)
	assert.Equal(t, l1.First().(int), 1)
	assert.Equal(t, l1.Last().(int), 5)

	l2 := New(6, 7, 8, 9, 10)
	l2.InsertList(l1)
	assert.True(t, l1.IsEmpty())
	assert.Equal(t, l1.Size(), 0)
	assert.Equal(t, l2.First().(int), 1)
	assert.Equal(t, l2.Last().(int), 10)

	assert.True(t, Fct.All(Fct.Zip(New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10), l2),
		func(pair interface{}) bool {
			return pair.(Fct.Pair).Item1.(int) == pair.(Fct.Pair).Item2.(int)
		}))
}

func TestDlist_Swap(t *testing.T) {

	l := New()
	l0 := New()

	l.Swap(l0)
	assert.True(t, l.IsEmpty())
	assert.True(t, l0.IsEmpty())

	l1 := New(1, 2, 3, 4, 5)

	l0.Swap(l1)

	assert.False(t, l0.IsEmpty())
	assert.Equal(t, l0.Size(), 5)
	assert.Equal(t, l1.Size(), 0)
	assert.True(t, l1.IsEmpty())

	l0.Swap(l1)
	assert.True(t, l0.IsEmpty())
	assert.Equal(t, l0.Size(), 0)
	assert.Equal(t, l1.Size(), 5)
	assert.False(t, l1.IsEmpty())

	l2 := New(6, 7, 8, 9, 10)
	l1.Swap(l2)

	assert.False(t, l1.IsEmpty())
	assert.False(t, l2.IsEmpty())
	assert.Equal(t, l1.Size(), 5)
	assert.Equal(t, l2.Size(), 5)
	assert.Equal(t, l1.First().(int), 6)
	assert.Equal(t, l1.Last().(int), 10)
	assert.Equal(t, l2.First().(int), 1)
	assert.Equal(t, l2.Last().(int), 5)

	l1.Swap(l2)
	assert.False(t, l1.IsEmpty())
	assert.False(t, l2.IsEmpty())
	assert.Equal(t, l1.Size(), 5)
	assert.Equal(t, l2.Size(), 5)
	assert.Equal(t, l1.First().(int), 1)
	assert.Equal(t, l1.Last().(int), 5)
	assert.Equal(t, l2.First().(int), 6)
	assert.Equal(t, l2.Last().(int), 10)
}

func TestDlist_Append(t *testing.T) {

	l1 := New()
	l1.Append(1, 2, 3, 4, 5)
	l2 := l1.Clone()
	assert.True(t, Fct.All(Fct.Zip(l1, l2), func(p interface{}) bool {
		return p.(Fct.Pair).Item1 == p.(Fct.Pair).Item2
	}))
}

func TestDlist_AppendList(t *testing.T) {

	l0 := New()
	l1 := New(1, 2, 3, 4, 5)
	l2 := New(6, 7, 8, 9, 10)

	l0.AppendList(l1)
	assert.Equal(t, l0.Size(), 5)
	assert.True(t, l1.IsEmpty())
	assert.True(t, Fct.All(Fct.Zip(New(1, 2, 3, 4, 5), l0), func(p interface{}) bool {
		return p.(Fct.Pair).Item1 == p.(Fct.Pair).Item2
	}))

	l0.AppendList(l2)
	assert.Equal(t, l0.Size(), 10)
	assert.True(t, l2.IsEmpty())
	assert.True(t, l2.Traverse(func(i interface{}) bool {
		return i.(int) <= 10
	}))
}

func TestDlist_RemoveFirst(t *testing.T) {

	l := New(1, 2, 3, 4, 5)
	for i := 1; !l.IsEmpty(); i++ {
		assert.Equal(t, i, l.RemoveFirst().(*Dlist).Item.(int))
	}
}

func TestDlist_RemoveLast(t *testing.T) {

	l := New(1, 2, 3, 4, 5)
	for i := 5; !l.IsEmpty(); i-- {
		assert.Equal(t, i, l.RemoveLast().(*Dlist).Item.(int))
	}
}

func TestDlist_Reverse(t *testing.T) {

	l := New(1, 2, 3, 4, 5)
	assert.True(t, Fct.All(Fct.Zip(l.Reverse(), l.ReverseInPlace()), func(p interface{}) bool {
		fmt.Printf("%d == %d\n", p.(Fct.Pair).Item1.(*Dlist).Item, p.(Fct.Pair).Item2.(*Dlist).Item)
		return p.(Fct.Pair).Item1.(*Dlist).Item == p.(Fct.Pair).Item2.(*Dlist).Item
	}))
}

func TestDlist_RotateLeft(t *testing.T) {

	l := New(1, 2, 3, 4, 5)
	Fct.ForEach(l.RotateLeft(3), func(i interface{}) {
		fmt.Printf("%d ", i)
	})
	fmt.Println()
}

func TestDlist_RotateRight(t *testing.T) {

	l := New(1, 2, 3, 4, 5)
	Fct.ForEach(l.RotateRight(3), func(i interface{}) {
		fmt.Printf("%d ", i)
	})
	fmt.Println()
}

func TestDlist_ToSlice(t *testing.T) {

	s := New(1, 2, 3, 4, 5).ToSlice()
	for i := 1; i <= 5; i++ {
		assert.Equal(t, i, s[i-1])
	}
}
