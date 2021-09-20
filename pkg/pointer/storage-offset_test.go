package pointer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetStorageOffset(t *testing.T) {
	p := NewPointer()
	x, y := p.GetStorageOffset()
	require.Equal(t, x, 0)
	require.Equal(t, y, 0)
	p.SetStorageOffset(3, 8)
	x, y = p.GetStorageOffset()
	require.Equal(t, x, 3)
	require.Equal(t, y, 8)
}

func TestCalculateNewStorageOffset(t *testing.T) {
	p := NewPointer()
	p.CalculateNewStorageOffset()
	x, y := p.GetStorageOffset()
	require.Equal(t, x, 1)
	require.Equal(t, y, 0)
	p.sox, p.soy = 3, 2
	x, y = p.GetStorageOffset()
	require.Equal(t, x, 3)
	require.Equal(t, y, 2)
	p.x, p.y = 8, 12
	p.CalculateNewStorageOffset()
	x, y = p.GetStorageOffset()
	require.Equal(t, x, 9)
	require.Equal(t, y, 12)
}
