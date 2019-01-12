package concurrency

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/kierachell/practice/data"
)

type Listener struct {
	id      int64
	state   data.State
	entries chan interface{}
	output  chan interface{}
	peers   chan interface{}
	quit    chan struct{}
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
	l.entries = make(chan interface{}, 5)
	l.peers = make(chan interface{}, 5)
	l.output = make(chan interface{}, 5)
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
			case input := <-l.entries:
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
	go func() {
		for {
			time.Sleep(20 * time.Millisecond)
			l.RequestVote()

		}
	}()
}

func (l *Listener) Respond(commit uint64, added bool) {
	info := &data.AppendEntryResponse{
		CommitIndex: commit,
		Added:       added,
	}
	response, err := proto.Marshal(info)
	if err != nil {
		log.Fatalf("Unable to marshal response: %v", err)
	}
	l.output <- response
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

func (l *Listener) RequestVote() bool {
	sent := false
	info := &data.RequestVoteRequest{
		Term:         l.state.CurrentTerm + 1,
		CandidateId:  fmt.Sprintf("%v", l.id),
		LastLogIndex: l.state.CommitIndex,
		LastLogTerm:  l.state.CurrentTerm,
	}
	request, err := proto.Marshal(info)
	if err != nil {
		log.Printf("Failed to request votes: %v", err)
		sent = false
	}
	l.peers <- request
	sent = true
	return sent
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

func (c *Controller) DetermineLeader() uint64 {
	var term uint64
	c.Workers.Iter(func(l *Listener) {
		select {
		case peer := <-l.peers:
			info := &data.RequestVoteRequest{}
			err := proto.Unmarshal(peer.([]byte), info)
			if err != nil {
				log.Printf("Failed reading message: %v", err)
			}
			log.Printf("Received vote request from %v", info.CandidateId)
			term = info.Term
		default:
			term = 0
		}
		term = 0
	})
	return term
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
			l.entries <- message
			resp := <-l.output
			info := &data.AppendEntryResponse{}
			err = proto.Unmarshal(resp.([]byte), info)
			if err != nil {
				log.Printf("Failed reading listener response: %v", err)
				c.Wg.Done()
				continue
			}
			if !info.Added {
				log.Printf("Failed to add entry commit: %v", info.CommitIndex)
				idx = info.CommitIndex
				retryMessage, err := proto.Marshal(logs[idx])
				if err != nil {
					log.Printf("Failed to retry addit log entry: %v", err)
					c.Wg.Done()
					continue
				}
				log.Printf("Retrying entry at commit: %v", logs[idx].LeaderCommit)
				l.entries <- retryMessage
				resp = <-l.output
				err = proto.Unmarshal(resp.([]byte), info)
				if err != nil {
					log.Printf("Failed to proccess listener response: %v", err)
					c.Wg.Done()
					continue
				}
			}
			c.Wg.Done()
		}
	})
	return true
}
