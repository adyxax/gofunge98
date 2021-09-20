package stack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewStack(t *testing.T) {
	require.Equal(t, NewStack(), &Stack{
		size:   32,
		height: 0,
		data:   make([]int, 32),
		next:   nil,
	})
}

func TestClear(t *testing.T) {
	s := NewStack()
	s.Clear()
	require.Equal(t, s.height, 0)
}

func TestDupicate(t *testing.T) {
	s := NewStack()
	s2 := NewStack()
	s.Duplicate()
	require.Equal(t, s.height, s2.height)
	s.Push(12)
	s.Duplicate()
	s2.Push(12)
	s2.Push(12)
	require.Equal(t, s.height, s2.height)
	require.Equal(t, s.data, s2.data)
}

func TestPop(t *testing.T) {
	s := NewStack()
	v := s.Pop()
	require.Equal(t, v, 0)
	s.Push(12)
	s.Push(14)
	v = s.Pop()
	require.Equal(t, v, 14)
	v = s.Pop()
	require.Equal(t, v, 12)
	v = s.Pop()
	require.Equal(t, v, 0)
}

func TestPush(t *testing.T) {
	s := NewStack()
	for i := 0; i < 32; i++ {
		s.Push(i)
	}
	require.Equal(t, s.size, 32)
	s.Push(-1)
	require.Equal(t, s.size, 64)
}

func TestSwap(t *testing.T) {
	s := NewStack()
	s2 := NewStack()
	s.Swap()
	s2.Push(0)
	s2.Push(0)
	require.Equal(t, s, s2)
	s.Clear()
	s.Push(1)
	s.Swap()
	s2.Clear()
	s2.Push(1)
	s2.Push(0)
	require.Equal(t, s, s2)
	s.Clear()
	s.Push(1)
	s.Push(2)
	s2.Clear()
	s2.Push(2)
	s2.Push(1)
	s.Swap()
	require.Equal(t, s, s2)
}
