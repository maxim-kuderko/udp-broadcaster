package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
)

func main() {
	go func() {
		conn, err := net.ListenUDP("udp", &net.UDPAddr{
			Port: 5000,
		})
		if err != nil {
			panic(err)
		}
		globalBuff := make([]byte, 100)
		last := 0
		for {
			n, _, _ := conn.ReadFrom(globalBuff[:])
			current, _ := strconv.Atoi(string(globalBuff[:n]))
			if last+1 != current {
				fmt.Println(fmt.Errorf("got %s wanted %d", string(globalBuff[:n]), current+1))
			}
		}
	}()
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Kill, os.Interrupt)
	fmt.Println(<-done)

}
