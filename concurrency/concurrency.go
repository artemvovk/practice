package concurrency

import (
	"log"
	"sync"
)

type Runnable func(int) bool

func AckChannels(work Runnable) bool {
	sendChan := make(chan chan bool)
	for i := 0; i < 10; i++ {
		go func(i int) {
			in := <-sendChan
			in <- work(i)
		}(i)
	}

	recvChan := make(chan bool)
	for i := 0; i < 10; i++ {
		sendChan <- recvChan
	}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			done := <-recvChan
			log.Printf("Receiving work: %v\n", done)
		}()
	}
	wg.Wait()
	return true
}
