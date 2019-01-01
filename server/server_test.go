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
