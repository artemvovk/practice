package main

import (
	"fmt"
	"github.com/kierachell/practice/data"
	"unsafe"
)

func main() {
	fmt.Printf("Main function")
}

func twoSum(arr []int, target int) []int {
	indicies := make([]int, 2)
	complements := make(map[int]int)
	for idx, val := range arr {
		if compl, ok := complements[val]; ok {
			indicies[0] = idx
			indicies[1] = compl
			return indicies
		}
		complements[target-val] = idx
	}
	return indicies
}

func addTwoNumbers(l1 *data.ListNode, l2 *data.ListNode) *data.ListNode {
	carry := 0
	sumNode := data.ListNode{
		Val: 0,
	}
	next := &sumNode
	for {
		next.Val += carry
		carry = 0
		if l2 != nil {
			next.Val += l2.Val
			l2 = l2.Next
		}
		if l1 != nil {
			next.Val += l1.Val
			l1 = l1.Next
		}
		if next.Val >= 10 {
			carry = 1
			next.Val = next.Val % 10
		}
		if l1 != nil || l2 != nil {
			next.Next = &data.ListNode{}
			next = next.Next
		} else {
			break
		}
	}
	if carry > 0 {
		next.Next = &data.ListNode{
			Val: carry,
		}
	}
	return &sumNode
}

func lengthOfLongestSubstring(s string) int {
	byteArray := []byte(s)
	maxLen := 0
OUTER:
	for outerIdx, _ := range byteArray {
		charSet := make(map[byte]bool)
		for innerIdx, char := range byteArray[outerIdx:] {
			if _, found := charSet[char]; found {
				if innerIdx > maxLen {
					maxLen = len(charSet)
				}
				continue OUTER
			}
			charSet[char] = true
		}
		if len(charSet) > maxLen {
			maxLen = len(charSet)
		}
	}
	return maxLen
}

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
