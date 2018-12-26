package hashers

import (
	"github.com/kierachell/practice/generators"
	"testing"
)

func BenchmarkRabinKarpHash(b *testing.B) {
	for n := 0; n < b.N; n++ {
		text := generators.GenerateString(1000, 100)
		pattern := generators.GenerateString(5, 2)
		b.Logf("Found %q at position %v\n", pattern, HashRabinKarp(text, pattern))
	}
}

func BenchmarkMurmurHash(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testString := generators.GenerateString(30, 20)
		hashed := HashMurmur3([]byte(testString), uint32(len(testString)))
		b.Logf("String: %s hashed as: %v\n", testString, hashed)
	}
}

func BenchmarkBloomFilter(b *testing.B) {
	for n := 0; n < b.N; n++ {
		hasher := func(data []byte, seed uint32) uint32 {
			return HashMurmur3(data, seed)
		}
		filter := NewBloomFilter(2, 0.1, hasher)
		testString := generators.GenerateString(20, 5)
		filter.Add(testString)
		b.Logf("Checking test string that was just added: %v\n", filter.Check(testString))
		b.Logf("Checking test string that was not added: %v\n", filter.Check(generators.GenerateString(100, 30)))
		b.Logf("Resulting bit array: %v\n", filter.BitArray)
	}
}

func BenchmarkFNV1a(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testString := generators.GenerateString(30, 20)
		hashed := HashFNV1a([]byte(testString), uint32(len(testString)))
		b.Logf("String: %s hashed as: %v\n", testString, hashed)
	}
}
