package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/kierachell/practice/data"
	"github.com/kierachell/practice/generators"
)

func BenchmarkChanOverChan(b *testing.B) {
	for n := 0; n < b.N; n++ {
		result := AckChannels(generators.GenerateWork)
		b.Logf("Did %v work", result)
	}
}

func BenchmarkPhilosophers(b *testing.B) {
	for n := 0; n < b.N; n++ {
		table := Init(5)
		b.Logf("Made a table of %v\n", table)
		var wg sync.WaitGroup
		for i := 0; i < n; i++ {
			index := rand.Intn(5)
			wg.Add(1)
			go func(seat Seat) {
				defer wg.Done()
				if !seat.Occupant.Thinking {
					seat.Occupant.Eat()
				}
				if !seat.Occupant.Eating {
					seat.Occupant.Think()
				}

			}(table[index])
		}
		wg.Wait()
	}
}

func BenchmarkPassingMessages(b *testing.B) {
	for n := 10; n < b.N; n++ {
		clients := &Clients{}
		for i := 0; i < 10; i++ {
			worker := NewListener()
			clients.Push(worker)
		}
		controller := &Controller{
			Workers: clients,
			Kill:    make(chan struct{}),
		}
		controller.Workers.Iter(func(l *Listener) {
			l.quit = controller.Kill
		})
		messages := make([]*data.AppendEntryRequest, n)
		r := rand.New(rand.NewSource(time.Now().Unix()))
		for _, i := range r.Perm(n) {
			req, _ := generators.GenerateAppendEntry(i)
			messages[i] = req
		}
		controller.SendMessages(messages)
		controller.DetermineLeader()
		controller.Wg.Wait()
		controller.Workers.Iter(func(l *Listener) {
			b.Logf("Client %v has %v entries", l.id, len(l.state.Log))
		})
		controller.KillAll()
	}
}

func BenchmarkDHT(b *testing.B) {
	ringSize := b.N
	testNode1 := NewChordNode(ringSize)
	b.Logf("New node: %v", testNode1.Id)

	for n := 10; n < b.N; n++ {
		testNode2 := NewChordNode(ringSize)
		for idx := 0; idx < n; idx++ {
			testNode2 = NewChordNode(ringSize)
			join := &data.JoinRequest{
				NodeId: testNode2.Id,
			}
			testNode1.Join(join, testNode2)
		}
		insert := &data.InsertRequest{
			Key:   fmt.Sprintf("%s-%v", "dank", n),
			Value: fmt.Sprintf("%s-%v", "meme", n),
		}
		testNode2.Insert(insert)
		lookup := &data.LookupRequest{
			Key: insert.Key,
		}
		value := testNode1.Lookup(lookup)
		b.Logf("Looked up: %v", value.Value)
	}
}
