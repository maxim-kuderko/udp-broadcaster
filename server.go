package main

import (
	"github.com/valyala/bytebufferpool"
	"net"
	"sync"
)

type Server struct {
	connections *sync.Map
	conn        *net.UDPConn
	broadcastC  chan []byte
}

func NewServer(c *net.UDPConn) *Server {
	s := &Server{
		connections: &sync.Map{},
		conn:        c,
		broadcastC:  make(chan []byte, 1),
	}
	go s.broadcast()
	return s
}

func (s *Server) broadcast() {
	for d := range s.broadcastC {
		s.connections.Range(func(key, value any) bool {
			s.conn.WriteTo(d, value.(net.Addr))
			return true
		})
	}
}

func (s *Server) Serve(data []byte, addr net.Addr) {
	buff := bytebufferpool.Get()
	defer bytebufferpool.Put(buff)
	buff.Write(data)
	s.serve(data, addr)
}

func (s *Server) serve(data []byte, addr net.Addr) {
	s.broadcastC <- data
	s.connections.Store(addr.String(), addr)
}
