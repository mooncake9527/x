package xcode_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	c := xcode.New(1, "custom error", "detailed description")
	assert.Equal(t, c.Code(), 1)
	assert.Equal(t, c.Message(), "custom error")
	assert.Equal(t, c.Detail(), "detailed description")
}
