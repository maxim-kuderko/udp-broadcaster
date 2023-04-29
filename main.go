package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		Port: 5000,
	})
	if err != nil {
		panic(err)
	}
	srv := &Server{
		conn: conn,
		fn: func(data []byte, addr net.Addr) {
			fmt.Println(string(data))
		}}
	globalBuff := make([]byte, 4096)
	for {
		n, addr, _ := conn.ReadFrom(globalBuff)
		srv.Serve(globalBuff[:n], addr)
	}
}
