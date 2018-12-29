package concurrency

import (
	"github.com/kierachell/practice/generators"
	"log"
	"math/rand"
	"time"
)

type Chopstick struct {
	ID     int
	IsUsed bool
}

type Philosopher struct {
	ID       int
	Place    *Seat
	Eating   bool
	Thinking bool
}

type Seat struct {
	Occupant   *Philosopher
	Chopsticks []*Chopstick
}

func Init(total int) []Seat {
	table := make([]Seat, total)
	chopsticks := make([]*Chopstick, total)
	for i := 0; i < total; i++ {
		chopsticks[i] = &Chopstick{
			ID: i,
		}
	}
	for i := 0; i < total; i++ {
		next := i + 1
		if i == total-1 {
			next = 0
		}
		seat := NewSeat(i, chopsticks[i], chopsticks[next])
		table[i] = seat
	}
	return table
}

func NewSeat(place int, chopstick1 *Chopstick, chopstick2 *Chopstick) Seat {
	occupant := Philosopher{ID: place}
	seat := Seat{
		Occupant:   &occupant,
		Chopsticks: []*Chopstick{chopstick1, chopstick2},
	}
	occupant.Place = &seat
	return seat
}

func (p Philosopher) Think() {
	p.Thinking = true
	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)
	log.Printf("Philosopher %v is thinking\n", &p)
	generators.GenerateWait(rand.Intn(10))
	p.Thinking = false
}

func (p Philosopher) Eat() {
	ready := false
	for !ready {
		generators.GenerateWait(rand.Intn(1))
		ready = true
		for _, chopstick := range p.Place.Chopsticks {
			if chopstick.IsUsed {
				ready = false
			}
		}
	}
	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)
	p.Eating = true
	log.Printf("Philosopher %v is eating\n", &p)
	for _, chopstick := range p.Place.Chopsticks {
		chopstick.IsUsed = true
	}
	generators.GenerateWait(rand.Intn(10))
	p.Eating = false
	for _, chopstick := range p.Place.Chopsticks {
		chopstick.IsUsed = false
	}
}
