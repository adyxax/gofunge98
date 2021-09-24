package pointer

import (
	"os"
	"testing"

	"git.adyxax.org/adyxax/gofunge98/pkg/field"
	"github.com/stretchr/testify/require"
)

func TestBegin(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		expected := NewStackStack()
		expected.head = &Stack{data: make([]int, 0), next: expected.head}
		expected.head.next.Push(0)
		expected.head.next.Push(0)
		expected.height++
		p := NewPointer()
		ss := p.ss
		ss.Begin(p)
		require.Equal(t, expected, ss)
		x, y := p.GetStorageOffset()
		require.Equal(t, 1, x)
		require.Equal(t, 0, y)
		// Let's push another one
		expected.head = &Stack{data: make([]int, 0), next: expected.head}
		expected.head.next.Push(1)
		expected.head.next.Push(0)
		expected.height++
		ss.Begin(p)
		require.Equal(t, expected, ss)
		x, y = p.GetStorageOffset()
		require.Equal(t, 1, x)
		require.Equal(t, 0, y)
	})
	t.Run("negative", func(t *testing.T) {
		expected := NewStackStack()
		expected.head = &Stack{size: 5, height: 0, data: make([]int, 5), next: expected.head}
		expected.head.next.Push(0)
		expected.head.next.Push(0)
		expected.head.next.Push(0)
		expected.head.next.Push(0)
		expected.head.next.Push(0)
		expected.head.next.Push(0)
		expected.head.next.Push(0)
		expected.height++
		p := NewPointer()
		file, err := os.Open("../field/test_data/hello.b98")
		require.NoError(t, err, "Failed to open file")
		f, err := field.Load(file)
		p.Step(*f)
		ss := p.ss
		ss.head.Push(-5)
		ss.Begin(p)
		require.Equal(t, expected, ss)
		x, y := p.GetStorageOffset()
		require.Equal(t, 2, x)
		require.Equal(t, 0, y)
	})
	t.Run("ask to copy more than we have", func(t *testing.T) {
		expected := NewStackStack()
		expected.head = &Stack{size: 34, height: 34, data: make([]int, 34), next: expected.head}
		expected.head.data[33] = 18
		expected.head.next.Push(2)
		expected.head.next.Push(3)
		expected.height++
		p := NewPointer()
		p.SetStorageOffset(2, 3)
		file, err := os.Open("../field/test_data/hello.b98")
		require.NoError(t, err, "Failed to open file")
		f, err := field.Load(file)
		p.Step(*f)
		ss := p.ss
		ss.head.Push(18)
		ss.head.Push(34)
		ss.Begin(p)
		require.Equal(t, expected, ss)
		x, y := p.GetStorageOffset()
		require.Equal(t, 2, x)
		require.Equal(t, 0, y)
	})
	t.Run("normal", func(t *testing.T) {
		expected := NewStackStack()
		expected.head = &Stack{size: 4, height: 4, data: []int{12, 14, -2, 5}, next: expected.head}
		expected.head.next.Push(7)
		expected.head.next.Push(36)
		expected.head.next.Push(42)
		expected.head.next.Push(-2)
		expected.head.next.Push(5)
		expected.head.next.Push(4)
		expected.head.next.Pop()
		expected.head.next.Pop()
		expected.head.next.Pop()
		expected.height++
		p := NewPointer()
		p.SetStorageOffset(36, 42)
		ss := p.ss
		ss.head.Push(7)
		ss.head.Push(12)
		ss.head.Push(14)
		ss.head.Push(-2)
		ss.head.Push(5)
		ss.head.Push(4)
		ss.Begin(p)
		require.Equal(t, expected, ss)
	})
}

