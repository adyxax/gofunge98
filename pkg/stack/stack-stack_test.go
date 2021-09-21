package stack

import (
	"os"
	"testing"

	"git.adyxax.org/adyxax/gofunge/pkg/field"
	"git.adyxax.org/adyxax/gofunge/pkg/pointer"
	"github.com/stretchr/testify/require"
)

func TestBegin(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		expected := NewStackStack()
		expected.head = &Stack{data: make([]int, 0), next: expected.head}
		expected.head.next.Push(0)
		expected.head.next.Push(0)
		expected.height++
		ss := NewStackStack()
		p := pointer.NewPointer()
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
		expected.head = &Stack{size: 5, height: 5, data: make([]int, 5), next: expected.head}
		expected.head.next.Push(0)
		expected.head.next.Push(0)
		expected.height++
		p := pointer.NewPointer()
		file, err := os.Open("../field/test_data/hello.b98")
		require.NoError(t, err, "Failed to open file")
		f, err := field.Load(file)
		p.Step(*f)
		ss := NewStackStack()
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
		expected.head.next.Push(18)
		expected.head.next.Push(2)
		expected.head.next.Push(3)
		expected.height++
		p := pointer.NewPointer()
		p.SetStorageOffset(2, 3)
		file, err := os.Open("../field/test_data/hello.b98")
		require.NoError(t, err, "Failed to open file")
		f, err := field.Load(file)
		p.Step(*f)
		ss := NewStackStack()
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
		expected.head.next.Push(12)
		expected.head.next.Push(14)
		expected.head.next.Push(-2)
		expected.head.next.Push(5)
		expected.head.next.Push(36)
		expected.head.next.Push(42)
		expected.height++
		p := pointer.NewPointer()
		p.SetStorageOffset(36, 42)
		ss := NewStackStack()
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
