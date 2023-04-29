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
		buff := make([]byte, 100)
		last := 0
		for {
			n, _, _ := conn.ReadFrom(buff[:])
			current, _ := strconv.Atoi(string(buff[:n]))
			if last+1 != current {
				fmt.Println(fmt.Errorf("got %s wanted %d", string(buff[:n]), last+1))
			}
			last = current
		}
	}()
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Kill, os.Interrupt)
	fmt.Println(<-done)

}
