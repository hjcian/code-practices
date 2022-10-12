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
