package xudp

import (
	"context"
	"fmt"
	"github.com/mooncake9527/x/xerrors/xcode"
	"github.com/mooncake9527/x/xerrors/xerror"
	"net"
	"strings"
	"sync"
)

const (
	// FreePortAddress 使用随机端口标记服务器监听。
	FreePortAddress = ":0"
)

// Server 定义 UDP 服务器。
type Server struct {
	err     error
	mu      sync.Mutex  // 用于 Server.Conn 并发安全。
	conn    *Conn       // UDP 服务器连接对象。
	network string      // UDP 服务器网络协议。
	address string      // UDP 服务器监听地址。
	handler func(*Conn) // UDP 连接的处理程序。
}

// NewServer 新建 UDP 服务器。
func NewServer(address string, handler func(*Conn)) *Server {
	srv := &Server{
		network: "udp",
		address: address,
		handler: handler,
	}
	srv.err = srv.listener()
	return srv
}

// listener 网络监听器。
func (s *Server) listener() (err error) {
	if s.handler == nil {
		err := xerror.NewCode(xcode.CodeMissingConfiguration, "start running failed: socket handler not defined")
		return err
	}
	addr, err := net.ResolveUDPAddr(s.network, s.address)
	if err != nil {
		err = xerror.Wrapf(err, `net.ResolveUDPAddr failed for address "%s"`, s.address)
		return err
	}
	conn, err := net.ListenUDP(s.network, addr)
	if err != nil {
		err = xerror.Wrapf(err, `net.ListenUDP failed for address "%s"`, s.address)
		return err
	}
	s.mu.Lock()
	s.conn = NewConnByNetConn(conn)
	s.mu.Unlock()
	return nil
}

// Close 关闭 UDP 服务器。
func (s *Server) Close(ctx context.Context) (err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	err = s.conn.Close()
	if err != nil {
		err = xerror.Wrap(err, "connection failed")
	}
	return
}

// Run 启动 UDP 服务器。
func (s *Server) Run(ctx context.Context) (err error) {
	if s.err != nil {
		return s.err
	}
	s.handler(s.conn)
	return nil
}

// GetListenedAddress 获取当前服务器监听地址。
func (s *Server) GetListenedAddress() string {
	if !strings.Contains(s.address, FreePortAddress) {
		return s.address
	}
	var (
		address      = s.address
		listenedPort = s.GetListenedPort()
	)
	address = strings.Replace(address, FreePortAddress, fmt.Sprintf(`:%d`, listenedPort), -1)
	return address
}

// GetListenedPort 获取当前服务器监听端口。
func (s *Server) GetListenedPort() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	if ln := s.conn; ln != nil {
		return ln.LocalAddr().(*net.UDPAddr).Port
	}
	return -1
}
