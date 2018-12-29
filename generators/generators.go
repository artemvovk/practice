package generators

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/kierachell/practice/data"
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

func GenerateWait(wait int) bool {
	done := true
	time.Sleep(time.Duration(wait) * time.Millisecond)
	fmt.Sprintf("Worked for %v milliseconds\n", wait)
	return done
}

func GenerateWork(number int) bool {
	done := true
	iterations := float64(number * 1000)
	for i := float64(0); i <= math.Abs(iterations); i++ {
		for j := float64(0); j <= math.Abs(iterations); j++ {
			out := math.Tan(math.Atan(math.Tan(math.Atan(math.Tan(math.Atan(iterations))))))
			fmt.Sprintf("Calculated %v\n", out)

		}
	}
	return done
}
