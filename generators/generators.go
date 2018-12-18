package generators

import (
	"github.com/kierachell/practice/data"
	"math/rand"
	"time"
)

func GenerateTwoSum() map[int][]int {
	rand.Seed(time.Now().UTC().UnixNano())
	testData := make(map[int][]int)
	size := rand.Intn(10)
	for idx := 0; idx < size; idx++ {
		length := rand.Intn(10) + 2
		array := make([]int, length)
		for i := 0; i < len(array); i++ {
			array[i] = rand.Intn(length)
		}
		target := array[rand.Intn(length)] + array[rand.Intn(length)]
		testData[target] = array
	}
	return testData
}

func GenerateListNodes() *data.ListNode {
	rand.Seed(time.Now().UTC().UnixNano())
	headNode := data.ListNode{
		Val: rand.Intn(9),
	}
	headNode.Next = &data.ListNode{}
	next := headNode.Next
	for idx := 0; idx <= rand.Intn(10); idx++ {
		next.Val = rand.Intn(9)
		next.Next = &data.ListNode{}
		next = next.Next

	}
	return &headNode
}

func GenerateString(maxLen int, minLen int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyz"
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(maxLen) + minLen
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
