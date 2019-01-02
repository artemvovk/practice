package server

import (
	"errors"
	"log"
	"net"
	"net/rpc"
)

type Response struct {
	Message string
}

type Request struct {
	Name string
}

type Handler struct{}

func ServerRPC(host, port string) string {
	server := rpc.NewServer()
	server.Register(&Handler{})

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Printf("Error starting server %v", err)
	}
	defer listener.Close()

	server.Accept(listener)
	return listener.Addr().String()
}

func ClientRPC(host, port, message string) error {
	client, err := rpc.Dial("tcp", net.JoinHostPort(host, port))
	if err != nil {
		return err
	}
	defer client.Close()

	response := new(Response)
	client.Call("Handler.Execute", Request{Name: message}, response)
	log.Printf("Client received: %v", response.Message)
	return nil
}

func (h *Handler) Execute(req Request, res *Response) (err error) {
	if req.Name == "" {
		err = errors.New("A name must be specified")
		return
	}
	log.Printf("Server received: %v", req.Name)
	res.Message = "Hello " + req.Name
	return
}
