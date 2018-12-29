package concurrency

import (
	"github.com/kierachell/practice/generators"
	"math/rand"
	"sync"
	"testing"
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
