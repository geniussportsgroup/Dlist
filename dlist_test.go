package Dlist

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
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
	}
	
	{
		l := New(1, 2, 3, 4, 5)
		assert.False(t, l.IsEmpty())
		assert.False(t, l.IsUnitarian())
		assert.False(t, l.IsUnitarianOrEmpty())
		
		for it := NewIterator(l); it.HasCurr(); it.Next() {
			fmt.Println(it.GetCurr().(int))
		}
	}
}
