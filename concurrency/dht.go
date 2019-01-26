package concurrency

import (
	"crypto/sha1"
	"encoding/hex"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/golang/protobuf/proto"

	"github.com/kierachell/practice/data"
)

type ChordNode struct {
	Id          string
	store       map[string]string
	fingerTable map[string]*ChordNode
	successor   *ChordNode
	predecessor *ChordNode
	output      chan *data.LookupResponse
	ringSize    uint64
}

func NewChordNode(ringSize int) *ChordNode {
	hash := sha1.New()
	seed := strconv.Itoa(rand.Intn(ringSize))
	n := ChordNode{
		Id: hex.EncodeToString(hash.Sum([]byte(seed))),
	}
	n.ringSize = uint64(ringSize)
	n.store = make(map[string]string)
	n.fingerTable = make(map[string]*ChordNode)
	n.output = make(chan *data.LookupResponse)
	return &n
}

func (n *ChordNode) Insert(req *data.InsertRequest) *data.InsertResponse {
	hash := sha1.New()
	ownHash, err := strconv.ParseUint(n.Id[:5], 16, 64)
	if err != nil {
		log.Printf("Own hash is broken: %v", err)
	}
	keyHash, err := strconv.ParseUint(hex.EncodeToString(hash.Sum([]byte(req.Key)))[:5], 16, 64)
	if err != nil {
		log.Printf("Key hash is broken: %v", err)
	}

	if ownHash%n.ringSize < keyHash%n.ringSize {
		for nodeId, node := range n.fingerTable {
			nodeHash, err := strconv.ParseUint(nodeId[:5], 16, 64)
			if err != nil {
				log.Printf("Node hash is broken: %v", err)
			}
			if nodeHash%n.ringSize >= keyHash%n.ringSize {
				log.Printf("Passing key %v to node %v", req.Key, nodeId)
				res := node.Insert(req)
				return res
			}
		}
	}

	log.Printf("Accepting key %v at node %v", req.Key, n.Id)
	n.store[req.Key] = req.Value
	res := &data.InsertResponse{}
	return res
}

func (n *ChordNode) Join(req *data.JoinRequest, node *ChordNode) *data.JoinResponse {
	if req == nil || node == nil {
		return nil
	}
	n.fingerTable[req.NodeId] = node
	node.fingerTable[n.Id] = n
	n.predecessor = node
	n.successor = node
	node.successor = n
	node.predecessor = n
	res := &data.JoinResponse{}
	return res
}

func (n *ChordNode) Lookup(req *data.LookupRequest) *data.LookupResponse {
	text, err := proto.Marshal(req)
	if err != nil {
		return nil
	}
	log.Printf("Looking up: %s in\n\t%s", text, n.Id)
	res := &data.LookupResponse{}
	val, ok := n.store[req.Key]
	if !ok {
		if req.Origin == n.Id {
			err := &data.Error{
				Text: "Value not found.",
			}
			res.Error = err
		} else {
			if req.Origin == "" {
				req.Origin = n.Id
			}
			n.successor.Lookup(req)
			select {
			case result := <-n.output:
				val = result.Value
			default:
				log.Printf("Checking other nodes")
			}
		}
	}
	if val != "" && req.Origin != "" {
		res.Value = val
		origin, ok := n.fingerTable[req.Origin]
		if ok {
			go origin.pass(res)
			time.Sleep(1 * time.Millisecond)
		}
	}
	return res
}

func (n *ChordNode) pass(res *data.LookupResponse) error {
	n.output <- res
	return nil
}

func (n *ChordNode) Leave() error {
	return nil
}
