package server

import (
	"testing"
	"time"
)

func TestUdpServer(t *testing.T) {
	host := "127.0.0.1"
	port := "8080"
	go ServerUDP(host, port)
	time.Sleep(time.Millisecond * 10)
	ClientUDP(host, port, "test")
}
