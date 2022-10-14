package implementqueueusingstacks

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_0232_ImplementQueueUsingStacks(t *testing.T) {
	q := Constructor()
	q.Push(1)
	q.Push(2)
	require.Equal(t, 1, q.Peek())
	require.Equal(t, 1, q.Pop())
	require.Equal(t, 2, q.Peek())
	require.False(t, q.Empty())
	require.Equal(t, 2, q.Pop())
	require.True(t, q.Empty())
}

func Test_0232_ImplementQueueUsingStacks_17_22(t *testing.T) {
	q := Constructor()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	require.Equal(t, 1, q.Pop())
	q.Push(5)
	require.Equal(t, 2, q.Pop())
	require.Equal(t, 3, q.Pop())
	require.Equal(t, 4, q.Pop())
	require.Equal(t, 5, q.Pop())
}
