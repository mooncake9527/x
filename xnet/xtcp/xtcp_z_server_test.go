package xtcp_test

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	simpleTimeout = time.Millisecond * 100
	sendData      = []byte("hello")
	invalidAddr   = "127.0.0.1:99999"
)

func startTCPServer(addr string) *xtcp.Server {
	ctx := context.Background()
	s := xtcp.NewServer(addr, func(conn *xtcp.Conn) {
		defer conn.Close()
		for {
			data, err := conn.Receive(-1)
			if err != nil {
				break
			}
			conn.Send(data)
		}
	})
	go s.Run(ctx)
	time.Sleep(simpleTimeout)
	return s
}

func startTCPPkgServer(addr string) *xtcp.Server {
	ctx := context.Background()
	s := xtcp.NewServer(addr, func(conn *xtcp.Conn) {
		defer conn.Close()
		for {
			data, err := conn.ReceivePkg()
			if err != nil {
				break
			}
			conn.SendPkg(data)
		}
	})
	go s.Run(ctx)
	time.Sleep(simpleTimeout)
	return s
}

func TestConnGetFreePorts(t *testing.T) {
	ports, _ := xtcp.GetFreePorts(2)
	assert.Greater(t, ports[0], 0)
	assert.Greater(t, ports[1], 0)

	startTCPServer(fmt.Sprintf("%s:%d", "127.0.0.1", ports[0]))

	conn, err := xtcp.NewConn(fmt.Sprintf("127.0.0.1:%d", ports[0]))
	assert.Nil(t, err)
	defer conn.Close()
	result, err := conn.SendReceive(sendData, -1)
	assert.Nil(t, err)
	assert.Equal(t, result, sendData)

	conn1, err1 := xtcp.NewConn(fmt.Sprintf("127.0.0.1:%d", 80))
	assert.NotNil(t, err1)
	assert.Nil(t, conn1)
}

func TestConnMustGetFreePort(t *testing.T) {
	port := xtcp.MustGetFreePort()
	addr := fmt.Sprintf("%s:%d", "127.0.0.1", port)
	startTCPServer(addr)

	result, err := xtcp.SendReceive(addr, sendData, -1)
	assert.Nil(t, err)
	assert.Equal(t, sendData, result)
}

func TestNewConn(t *testing.T) {
	addr := xtcp.FreePortAddress

	conn, err := xtcp.NewConn(addr, simpleTimeout)
	assert.Nil(t, conn)
	assert.NotNil(t, err)

	s := startTCPServer(xtcp.FreePortAddress)

	conn1, err1 := xtcp.NewConn(s.GetListenedAddress(), simpleTimeout)
	assert.Nil(t, err1)
	assert.NotEqual(t, conn1, nil)
	defer conn1.Close()
	result1, err1 := conn1.SendReceive(sendData, -1)
	assert.Nil(t, err1)
	assert.Equal(t, result1, sendData)
}

func TestConn_Send(t *testing.T) {
	s := startTCPServer(xtcp.FreePortAddress)

	conn, err := xtcp.NewConn(s.GetListenedAddress())
	assert.Nil(t, err)
	assert.NotNil(t, conn)
	err = conn.Send(sendData, xtcp.Retry{Count: 1})
	assert.Nil(t, err)
	result, err := conn.Receive(-1)
	assert.Nil(t, err)
	assert.Equal(t, result, sendData)
}

func TestConn_SendWithTimeout(t *testing.T) {
	s := startTCPServer(xtcp.FreePortAddress)

	conn, err := xtcp.NewConn(s.GetListenedAddress())
	assert.Nil(t, err)
	assert.NotNil(t, conn)
	err = conn.SendWithTimeout(sendData, time.Second, xtcp.Retry{Count: 1})
	assert.Nil(t, err)
	result, err := conn.Receive(-1)
	assert.Nil(t, err)
	assert.Equal(t, result, sendData)
}

func TestConn_SendReceive(t *testing.T) {
	s := startTCPServer(xtcp.FreePortAddress)

	conn, err := xtcp.NewConn(s.GetListenedAddress())
	assert.Nil(t, err)
	assert.NotNil(t, conn)
	result, err := conn.SendReceive(sendData, -1, xtcp.Retry{Count: 1})
	assert.Nil(t, err)
	assert.Equal(t, result, sendData)
}