func TestEnd(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		expected := NewStackStack()
		p := NewPointer()
		ss := p.ss
		ss.Begin(p)
		reflect := ss.End(p)
		require.Equal(t, false, reflect)
		require.Equal(t, expected, ss)
	})
	t.Run("drop", func(t *testing.T) {
		expected := NewStackStack()
		expected.head.Push(7)
		expected.head.Push(12)
		expected.head.Push(14)
		expected.head.Push(-2)
		expected.head.Push(5)
		expected.head.Pop()
		expected.head.Pop()
		expected.head.Pop()
		p := NewPointer()
		ss := p.ss
		ss.head.Push(7)
		ss.head.Push(12)
		ss.head.Push(14)
		ss.head.Push(-2)
		ss.head.Push(5)
		ss.head.Push(0)
		ss.Begin(p)
		ss.head.Push(18)
		ss.head.Push(42)
		ss.head.Push(-3)
		reflect := ss.End(p)
		require.Equal(t, false, reflect)
		require.Equal(t, expected, ss)
	})
	t.Run("drop too much", func(t *testing.T) {
		expected := NewStackStack()
		p := NewPointer()
		ss := p.ss
		ss.Begin(p)
		ss.head.Push(-3)
		reflect := ss.End(p)
		require.Equal(t, false, reflect)
		require.Equal(t, expected, ss)
	})
	t.Run("reflect", func(t *testing.T) {
		expected := NewStackStack()
		p := NewPointer()
		ss := p.ss
		reflect := ss.End(p)
		require.Equal(t, true, reflect)
		require.Equal(t, expected, ss)
	})
	t.Run("transfert", func(t *testing.T) {
		expected := NewStackStack()
		expected.head.size = 5
		expected.head.data = make([]int, 5)
		expected.head.Push(7)
		expected.head.Push(12)
		expected.head.Push(14)
		expected.head.Push(-2)
		expected.head.Push(5)
		p := NewPointer()
		ss := p.ss
		ss.head.size = 4
		ss.head.data = make([]int, 4)
		ss.head.Push(7)
		ss.head.Push(0)
		ss.Begin(p)
		ss.head.Push(0)
		ss.head.Push(18)
		ss.head.Push(42)
		ss.head.Push(7)
		ss.head.Push(12)
		ss.head.Push(14)
		ss.head.Push(-2)
		ss.head.Push(5)
		ss.head.Push(4)
		reflect := ss.End(p)
		require.Equal(t, false, reflect)
		require.Equal(t, expected, ss)
	})
}

func TestUnder(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		expected := NewStackStack()
		p := NewPointer()
		reflect := p.ss.Under()
		require.Equal(t, true, reflect)
		require.Equal(t, expected, p.ss)
	})
	t.Run("positive", func(t *testing.T) {
		pe := NewPointer()
		expected := NewStackStack()
		expected.head.Push(1)
		expected.head.Push(2)
		expected.head.Push(3)
		expected.head.Push(6)
		expected.head.Push(0)
		expected.Begin(pe)
		expected.head.Push(0)
		expected.head.Push(0)
		expected.head.Push(6)
		expected.head.next.Pop()
		expected.head.next.Pop()
		expected.head.next.Pop()
		p := NewPointer()
		ss := p.ss
		ss.head.Push(1)
		ss.head.Push(2)
		ss.head.Push(3)
		ss.head.Push(6)
		ss.head.Push(0)
		ss.Begin(p)
		ss.head.Push(3)
		reflect := ss.Under()
		require.Equal(t, false, reflect)
		require.Equal(t, expected, ss)
	})
	t.Run("negative", func(t *testing.T) {
		pe := NewPointer()
		expected := NewStackStack()
		expected.Begin(pe)
		expected.head.next.Push(12)
		expected.head.next.Push(5)
		expected.head.next.Push(8)
		expected.head.Push(8)
		expected.head.Push(5)
		expected.head.Push(12)
		expected.head.Push(-3)
		expected.head.Pop()
		expected.head.Pop()
		expected.head.Pop()
		expected.head.Pop()
		p := NewPointer()
		ss := p.ss
		ss.Begin(p)
		ss.head.Push(8)
		ss.head.Push(5)
		ss.head.Push(12)
		ss.head.Push(-3)
		reflect := ss.Under()
		require.Equal(t, false, reflect)
		require.Equal(t, expected, ss)
	})
}
