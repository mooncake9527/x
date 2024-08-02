package xtcp_test

import (
	"github.com/mooncake9527/x/xnet/xtcp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFreePort(t *testing.T) {
	_, err := xtcp.GetFreePort()
	if err != nil {
		t.Error(err)
	}
}

func TestGetFreePorts(t *testing.T) {
	ports, err := xtcp.GetFreePorts(2)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, len(ports), 2)
}
