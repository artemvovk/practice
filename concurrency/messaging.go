package concurrency

import (
	"log"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/kierachell/practice/data"
)

type Listener struct {
	id     int64
	state  data.State
	source chan interface{}
	output chan ListenerResponse
	quit   chan struct{}
}

type ListenerResponse struct {
	CommitIndex uint64
	Added       bool
}

type Controller struct {
	Wg      sync.WaitGroup
	Workers *Clients
	Kill    chan struct{}
}

type Clients struct {
	sync.Mutex
	listeners []*Listener
}

func (c *Controller) KillAll() {
	close(c.Kill)
}

func NewListener() *Listener {
	l := &Listener{}
	l.id = time.Now().UTC().UnixNano()
	l.source = make(chan interface{}, 5)
	l.output = make(chan ListenerResponse, 5)
	l.InitState()
	l.Start()
	return l
}

func (l *Listener) InitState() {
	newstate := data.State{
		CurrentTerm: 0,
		CommitIndex: 0,
		LastApplied: 0,
	}
	l.state = newstate
}

func (l *Listener) Start() {
	go func() {
		for {
			select {
			case input := <-l.source:
				info := &data.AppendEntryRequest{}
				err := proto.Unmarshal(input.([]byte), info)
				if err != nil {
					log.Printf("Failed reading message: %v", err)
				}
				l.Respond(l.state.CommitIndex, l.ProcessEntry(info))
			case <-l.quit:
				return
			}
		}
	}()
}

func (l *Listener) Respond(commit uint64, added bool) {
	l.output <- ListenerResponse{
		CommitIndex: commit,
		Added:       added,
	}
}

func (l *Listener) ProcessEntry(info *data.AppendEntryRequest) bool {
	added := false
	if l.state.CurrentTerm > info.Term {
		log.Printf("Wrong term: %v", info.Term)
		return added
	}
	if l.state.CommitIndex != info.PrevLogEntry {
		log.Printf("Last log entry mismatch: have %v got %v", l.state.CommitIndex, info.PrevLogEntry)
		return added
	} else {
		log.Printf("Worker %v | Got log entry: %v-%v", l.id, info.LeaderCommit, info.Entries)
		for _, entry := range info.Entries {
			l.state.Log = append(l.state.Log, entry)
		}
	}
	if info.LeaderCommit > l.state.CommitIndex {
		if l.state.CommitIndex < info.LeaderCommit {
			l.state.CommitIndex = uint64(len(l.state.Log) - 1)
		} else {
			l.state.CommitIndex = info.LeaderCommit

		}
	}
	return true
}

func (c *Clients) Push(l *Listener) {
	c.Lock()
	defer c.Unlock()
	c.listeners = append(c.listeners, l)
}

func (c *Clients) Iter(apply func(*Listener)) {
	c.Lock()
	defer c.Unlock()

	for _, listener := range c.listeners {
		apply(listener)
	}
}

func (c *Controller) SendMessages(logs []*data.AppendEntryRequest) bool {
	c.Workers.Iter(func(l *Listener) {
		for idx := uint64(0); idx < uint64(len(logs)); idx++ {
			c.Wg.Add(1)
			message, err := proto.Marshal(logs[idx])
			if err != nil {
				c.Wg.Done()
				continue
			}
			l.source <- message
			resp := <-l.output
			if !resp.Added {
				log.Printf("Failed to add entry commit: %v", resp.CommitIndex)
				idx = resp.CommitIndex
				retryMessage, err := proto.Marshal(logs[idx])
				if err != nil {
					c.Wg.Done()
					continue
				}
				log.Printf("Retrying entry at commit: %v", logs[idx].LeaderCommit)
				l.source <- retryMessage
				resp = <-l.output

			}
			c.Wg.Done()
		}
	})
	return true
}
