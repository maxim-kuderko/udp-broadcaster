package main

import (
	"net"
)

type Server struct {
	fn   func(data []byte, addr net.Addr)
	conn *net.UDPConn
}

func (s *Server) Serve(data []byte, addr net.Addr) {
	s.fn(data, addr)
}
