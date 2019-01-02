package server

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/net/websocket"
)

var echoTimes = 1

func ServerWebSocket(host, port string) error {
	router := mux.NewRouter()
	ws := websocket.Server{
		Handler: websocket.Handler(echoRequest),
	}
	router.PathPrefix("/").Handler(ws)
	server := &http.Server{
		Handler:      router,
		Addr:         net.JoinHostPort(host, port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	go func() {
		select {
		case err := <-errHandler:
			log.Printf("Error happened: %v", err.Error())
		default:
		}
	}()
	return server.ListenAndServe()
}

func echoRequest(ws *websocket.Conn) {
	var err error
	buf := make([]byte, maxBufferSize)
	if _, err := ws.Read(buf); err != nil {
		errHandler <- err
	}
	// All this work is just to test multiple writes from the server
	// 'Cause that's what sockets are all about
	re := regexp.MustCompile("[0-9]+")
	echoStr := re.FindAllString(string(buf), 1)[0]
	if echoTimes, err = strconv.Atoi(echoStr); err != nil {
		errHandler <- err
	}
	for i := 0; i < echoTimes; i++ {
		message := fmt.Sprintf("Server says: %s + %v\n", buf, i)
		if _, err := ws.Write([]byte(message)); err != nil {
			errHandler <- err
		}

		log.Printf("Server sent: %s.\n", message)
	}
}

func ClientWebSocket(host, port, message string, times int) error {
	url := "ws://" + host + ":" + port + "/"
	origin := "http://" + host + "/"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		return err
	}
	send := fmt.Sprintf("Client: %s+%v", message, times)
	if _, err := ws.Write([]byte(send)); err != nil {
		errHandler <- err
	}
	for {
		var msg = make([]byte, maxBufferSize)
		if _, err = ws.Read(msg); err != nil {
			return err
		}
		log.Printf("Client received: %s.\n", msg)
	}
	return nil
}
