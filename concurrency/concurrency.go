package concurrency

import (
	"log"
	"sync"

	"github.com/kierachell/practice/generators"
)

func AckChannels() bool {
	sendChan := make(chan chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			in := <-sendChan
			in <- generators.GenerateWork(10)
		}()
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
