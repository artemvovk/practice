package server

import (
	"fmt"
	"testing"
	"time"
)

func BenchmarkUdpServer(b *testing.B) {
	host := "127.0.0.1"
	port := "8080"
	go ServerUDP(host, port)
	// wait for server to start
	time.Sleep(time.Millisecond * 10)

	for n := 0; n < b.N; n++ {
		ClientUDP(host, port, fmt.Sprintf("test-%v", n))
	}
}

func BenchmarkWebSocketServer(b *testing.B) {
	host := "localhost"
	port := "8080"
	go ServerWebSocket(host, port)
	time.Sleep(time.Millisecond * 10)
	for n := 0; n < b.N; n++ {
		go b.Logf("Result: %v", ClientWebSocket(host, port, "test message", n))
	}
	// Just wait for terminal io
	time.Sleep(time.Millisecond * 100)
}
