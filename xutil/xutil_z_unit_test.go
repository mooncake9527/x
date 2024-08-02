package xutil_test

import (
	"github.com/mooncake9527/x/xutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInArray(t *testing.T) {
	assert.True(t, xutil.InArray(1, []int{1, 2, 3, 4}), true)
	assert.False(t, xutil.InArray(5, []int{1, 2, 3, 4}), true)
	assert.True(t, xutil.InArray("cat", []string{"dog", "cat", "pig"}), true)
	assert.False(t, xutil.InArray("monkey", []string{"dog", "cat", "pig"}), true)
}
