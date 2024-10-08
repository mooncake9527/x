package xtcp

import (
	"context"
	"crypto/tls"
	"errors"
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

// Server 定义 TCP 服务包装器。
type Server struct {
	err       error
	mu        sync.Mutex   // 用于 Server.listen 并发安全。
	listen    net.Listener // 网络监听器。
	network   string       // 服务器网络协议。
	address   string       // 服务器监听地址。
	handler   func(*Conn)  // 连接处理器。
	tlsConfig *tls.Config  // TLS 配置。
}

// NewServer 新建 TCP 服务器。
func NewServer(address string, handler func(*Conn)) *Server {
	srv := &Server{
		network: "tcp",
		address: address,
		handler: handler,
	}
	srv.err = srv.listener()
	return srv
}

// NewServerTLS 新建 TCP TLS 服务器。
func NewServerTLS(address string, tlsConfig *tls.Config, handler func(*Conn)) *Server {
	srv := &Server{
		network:   "tcp",
		address:   address,
		handler:   handler,
		tlsConfig: tlsConfig,
	}
	srv.err = srv.listener()
	return srv
}

// listener 网络监听器。
func (s *Server) listener() (err error) {
	if s.handler == nil {
		err = xerror.NewCode(xcode.CodeMissingConfiguration, "start running failed: socket handler not defined")
		return
	}
	if s.tlsConfig != nil {
		// TLS Server
		s.mu.Lock()
		s.listen, err = tls.Listen(s.network, s.address, s.tlsConfig)
		s.mu.Unlock()
		if err != nil {
			err = xerror.Wrapf(err, `tls.Listen failed for address "%s"`, s.address)
			return
		}
	} else {
		// Normal Server
		var tcpAddr *net.TCPAddr
		if tcpAddr, err = net.ResolveTCPAddr(s.network, s.address); err != nil {
			err = xerror.Wrapf(err, `net.ResolveTCPAddr failed for address "%s"`, s.address)
			return err
		}
		s.mu.Lock()
		s.listen, err = net.ListenTCP(s.network, tcpAddr)
		s.mu.Unlock()
		if err != nil {
			err = xerror.Wrapf(err, `net.ListenTCP failed for address "%s"`, s.address)
			return err
		}
	}
	return nil
}

// Close 关闭 TCP 服务器。
func (s *Server) Close(ctx context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.listen == nil {
		return nil
	}
	return s.listen.Close()
}

// Run 启动 TCP 服务器。
func (s *Server) Run(ctx context.Context) (err error) {
	if s.err != nil {
		return s.err
	}
	if s.listen == nil {
		err = errors.New("xtcp start running failed: socket Listener not defined")
		return
	}
	if s.handler == nil {
		err = errors.New("xtcp start running failed: socket handler not defined")
		return
	}
	// Listening loop.
	for {
		var conn net.Conn
		if conn, err = s.listen.Accept(); err != nil {
			err = xerror.Wrapf(err, `Listener.Accept failed`)
			return err
		} else if conn != nil {
			go s.handler(NewConnByNetConn(conn))
		}
	}
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
	if ln := s.listen; ln != nil {
		return ln.Addr().(*net.TCPAddr).Port
	}
	return -1
}
