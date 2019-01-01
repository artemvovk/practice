package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const maxBufferSize = 1024

var (
	errHandler = make(chan error, 1)
	addresses  = make(chan *net.UDPAddr, 10)
)

func ServerUDP(host, port string) error {
	udpAddr, err := net.ResolveUDPAddr("udp", net.JoinHostPort(host, port))
	if err != nil {
		return err
	}
	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return err
	}
	signalChan := make(chan os.Signal, 1)
	handleUDP(udpConn)
	for {
		<-signalChan
		log.Printf("Closing connection")
		close(signalChan)
		close(addresses)
		udpConn.Close()
	}

	return err
}

func ClientUDP(host, port, message string) error {
	packet := make([]byte, maxBufferSize)
	conn, err := net.Dial("udp", net.JoinHostPort(host, port))
	defer conn.Close()
	if err != nil {
		return err
	}
	fmt.Fprintf(conn, message)
	_, err = bufio.NewReader(conn).Read(packet)
	if err == nil {
		log.Printf("%s\n", packet)
	} else {
		return err
	}
	return nil
}

func handleUDP(pc *net.UDPConn) {

	for {
		message := make([]byte, maxBufferSize)
		n, addr, err := pc.ReadFromUDP(message)
		addresses <- addr

		if err != nil {
			errHandler <- err
			return
		}
		go func() {
			<-addresses
			log.Printf("Received: bytes=%d from=%s\n",
				n, addr.String())
			log.Printf("Packet: %v", string(message))
			n, err = pc.WriteTo(message[:n], addr)
			if err != nil {
				errHandler <- err
				return
			}
			log.Printf("Sent: bytes=%d to=%s\n", n, addr.String())
		}()

	}
}
