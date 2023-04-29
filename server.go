package main

import (
	"github.com/valyala/bytebufferpool"
	"net"
)

type Server struct {
	fn   func(data []byte, addr net.Addr)
	conn *net.UDPConn
}

func (s *Server) Serve(data []byte, addr net.Addr) {
	buff := bytebufferpool.Get()
	defer bytebufferpool.Put(buff)
	buff.Write(data)
	s.fn(buff.Bytes(), addr)
}