func TestConn_SendReceiveWithTimeout(t *testing.T) {
	s := startTCPServer(xtcp.FreePortAddress)

	conn, err := xtcp.NewConn(s.GetListenedAddress())
	assert.Nil(t, err)
	assert.NotNil(t, conn)
	result, err := conn.SendReceiveWithTimeout(sendData, -1, time.Second, xtcp.Retry{Count: 1})
	assert.Nil(t, err)
	assert.Equal(t, result, sendData)
}

func TestConn_ReceiveWithTimeout(t *testing.T) {
	s := startTCPServer(xtcp.FreePortAddress)

	conn, err := xtcp.NewConn(s.GetListenedAddress())
	assert.Nil(t, err)
	assert.NotNil(t, conn)
	conn.Send(sendData)
	result, err := conn.ReceiveWithTimeout(-1, time.Second, xtcp.Retry{Count: 1})
	assert.Nil(t, err)
	assert.Equal(t, result, sendData)
}

func TestConn_ReceiveLine(t *testing.T) {
	s := startTCPServer(xtcp.FreePortAddress)

	conn, err := xtcp.NewConn(s.GetListenedAddress())
	assert.Nil(t, err)
	assert.NotNil(t, conn)
	data := []byte("hello\n")
	conn.Send(data)
	result, err := conn.ReceiveLine(xtcp.Retry{Count: 1})
	assert.Nil(t, err)
	splitData := strings.Split(string(data), "\n")
	assert.Equal(t, string(result), splitData[0])
}

func TestConn_ReceiveTill(t *testing.T) {
	s := startTCPServer(xtcp.FreePortAddress)

	conn, err := xtcp.NewConn(s.GetListenedAddress())
	assert.Nil(t, err)
	assert.NotNil(t, conn)
	conn.Send(sendData)
	result, err := conn.ReceiveTill([]byte("hello"), xtcp.Retry{Count: 1})
	assert.Nil(t, err)
	assert.Equal(t, result, sendData)
}

func TestConn_SetDeadline(t *testing.T) {
	s := startTCPServer(xtcp.FreePortAddress)

	conn, err := xtcp.NewConn(s.GetListenedAddress())
	assert.Nil(t, err)
	assert.NotNil(t, conn)
	conn.SetDeadline(time.Time{})
	err = conn.Send(sendData, xtcp.Retry{Count: 1})
	assert.Nil(t, err)
	result, err := conn.Receive(-1)
	assert.Nil(t, err)
	assert.Equal(t, result, sendData)
}

func TestConn_SetReceiveBufferWait(t *testing.T) {
	s := startTCPServer(xtcp.FreePortAddress)

	conn, err := xtcp.NewConn(s.GetListenedAddress())
	assert.Nil(t, err)
	assert.NotNil(t, conn)
	conn.SetReceiveBufferWait(time.Millisecond * 100)
	err = conn.Send(sendData, xtcp.Retry{Count: 1})
	assert.Nil(t, err)
	result, err := conn.Receive(-1)
	assert.Nil(t, err)
	assert.Equal(t, result, sendData)
}

func TestSend(t *testing.T) {
	s := startTCPServer(xtcp.FreePortAddress)

	err1 := xtcp.Send(invalidAddr, sendData, xtcp.Retry{Count: 1})
	assert.NotNil(t, err1)

	err2 := xtcp.Send(s.GetListenedAddress(), sendData, xtcp.Retry{Count: 1})
	assert.Nil(t, err2)
}

func TestSendReceive(t *testing.T) {
	s := startTCPServer(xtcp.FreePortAddress)

	result1, err1 := xtcp.SendReceive(invalidAddr, sendData, -1)
	assert.NotNil(t, err1)
	assert.Nil(t, result1)

	result2, err2 := xtcp.SendReceive(s.GetListenedAddress(), sendData, -1)
	assert.Nil(t, err2)
	assert.Equal(t, result2, sendData)
}

func TestSendWithTimeout(t *testing.T) {
	s := startTCPServer(xtcp.FreePortAddress)

	err := xtcp.SendWithTimeout(invalidAddr, sendData, time.Millisecond*500)
	assert.NotNil(t, err)
	err = xtcp.SendWithTimeout(s.GetListenedAddress(), sendData, time.Millisecond*500)
	assert.Nil(t, err)
}

