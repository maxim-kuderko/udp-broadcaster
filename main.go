package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
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
		buff := make([]byte, 100)
		last := 0
		c := atomic.Int32{}
		rate := atomic.Int32{}
		var latency time.Duration
		go func() {
			for range time.NewTicker(time.Second).C {
				fmt.Printf("rate %d, out of order %d/s, latency %d\n", rate.Swap(0), c.Swap(0), latency.Milliseconds())
			}
		}()
		go func() {
			for range time.NewTicker(time.Second).C {
				conn.Write([]byte(fmt.Sprintf("l%d", time.Now().UnixMilli())))
			}
		}()
		for {
			n, _, _ := conn.ReadFrom(buff[:])
			if buff[0] == 'l' {
				ts, _ := strconv.Atoi(string(buff[1:n]))
				latency = time.Since(time.UnixMilli(int64(ts)))
				continue
			}
			current, _ := strconv.Atoi(string(buff[:n]))
			rate.Add(1)
			if last+1 != current {
				c.Add(1)
			}
			last = current
		}
	}()
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Kill, os.Interrupt)
	fmt.Println(<-done)

}
