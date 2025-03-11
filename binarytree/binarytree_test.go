package binarytree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromHeapSlice(t *testing.T) {
	tree := FromHeapSlice([]any{3, 9, 20, nil, nil, 15, 7})
	assert.Equal(t, 3, tree.Value)
	assert.Equal(t, 9, tree.Left.Value)
	assert.Equal(t, 20, tree.Right.Value)
	assert.Nil(t, tree.Left.Left)
	assert.Nil(t, tree.Left.Right)
	assert.Equal(t, 15, tree.Right.Left.Value)
	assert.Equal(t, 7, tree.Right.Right.Value)
}
