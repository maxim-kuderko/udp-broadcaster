package main

import (
	"log"
	"net"
	"strconv"
	"time"
)

func main() {
	run()
}

func run() {
	udpServer, err := net.ResolveUDPAddr("udp", ":5000")

	if err != nil {
		log.Panicf("ResolveUDPAddr failed: %s", err.Error())
	}

	conn, err := net.DialUDP("udp", nil, udpServer)
	if err != nil {
		log.Panicf("Listen failed: %s", err.Error())
	}
	i := 1
	go func() {
		b := make([]byte, 100)
		for {
			n, _ := conn.Read(b[:])
			conn.Write(b[:n])
		}
	}()
	for {
		conn.Write([]byte(strconv.Itoa(i)))
		i++
		time.Sleep(time.Millisecond)
	}
}
