package stack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClear(t *testing.T) {
	s := NewStack(32, nil)
	s.Clear()
	require.Equal(t, 0, s.height)
}

func TestDupicate(t *testing.T) {
	expected := NewStack(32, nil)
	expected.height = 2
	s := NewStack(32, nil)
	s.Duplicate()
	require.Equal(t, expected.height, s.height)
	s.Push(12)
	s.Duplicate()
	expected.Push(12)
	expected.Push(12)
	require.Equal(t, expected.height, s.height)
	require.Equal(t, expected.data, s.data)
}

func TestPop(t *testing.T) {
	s := NewStack(32, nil)
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
	s := NewStack(32, nil)
	for i := 0; i < 32; i++ {
		s.Push(i)
	}
	require.Equal(t, 32, s.size)
	s.Push(-1)
	require.Equal(t, 64, s.size)
}

func TestSwap(t *testing.T) {
	s := NewStack(32, nil)
	s2 := NewStack(32, nil)
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

func TestHeights(t *testing.T) {
	// TODO
}

func TestTransfert(t *testing.T) {
	// TODO
}
