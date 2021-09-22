package stack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClear(t *testing.T) {
	s := NewStack()
	s.Clear()
	require.Equal(t, 0, s.height)
}

func TestDupicate(t *testing.T) {
	s := NewStack()
	s2 := NewStack()
	s.Duplicate()
	require.Equal(t, s2.height, s.height)
	s.Push(12)
	s.Duplicate()
	s2.Push(12)
	s2.Push(12)
	require.Equal(t, s2.height, s.height)
	require.Equal(t, s2.data, s.data)
}

func TestPop(t *testing.T) {
	s := NewStack()
	v := s.Pop()
	require.Equal(t, 0, v)
	s.Push(12)
	s.Push(14)
	v = s.Pop()
	require.Equal(t, 14, v)
	v = s.Pop()
	require.Equal(t, 12, v)
	v = s.Pop()
	require.Equal(t, 0, v)
}

func TestPush(t *testing.T) {
	s := NewStack()
	for i := 0; i < 32; i++ {
		s.Push(i)
	}
	require.Equal(t, 32, s.size)
	s.Push(-1)
	require.Equal(t, 64, s.size)
}

func TestSwap(t *testing.T) {
	s := NewStack()
	s2 := NewStack()
	s.Swap()
	s2.Push(0)
	s2.Push(0)
	require.Equal(t, s2, s)
	s.Clear()
	s.Push(1)
	s.Swap()
	s2.Clear()
	s2.Push(1)
	s2.Push(0)
	require.Equal(t, s2, s)
	s.Clear()
	s.Push(1)
	s.Push(2)
	s2.Clear()
	s2.Push(2)
	s2.Push(1)
	s.Swap()
	require.Equal(t, s2, s)
}
