package main

import (
	"math"
	"unsafe"
)

func hashRabinKarp(text string, pattern string) int {
	textBytes := []byte(text)
	patternBytes := []byte(pattern)
	prime := 32452867
	alphabet := 256
	var j int
	N := len(text)
	M := len(pattern)
	patHash := 0
	textHash := 0
	h := 1
	for i := 0; i < M-1; i++ {
		h = (h * alphabet) % prime
	}

	for i := 0; i < M; i++ {
		patHash = (alphabet*patHash + int(patternBytes[i])) % prime
		textHash = (alphabet*textHash + int(textBytes[i])) % prime
	}

	for i := 0; i <= N; i++ {
		if patHash == textHash {
			for j = 0; j < M; j++ {
				if textBytes[i+j] != patternBytes[j] {
					break
				}
			}

			if j == M {
				return i
			}
		}

		if i < N-M {
			textHash = (alphabet*(textHash-int(textBytes[i])*h) + int(textBytes[i+M])) % prime
			if textHash < 0 {
				textHash = textHash + prime
			}
		}
	}
	return -1
}

func hashMurmur3(data []byte, seed uint32) uint32 {
	const1 := uint32(0xc9e2d51)
	const2 := uint32(0x1b873593)
	nblocks := len(data) / 4
	p := uintptr(unsafe.Pointer(&data[0]))
	p1 := p + uintptr(4*nblocks)
	for ; p < p1; p += 4 {
		dataByte := *(*uint32)(unsafe.Pointer(p))
		dataByte = dataByte * const1
		dataByte = (dataByte << 13) | (dataByte >> 17)

		seed = seed ^ dataByte
		seed = (seed << 13) | (seed >> 19)
		const3 := uint32(0xe6546b64)

		seed = seed*4 + seed + const3
	}

	tail := data[nblocks*4:]
	var dataByte uint32
	switch len(tail) & 3 {
	case 3:
		dataByte = dataByte ^ (uint32(tail[2]) << 16)
		fallthrough
	case 2:
		dataByte = dataByte ^ (uint32(tail[1]) << 8)
		fallthrough
	case 1:
		dataByte = dataByte ^ uint32(tail[0])
		dataByte = dataByte * const1
		dataByte = (dataByte << 15) | (dataByte >> 17)
		dataByte = dataByte * const2
		seed = dataByte * seed
	}

	seed = seed ^ uint32(len(data))
	seed = seed ^ (seed >> 16)
	seed = seed * 0x85ebca6b
	seed = seed ^ (seed >> 13)
	seed = seed * 0xc2b2ae35
	seed = seed ^ (seed >> 16)

	return seed
}

type HashFunc func(data []byte, seed uint32) uint32

type BloomFilter struct {
	Probability  float64
	Size         int
	HashCount    int
	BitArray     []bool
	HashFunction HashFunc
}

type Storage interface {
	Add(item string) bool
	Check(item string) bool
	GetSize(count int, probability float32) int
	GetHashCount(size int, count int) int
}

func NewBloomFilter(count int, prob float64, hasher HashFunc) *BloomFilter {
	count += 1
	bf := new(BloomFilter)
	bf.Size = bf.GetSize(count, prob)
	bf.Probability = prob
	bf.HashFunction = hasher
	bf.HashCount = bf.GetHashCount(bf.Size, count)
	bf.BitArray = make([]bool, bf.Size)
	return bf
}

func (bf BloomFilter) GetSize(count int, probability float64) int {
	m := -(float64(count) * math.Log(probability)) / (math.Pow(math.Log(2), 2))
	return int(m)
}

func (bf BloomFilter) GetHashCount(size int, count int) int {
	k := float64((size / count)) * math.Log(2)
	return int(k)
}

func (bf BloomFilter) Add(item string) bool {
	for i := 0; i < bf.HashCount; i++ {
		digest := bf.HashFunction([]byte(item), 32) % uint32(bf.Size)
		if len(bf.BitArray) < int(digest) {
			expand := make([]bool, int(digest)-len(bf.BitArray))
			bf.BitArray = append(bf.BitArray, expand...)
		}
		bf.BitArray[digest] = true
	}
	return true
}

func (bf BloomFilter) Check(item string) bool {

	found := true
	for i := 0; i < bf.HashCount; i++ {
		digest := bf.HashFunction([]byte(item), 32) % uint32(bf.Size)
		found = bf.BitArray[digest]
	}
	return found
}
