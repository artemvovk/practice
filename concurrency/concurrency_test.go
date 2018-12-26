package concurrency

import (
	"github.com/kierachell/practice/generators"
	"testing"
)

func TestChanOverChan(t *testing.T) {
	result := AckChannels(generators.GenerateWork)
	t.Logf("Did %v work", result)
}
