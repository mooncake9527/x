package xudp_test

import (
	"context"
	"fmt"
	"github.com/mooncake9527/x/xerrors/xerror"
	"github.com/mooncake9527/x/xnet/xudp"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	simpleTimeout = time.Millisecond * 100
	sendData      = []byte("hello")
)

func startUDPServer(addr string) *xudp.Server {
	ctx := context.Background()
	s := xudp.NewServer(addr, func(conn *xudp.Conn) {
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

func Test_Basic(t *testing.T) {
	var ctx = context.TODO()
	s := xudp.NewServer(xudp.FreePortAddress, func(conn *xudp.Conn) {
		defer conn.Close()
		for {
			data, err := conn.Receive(-1)
			if len(data) > 0 {
				if err := conn.Send(append([]byte("> "), data...)); err != nil {
					t.Error(xerror.Stack(err))
				}
			}
			if err != nil {
				break
			}
		}
	})
	go s.Run(ctx)
	defer s.Close(ctx)
	time.Sleep(100 * time.Millisecond)

	// xudp.Conn.Send
	for i := 0; i < 100; i++ {
		conn, err := xudp.NewConn(s.GetListenedAddress())
		assert.Nil(t, err)
		assert.Nil(t, conn.Send([]byte(strconv.Itoa(i))))
		assert.Nil(t, conn.RemoteAddr())
		result, err := conn.Receive(-1)
		assert.Nil(t, err)
		assert.NotNil(t, conn.RemoteAddr())
		assert.Equal(t, string(result), fmt.Sprintf(`> %d`, i))
		conn.Close()
	}

	// xudp.Conn.SendReceive
	for i := 0; i < 100; i++ {
		conn, err := xudp.NewConn(s.GetListenedAddress())
		assert.Nil(t, err)
		_, err = conn.SendReceive([]byte(strconv.Itoa(i)), -1)
		assert.Nil(t, err)
		conn.Close()
	}

	// xudp.Conn.SendWithTimeout
	for i := 0; i < 100; i++ {
		conn, err := xudp.NewConn(s.GetListenedAddress())
		assert.Nil(t, err)
		err = conn.SendWithTimeout([]byte(strconv.Itoa(i)), time.Second)
		assert.Nil(t, err)
		conn.Close()
	}

	// xudp.Conn.ReceiveWithTimeout
	for i := 0; i < 100; i++ {
		conn, err := xudp.NewConn(s.GetListenedAddress())
		assert.Nil(t, err)
		err = conn.Send([]byte(strconv.Itoa(i)))
		assert.Nil(t, err)
		conn.SetReceiveBufferWait(time.Millisecond * 100)
		result, err := conn.ReceiveWithTimeout(-1, time.Second)
		assert.Nil(t, err)
		assert.Equal(t, string(result), fmt.Sprintf(`> %d`, i))
		conn.Close()
	}

	// xudp.Conn.SendReceiveWithTimeout
	for i := 0; i < 100; i++ {
		conn, err := xudp.NewConn(s.GetListenedAddress())
		assert.Nil(t, err)
		result, err := conn.SendReceiveWithTimeout([]byte(strconv.Itoa(i)), -1, time.Second)
		assert.Nil(t, err)
		assert.Equal(t, string(result), fmt.Sprintf(`> %d`, i))
		conn.Close()
	}

	// xudp.Send
	for i := 0; i < 100; i++ {
		err := xudp.Send(s.GetListenedAddress(), []byte(strconv.Itoa(i)))
		assert.Nil(t, err)
	}

	// xudp.SendReceive
	for i := 0; i < 100; i++ {
		result, err := xudp.SendReceive(s.GetListenedAddress(), []byte(strconv.Itoa(i)), -1)
		assert.Nil(t, err)
		assert.Equal(t, string(result), fmt.Sprintf(`> %d`, i))
	}
}

// 如果读取缓冲区大小小于已发送的软件包大小，则将删除其余数据。
func Test_Buffer(t *testing.T) {
	var ctx = context.TODO()
	s := xudp.NewServer(xudp.FreePortAddress, func(conn *xudp.Conn) {
		defer conn.Close()
		for {
			data, err := conn.Receive(-1)
			if len(data) > 0 {
				if err := conn.Send(data[:1]); err != nil {
					t.Error(xerror.Stack(err))
				}
			}
			if err != nil {
				t.Error(err)
				break
			}
		}
	})
	go s.Run(ctx)
	defer s.Close(ctx)
	time.Sleep(100 * time.Millisecond)

	result1, err1 := xudp.SendReceive(s.GetListenedAddress(), []byte("123"), -1)
	assert.Nil(t, err1)
	assert.Equal(t, string(result1), "1")

	result2, err2 := xudp.SendReceive(s.GetListenedAddress(), []byte("456"), -1)
	assert.Nil(t, err2)
	assert.Equal(t, string(result2), "4")
}

func Test_NewConn(t *testing.T) {
	s := startUDPServer(xudp.FreePortAddress)

	conn1, err1 := xudp.NewConn(s.GetListenedAddress(), fmt.Sprintf("127.0.0.1:%d", xudp.MustGetFreePort()))
	assert.Nil(t, err1)
	conn1.SetDeadline(time.Now().Add(time.Second))
	assert.Nil(t, conn1.Send(sendData))
	conn1.Close()

	conn2, err2 := xudp.NewConn(s.GetListenedAddress(), fmt.Sprintf("127.0.0.1:%d", 99999))
	assert.Nil(t, conn2)
	assert.NotNil(t, err2)

	conn3, err3 := xudp.NewConn(fmt.Sprintf("127.0.0.1:%d", 99999))
	assert.Nil(t, conn3)
	assert.NotNil(t, err3)
}

func Test_NewServer(t *testing.T) {
	s := xudp.NewServer(xudp.FreePortAddress, func(conn *xudp.Conn) {
		defer conn.Close()
		for {
			data, err := conn.Receive(1)
			if len(data) > 0 {
				conn.Send(data)
			}
			if err != nil {
				break
			}
		}
	})

	assert.NotNil(t, s)
}
