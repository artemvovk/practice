package server

import (
	"fmt"
	"testing"
	"time"
)

const (
	host    = "localhost"
	udpPort = "8886"
	wsPort  = "8887"
	rpcPort = "8888"
)

// RPC only allows one server at a time
// benchmarks run in parallel so I must
// init the server as a singleton
func init() {
	go func() {
		ServerRPC(host, rpcPort)
	}()

	time.Sleep(time.Millisecond * 10)
}

func BenchmarkUDPServer(b *testing.B) {
	go ServerUDP(host, udpPort)
	// wait for server to start
	time.Sleep(time.Millisecond * 10)

	for n := 0; n < b.N; n++ {
		// UDP doesn't do hostname resolution
		ClientUDP("127.0.0.1", udpPort, fmt.Sprintf("test-%v", n))
	}
}

func BenchmarkWebSocketServer(b *testing.B) {
	go ServerWebSocket(host, wsPort)
	time.Sleep(time.Millisecond * 10)
	for n := 0; n < b.N; n++ {
		go ClientWebSocket(host, wsPort, "test message", n)
	}
	// Just wait for terminal io
	time.Sleep(time.Millisecond * 100)
}

func BenchmarkRPCServer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		go ClientRPC(host, rpcPort, fmt.Sprintf("Tester-%v", n))
		time.Sleep(time.Millisecond * 10)
	}
}