func TestSendReceiveWithTimeout(t *testing.T) {
	s := startTCPServer(xtcp.FreePortAddress)

	result, err := xtcp.SendReceiveWithTimeout(invalidAddr, sendData, -1, time.Millisecond*500)
	assert.Nil(t, result)
	assert.NotNil(t, err)
	result, err = xtcp.SendReceiveWithTimeout(s.GetListenedAddress(), sendData, -1, time.Millisecond*500)
	assert.Nil(t, err)
	assert.Equal(t, result, sendData)
}

func TestSendPkg(t *testing.T) {
	s := startTCPPkgServer(xtcp.FreePortAddress)

	err1 := xtcp.SendPkg(s.GetListenedAddress(), sendData)
	assert.Nil(t, err1)
	err1 = xtcp.SendPkg(invalidAddr, sendData)
	assert.NotNil(t, err1)

	err2 := xtcp.SendPkg(s.GetListenedAddress(), sendData, xtcp.PkgOption{Retry: xtcp.Retry{Count: 3}})
	assert.Nil(t, err2)
	err2 = xtcp.SendPkg(s.GetListenedAddress(), sendData)
	assert.Nil(t, err2)
}

func TestSendReceivePkg(t *testing.T) {
	s := startTCPPkgServer(xtcp.FreePortAddress)

	err1 := xtcp.SendPkg(s.GetListenedAddress(), sendData)
	assert.Nil(t, err1)
	_, err1 = xtcp.SendReceivePkg(invalidAddr, sendData)
	assert.NotNil(t, err1)

	err2 := xtcp.SendPkg(s.GetListenedAddress(), sendData)
	assert.Nil(t, err2)
	result, err2 := xtcp.SendReceivePkg(s.GetListenedAddress(), sendData)
	assert.Nil(t, err2)
	assert.Equal(t, result, sendData)
}

func TestSendPkgWithTimeout(t *testing.T) {
	s := startTCPPkgServer(xtcp.FreePortAddress)

	err1 := xtcp.SendPkg(s.GetListenedAddress(), sendData)
	assert.Nil(t, err1)
	err1 = xtcp.SendPkgWithTimeout(invalidAddr, sendData, time.Second)
	assert.NotNil(t, err1)

	err2 := xtcp.SendPkg(s.GetListenedAddress(), sendData)
	assert.Nil(t, err2)
	err2 = xtcp.SendPkgWithTimeout(s.GetListenedAddress(), sendData, time.Second)
	assert.Nil(t, err2)
}

func TestSendReceivePkgWithTimeout(t *testing.T) {
	s := startTCPPkgServer(xtcp.FreePortAddress)

	err1 := xtcp.SendPkg(s.GetListenedAddress(), sendData)
	assert.Nil(t, err1)
	_, err1 = xtcp.SendReceivePkgWithTimeout(invalidAddr, sendData, time.Second)
	assert.NotNil(t, err1)

	err2 := xtcp.SendPkg(s.GetListenedAddress(), sendData)
	assert.Nil(t, err2)
	result, err2 := xtcp.SendReceivePkgWithTimeout(s.GetListenedAddress(), sendData, time.Second)
	assert.Nil(t, err2)
	assert.Equal(t, result, sendData)
}

func TestNewServer(t *testing.T) {
	ctx := context.Background()
	s := xtcp.NewServer(xtcp.FreePortAddress, func(conn *xtcp.Conn) {
		defer conn.Close()
		for {
			data, err := conn.Receive(-1)
			if err != nil {
				break
			}
			conn.Send(data)
		}
	})
	defer s.Close(ctx)
	go s.Run(ctx)

	time.Sleep(simpleTimeout)

	result, err := xtcp.SendReceive(s.GetListenedAddress(), sendData, -1)
	assert.Nil(t, err)
	assert.Equal(t, result, sendData)
}

func TestServer_Run(t *testing.T) {
	ctx := context.Background()
	s := xtcp.NewServer(xtcp.FreePortAddress, func(conn *xtcp.Conn) {
		defer conn.Close()
		for {
			data, err := conn.Receive(-1)
			if err != nil {
				break
			}
			conn.Send(data)
		}
	})
	defer s.Close(ctx)
	go s.Run(ctx)

	time.Sleep(simpleTimeout)

	result, err := xtcp.SendReceive(s.GetListenedAddress(), sendData, -1)
	assert.Nil(t, err)
	assert.Equal(t, result, sendData)

	s1 := xtcp.NewServer(xtcp.FreePortAddress, nil)
	defer s1.Close(ctx)
	go func() {
		err := s1.Run(ctx)
		assert.NotNil(t, err)
	}()
}
