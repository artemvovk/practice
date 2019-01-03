package algorithms

import (
	"math/rand"
	"testing"

	"github.com/kierachell/practice/generators"
)

func BenchmarkTwoSum10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testData := generators.GenerateTwoSum()
		for target, arr := range testData {
			result := TwoSum(arr, target)
			b.Logf("Testing array %v with target %v", arr, target)
			b.Logf("Resulting indicies %v", result)
		}
	}
}

func BenchmarkAddTwoNumbers(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var sumNumber []int
		l1 := generators.GenerateListNodes(10)
		l2 := generators.GenerateListNodes(10)
		sumNode := AddTwoNumbers(l1, l2)
		next := sumNode.Next
		for {
			sumNumber = append(sumNumber, sumNode.Val)
			if next == nil {
				break
			}
			sumNode = next
			next = sumNode.Next
		}
		b.Logf("Resulting number: %v\n", sumNumber)
	}
}

func BenchmarkLengthOfLongestSubstring(b *testing.B) {
	for n := 0; n < b.N; n++ {
		randomString := generators.GenerateString(30, 1)
		b.Logf("Generated string: %s\n", randomString)
		b.Logf("Length of substring: %v\n", LengthOfLongestSubstring(randomString))
	}
}

func BenchmarkAnagramSearch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		text := generators.GenerateString(300, 20)
		pattern := "kjhll"
		indecies := AnagramSearch(pattern, text)
		if len(indecies) > 0 {
			b.Logf("Found %s in %s at %v\n", pattern, text, indecies)
		}
	}
}

func BenchmarkListIntersection(b *testing.B) {
	l1 := generators.GenerateListNodes(10)
	l2 := generators.GenerateListNodes(20)
	index := 0
	r := rand.Intn(9)
	c1 := l1
	c2 := l2
	for c1 != nil {
		if index == r {
			c1.Next = c2.Next
			break
		}
		c1 = c1.Next
		c2 = c2.Next
		index += 1
	}
	for n := 0; n < b.N; n++ {
		b.Logf("Intersect node: %v", LinkedListIntersect(l1, l2))
	}
}
