package main

import (
	"github.com/kierachell/practice/generators"
	"testing"
)

func BenchmarkTwoSum10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testData := generators.GenerateTwoSum()
		for target, arr := range testData {
			result := twoSum(arr, target)
			b.Logf("Testing array %v with target %v", arr, target)
			b.Logf("Resulting indicies %v", result)
		}
	}
}

func BenchmarkAddTwoNumbers(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var sumNumber []int
		l1 := generators.GenerateListNodes()
		l2 := generators.GenerateListNodes()
		sumNode := addTwoNumbers(l1, l2)
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
		b.Logf("Length of substring: %v\n", lengthOfLongestSubstring(randomString))
	}
}

func BenchmarkRabinKarpHash(b *testing.B) {
	for n := 0; n < b.N; n++ {
		text := generators.GenerateString(1000, 100)
		pattern := generators.GenerateString(5, 2)
		b.Logf("Found %q at position %v\n", pattern, hashRabinKarp(text, pattern))
	}
}
