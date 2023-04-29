package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync/atomic"
	"time"
)

func main() {
	go func() {
		conn, err := net.ListenUDP("udp", &net.UDPAddr{
			Port: 5000,
		})
		if err != nil {
			panic(err)
		}
		conn.SetReadBuffer(1024 * 1024)
		srv := &Server{
			conn: conn,
			fn: func(data []byte, addr net.Addr) {
				//log.Println(string(data))
			}}
		c := atomic.Int32{}
		go func() {
			for range time.NewTicker(time.Second).C {
				fmt.Printf("%d\n", c.Swap(0))
			}
		}()
		globalBuff := make([]byte, 100)
		for {
			n, addr, _ := conn.ReadFrom(globalBuff[:])
			srv.Serve(globalBuff[:n], addr)
			c.Add(1)
		}
	}()
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Kill, os.Interrupt)
	fmt.Println(<-done)

}
