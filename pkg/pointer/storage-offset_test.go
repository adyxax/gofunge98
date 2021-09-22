package pointer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetStorageOffset(t *testing.T) {
	p := NewPointer()
	x, y := p.GetStorageOffset()
	require.Equal(t, 0, x)
	require.Equal(t, 0, y)
	p.SetStorageOffset(3, 8)
	x, y = p.GetStorageOffset()
	require.Equal(t, 3, x)
	require.Equal(t, 8, y)
}

func TestCalculateNewStorageOffset(t *testing.T) {
	p := NewPointer()
	p.CalculateNewStorageOffset()
	x, y := p.GetStorageOffset()
	require.Equal(t, 1, x)
	require.Equal(t, 0, y)
	p.sox, p.soy = 3, 2
	x, y = p.GetStorageOffset()
	require.Equal(t, 3, x)
	require.Equal(t, 2, y)
	p.x, p.y = 8, 12
	p.CalculateNewStorageOffset()
	x, y = p.GetStorageOffset()
	require.Equal(t, 9, x)
	require.Equal(t, 12, y)
}
